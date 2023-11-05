package redis

import (
	"MusicServiceBackend/internal/role"
	"context"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	Id     int      `json:"id"`
	Scopes []string `json:"scopes" json:"scopes"`
	Role   role.Role
}

func getJWTKey(token string) string {
	return token
}

func (c *Client) WriteJWTToBlacklist(ctx context.Context, jwtStr string, jwtTTL time.Duration) error {
	return c.client.Set(ctx, jwtStr, true, jwtTTL).Err()
}

func (c *Client) CheckJWTInBlacklist(ctx context.Context, jwtStr string) error {
	return c.client.Get(ctx, getJWTKey(jwtStr)).Err()
}
