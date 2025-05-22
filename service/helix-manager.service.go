package service

import (
	"context"
	"log/slog"
	"sync"

	"github.com/nicklaw5/helix/v2"

	"arnobot-shared/applog"
	"arnobot-shared/data"
	"arnobot-shared/pkg/errs"
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

	appClient, _ := helix.NewClient(&helix.Options{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	})

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

	return nil, errs.New(errs.CodeNotFound, "helix client is not found", nil)
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
		err := hm.authModuleService.AuthProviderUpdateTokens(ctx, data.AuthProviderUpdateTokens{
			ID:           int(provider.ID),
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
