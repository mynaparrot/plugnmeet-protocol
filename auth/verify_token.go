package auth

import (
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/livekit/protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
	"time"
)

// VerifyPlugNmeetAccessToken can be use to verify plugNmeet access token
func VerifyPlugNmeetAccessToken(apiKey, secret, token string) (*plugnmeet.PlugNmeetTokenClaims, error) {
	tok, err := jwt.ParseSigned(token)
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	claims := plugnmeet.PlugNmeetTokenClaims{}
	if err = tok.Claims([]byte(secret), &out, &claims); err != nil {
		return nil, err
	}

	if err = out.Validate(jwt.Expected{Issuer: apiKey, Time: time.Now()}); err != nil {
		return nil, err
	}
	claims.UserId = out.Subject

	return &claims, nil
}

// ValidateLivekitWebhookToken can be use to validate both livekit token
func ValidateLivekitWebhookToken(secret, token string) (string, error) {
	grant, err := auth.ParseAPIToken(token)
	if err != nil {
		return "", err
	}

	claims, err := grant.Verify(secret)
	if err != nil {
		return "", err
	}

	return claims.Sha256, nil
}
