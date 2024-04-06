package auth

import (
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"time"
)

// VerifyPlugNmeetAccessToken can be use to verify plugNmeet access token
func VerifyPlugNmeetAccessToken(apiKey, secret, token string) (*plugnmeet.PlugNmeetTokenClaims, error) {
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	claims := plugnmeet.PlugNmeetTokenClaims{}
	if err = tok.Claims([]byte(secret), &out, &claims); err != nil {
		return nil, err
	}

	if err = out.Validate(jwt.Expected{Issuer: apiKey, Time: time.Now().UTC()}); err != nil {
		return nil, err
	}
	claims.UserId = out.Subject

	return &claims, nil
}
