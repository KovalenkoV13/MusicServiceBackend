package main

import (
	redisCL "MusicServiceBackend/internal/redis"
	"MusicServiceBackend/internal/role"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"strings"
)

const jwtPrefix = "Bearer "

func (a *Application) WithAuthCheck(next http.Handler, assignedRoles ...role.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtStr := r.Header.Get("Authorization")
		f := false
		if !strings.HasPrefix(jwtStr, jwtPrefix) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		jwtStr = jwtStr[len(jwtPrefix):]
		fmt.Printf("middleware ")
		err := a.redis.CheckJWTInBlacklist(r.Context(), jwtStr)
		if err == nil {
			return
		}
		if !errors.Is(err, redis.Nil) {
			return
		}

		token, err := jwt.ParseWithClaims(jwtStr, &redisCL.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("MusicService"), nil
		})
		if err != nil {
			log.Println(err)
			return
		}

		myClaims := token.Claims.(*redisCL.JWTClaims)

		for _, oneOfAssignedRole := range assignedRoles {
			if myClaims.Role == oneOfAssignedRole {
				fmt.Println(oneOfAssignedRole)
				next.ServeHTTP(w, r)
				f = true
			}
		}
		if f != true {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
		return
	})
}
