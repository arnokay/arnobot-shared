package data

import (
	"time"

	"arnobot-shared/db"
)

type AuthProvider struct {
	ID             int
	UserID         int
	Provider       string
	ProviderUserID string
	AccessToken    string
	RefreshToken   string
	AccessType     string
	Scopes         []string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewProviderAuthFromDB(fromDB db.AuthProvider) AuthProvider {
	return AuthProvider{
		ID:             int(fromDB.ID),
		UserID:         int(fromDB.UserID),
		Provider:       fromDB.Provider,
		ProviderUserID: fromDB.ProviderUserID,
		AccessToken:    fromDB.AccessToken,
		RefreshToken:   fromDB.RefreshToken,
		AccessType:     fromDB.AccessType,
		Scopes:         fromDB.Scopes,
		CreatedAt:      fromDB.CreatedAt.Time,
		UpdatedAt:      fromDB.UpdatedAt.Time,
	}
}

type AuthProviderCreate struct {
	UserID         int
	Provider       string
	ProviderUserID string
	AccessToken    string
	RefreshToken   string
	AccessType     string
	Scopes         []string
}

func (p AuthProviderCreate) ToDB() db.AuthProviderCreateParams {
	return db.AuthProviderCreateParams{
		UserID:         int32(p.UserID),
		Provider:       p.Provider,
		ProviderUserID: p.ProviderUserID,
		AccessToken:    p.AccessToken,
		RefreshToken:   p.RefreshToken,
		AccessType:     p.AccessType,
		Scopes:         p.Scopes,
	}
}

type AuthProviderUpdateTokens struct {
	ID           int
	AccessToken  string
	RefreshToken *string
}

func (p AuthProviderUpdateTokens) ToDB() db.AuthProviderUpdateTokensParams {
	return db.AuthProviderUpdateTokensParams{
		ID:           int32(p.ID),
		AccessToken:  p.AccessToken,
		RefreshToken: p.RefreshToken,
	}
}

type AuthProviderUpdate struct {
	UserID         int
	ProviderUserID string
	AccessToken    string
	RefreshToken   string
	AccessType     string
}

type AuthProviderGet struct {
	ProviderUserID string
	Provider       string
}
