package auth

import (
	"context"
	"github.com/869413421/micro-chat/pkg/enforcer"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

const (

	// bearerWord the bearer key word for authorization
	bearerWord string = "Bearer"

	// bearerFormat authorization token format
	bearerFormat string = "Bearer %s"

	// authorizationKey holds the key used to store the JWT Token in the request tokenHeader.
	authorizationKey string = "Authorization"

	// reason holds the error reason.
	reason string = "UNAUTHORIZED"

	reasonPermission string = "PERMISSION_DENIED"
)

var (
	ErrMissingJwtToken    = errors.Unauthorized(reason, "JWT token is missing")
	ErrTokenInvalid       = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired       = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail     = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrWrongContext       = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrorPermissionDenied = errors.Forbidden(reasonPermission, "Permission denied")
)

type JwtClaims struct {
	ID   uint64
	Name string
	jwt.RegisteredClaims
}

// GenerateToken 生成token
func GenerateToken(id uint64, name, key string) (string, error) {
	claims := JwtClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "micro-chat",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

// ParseToken 解析token
func ParseToken(tokenString, key string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Server is a server auth middleware. Check the token and extract the info from token.
func Server(jwtSecret string) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				// 验证token
				auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, ErrMissingJwtToken
				}
				jwtToken := auths[1]

				claims, err := ParseToken(jwtToken, jwtSecret)
				if err != nil {
					ve, ok := err.(*jwt.ValidationError)
					if !ok {
						return nil, errors.Unauthorized(reason, err.Error())
					}
					if ve.Errors&jwt.ValidationErrorMalformed != 0 {
						return nil, ErrTokenInvalid
					}
					if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
						return nil, ErrTokenExpired
					}
					return nil, ErrTokenParseFail
				}

				// 鉴定权限
				enf, err := enforcer.GetSyncedEnforcer()
				if err != nil {
					return nil, errors.Forbidden(reasonPermission, err.Error())
				}

				// 获取请求的路径和方法
				ht, _ := header.(*http.Transport)
				request := ht.Request()
				resource := request.URL.Path
				action := request.Method

				// 检查权限
				ok, err = enf.Enforce(claims.Name, resource, action)
				if err != nil {
					return nil, errors.Forbidden(reasonPermission, err.Error())
				}

				if !ok {
					return nil, ErrorPermissionDenied
				}

				ctx = NewContext(ctx, claims)
				return handler(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
}

type authKey struct{}

// NewContext put auth info into context
func NewContext(ctx context.Context, info *JwtClaims) context.Context {
	return context.WithValue(ctx, authKey{}, info)
}

// FromContext extract auth info from context
func FromContext(ctx context.Context) (token *JwtClaims, ok bool) {
	token, ok = ctx.Value(authKey{}).(*JwtClaims)
	return
}
