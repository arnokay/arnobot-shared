package service

import (
	"context"
	"log/slog"
	"sync"

	"github.com/nicklaw5/helix/v2"

	"github.com/arnokay/arnobot-shared/applog"
	"github.com/arnokay/arnobot-shared/data"
	"github.com/arnokay/arnobot-shared/pkg/assert"
	"github.com/arnokay/arnobot-shared/apperror"
)

// TODO: right now there is no cleanup for clients
type HelixManager struct {
	logger       *slog.Logger
	clientID     string
	clientSecret string

	appClient *helix.Client

	clients map[string]*helix.Client
	mu      sync.RWMutex

	authModuleService *AuthModuleService
}

func NewHelixManager(authModuleSerivce *AuthModuleService, clientID, clientSecret string) *HelixManager {
	logger := applog.NewServiceLogger("helix-manager")

	appClient, err := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})
	assert.NoError(err, "helix client needs to be initialized")
	
  token, err := appClient.RequestAppAccessToken([]string{
		"user:read:chat",
		"user:write:chat",
		"user:bot",
		"channel:bot",
		"bits:read",
	})
  assert.NoError(err, "cannot get access tokens for app client")
  appClient.SetAppAccessToken(token.Data.AccessToken)

	return &HelixManager{
		logger:            logger,
		clientID:          clientID,
		clientSecret:      clientSecret,
		appClient:         appClient,
		clients:           make(map[string]*helix.Client),
		authModuleService: authModuleSerivce,
	}
}

func (hm *HelixManager) GetApp(ctx context.Context) *helix.Client {
	return hm.appClient
}

func (hm *HelixManager) GetByID(ctx context.Context, twitchID string) (*helix.Client, error) {
	hm.mu.RLock()
	client, exists := hm.clients[twitchID]
	hm.mu.RUnlock()

	if exists {
		return client, nil
	}

	return nil, apperror.New(apperror.CodeNotFound, "helix client is not found", nil)
}

func (hm *HelixManager) GetByProvider(ctx context.Context, provider data.AuthProvider) *helix.Client {
	hm.mu.RLock()
	client, exists := hm.clients[provider.ProviderUserID]
	hm.mu.RUnlock()

	if exists {
		return client
	}

	hm.mu.Lock()
	defer hm.mu.Unlock()

	if client, exists := hm.clients[provider.ProviderUserID]; exists {
		return client
	}

	client, _ = helix.NewClient(&helix.Options{
		ClientID:        hm.clientID,
		ClientSecret:    hm.clientSecret,
		UserAccessToken: provider.AccessToken,
		RefreshToken:    provider.RefreshToken,
	})

	client.OnUserAccessTokenRefreshed(func(newAccessToken, newRefreshToken string) {
		hm.logger.InfoContext(ctx, "token refreshed", "providerUserID", provider.ProviderUserID)
		err := hm.authModuleService.AuthProviderUpdateTokens(ctx, provider.ID, data.AuthProviderUpdateTokens{
			AccessToken:  newAccessToken,
			RefreshToken: &newRefreshToken,
		})
		if err != nil {
			hm.logger.ErrorContext(ctx, "failed to update tokens", "providerID", provider.ID, "providerUserID", provider.ProviderUserID)
		}
	})

	hm.clients[provider.ProviderUserID] = client

	return client
}
