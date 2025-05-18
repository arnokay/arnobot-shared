package service

import (
	"context"
	"log/slog"

	"github.com/nicklaw5/helix/v2"

	"arnobot-shared/applog"
	"arnobot-shared/db"
	"arnobot-shared/pkg/errs"
)

type HelixManager struct {
	query        db.Querier
	logger       *slog.Logger
	clientID     string
	clientSecret string
}

func NewHelixManager(querier db.Querier, clientID, clientSecret string) *HelixManager {
	logger := applog.NewServiceLogger("HelixManager")

	return &HelixManager{
		query:        querier,
		logger:       logger,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (hm *HelixManager) GetProviderByTwitchID(ctx context.Context, twitchID string) (*db.AuthProvider, error) {
	provider, err := hm.query.AuthProviderGetByProviderUserId(ctx, db.AuthProviderGetByProviderUserIdParams{
		ProviderUserID: twitchID,
		Provider:       "twitch",
	})
	if err != nil {
		hm.logger.Error("cannot get auth provider", "id", twitchID, "err", err)
		return nil, errs.ErrNotFound
	}

	return &provider, nil
}

func (hm *HelixManager) GetByTwitchID(ctx context.Context, provider db.AuthProvider) (*helix.Client, error) {
	client, _ := helix.NewClient(&helix.Options{
		ClientID:        hm.clientID,
		ClientSecret:    hm.clientSecret,
		UserAccessToken: provider.AccessToken,
		RefreshToken:    provider.RefreshToken,
	})

	client.OnUserAccessTokenRefreshed(func(newAccessToken, newRefreshToken string) {
		count, err := hm.query.AuthProviderUpdateTokens(ctx, db.AuthProviderUpdateTokensParams{
			ID:           provider.ID,
			AccessToken:  newAccessToken,
			RefreshToken: &newRefreshToken,
		})

		if err != nil || count == 0 {
			hm.logger.ErrorContext(ctx, "cannot update provider tokens", "provider_id", provider.ID, "err", err, "count", count)
		}
	})

	return client, nil
}
