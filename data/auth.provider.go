package data

import (
	"time"

	"arnobot-shared/db"
)

type AuthProvider struct {
	ID             int32
	UserID         int32
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
		ID:             fromDB.ID,
		UserID:         fromDB.UserID,
		Provider:       fromDB.Provider,
		ProviderUserID: fromDB.ProviderUserID,
		AccessToken:    fromDB.AccessToken,
		RefreshToken:   fromDB.RefreshToken,
		AccessType:     fromDB.AccessType,
		Scopes:         fromDB.Scopes,
		CreatedAt:      fromDB.CreatedAt,
		UpdatedAt:      fromDB.UpdatedAt,
	}
}

type AuthProviderCreate struct {
	UserID         int32
	Provider       string
	ProviderUserID string
	AccessToken    string
	RefreshToken   string
	AccessType     string
	Scopes         []string
}

func (p AuthProviderCreate) ToDB() db.AuthProviderCreateParams {
	scopes := p.Scopes
	if scopes == nil {
		scopes = make([]string, 0)
	}

	return db.AuthProviderCreateParams{
		UserID:         p.UserID,
		Provider:       p.Provider,
		ProviderUserID: p.ProviderUserID,
		AccessToken:    p.AccessToken,
		RefreshToken:   p.RefreshToken,
		AccessType:     p.AccessType,
		Scopes:         scopes,
	}
}

type AuthProviderUpdateTokens struct {
	AccessToken  string
	RefreshToken *string
}

func (p AuthProviderUpdateTokens) ToDB(id int32) db.AuthProviderUpdateTokensParams {
	return db.AuthProviderUpdateTokensParams{
		ID:           id,
		AccessToken:  p.AccessToken,
		RefreshToken: p.RefreshToken,
	}
}

type AuthProviderUpdate struct {
	UserID         int32
	ProviderUserID string
	AccessToken    string
	RefreshToken   string
	AccessType     string
}

type AuthProviderGet struct {
	ProviderUserID *string
	UserID         *int32
	Provider       string
}

func (p AuthProviderGet) ToDB() db.AuthProviderGetParams {
	return db.AuthProviderGetParams{
		ProviderUserID: p.ProviderUserID,
		UserID:         p.UserID,
		Provider:       p.Provider,
	}
}
