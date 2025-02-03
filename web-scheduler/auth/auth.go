package auth

import "context"

func SetClientID(ctx context.Context, clientID string) context.Context {
	return context.WithValue(ctx, "clientID", clientID)
}

func GetClientID(ctx context.Context) string {
	return ctx.Value("clientID").(string)
}
