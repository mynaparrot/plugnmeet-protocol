package auth

import (
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"time"
)

// VerifyPlugNmeetAccessToken can be use to verify plugNmeet access token
func VerifyPlugNmeetAccessToken(apiKey, secret, token string, withTime bool) (*plugnmeet.PlugNmeetTokenClaims, error) {
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	claims := plugnmeet.PlugNmeetTokenClaims{}
	if err = tok.Claims([]byte(secret), &out, &claims); err != nil {
		return nil, err
	}

	exp := jwt.Expected{Issuer: apiKey, Subject: claims.UserId}
	if withTime {
		exp.Time = time.Now().UTC()
	} else {
		// so the token will not be expired
		out.Expiry = nil
		out.NotBefore = nil
		out.IssuedAt = nil
	}

	if err = out.Validate(exp); err != nil {
		return nil, err
	}
	claims.UserId = out.Subject

	return &claims, nil
}
