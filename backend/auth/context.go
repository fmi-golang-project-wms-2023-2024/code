package auth

import "context"

type UserContextKey string

const (
	subKey  UserContextKey = "sub"
	roleKey UserContextKey = "role"
)

func SetSubInContext(ctx context.Context, sub string) context.Context {
	return context.WithValue(ctx, subKey, sub)
}

func SetRoleInContext(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func GetSubFromContext(ctx context.Context) (string, bool) {
	sub, ok := ctx.Value(subKey).(string)
	return sub, ok
}

func GetRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(roleKey).(string)
	return role, ok
}
