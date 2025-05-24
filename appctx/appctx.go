package appctx

import (
	"context"

	"arnobot-shared/data"
)

type ctxKey string

const (
	USER_KEY ctxKey = "user"
	TRACE_ID ctxKey = "trace"
)

func SetUser(ctx context.Context, user *data.User) context.Context {
	ctx2 := context.WithValue(ctx, USER_KEY, user)
  return ctx2
}

func GetUser(ctx context.Context) *data.User {
	rawUser := ctx.Value(USER_KEY)
	if rawUser == nil {
		return nil
	}
	user, ok := rawUser.(*data.User)
	if !ok {
		return nil
	}
	return user
}
