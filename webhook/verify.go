package webhook

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/livekit/protocol/auth"
	"time"
)

// VerifyRequest will verify webhook request both for livekit & plugNmeet
// in plugNmeet we're following the same token system as livekit is using
// in this method we'll verify the provided body request
func VerifyRequest(body []byte, apiKey, secret, token string) (bool, error) {
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return false, err
	}

	out := jwt.Claims{}
	claims := auth.ClaimGrants{}
	if err := tok.Claims([]byte(secret), &out, &claims); err != nil {
		return false, err
	}

	if err := out.Validate(jwt.Expected{Issuer: apiKey, Time: time.Now()}); err != nil {
		return false, err
	}

	sha := sha256.Sum256(body)
	hash := base64.StdEncoding.EncodeToString(sha[:])

	if subtle.ConstantTimeCompare([]byte(claims.Sha256), []byte(hash)) != 1 {
		return false, errors.New("authorization token didn't match")
	}

	return true, nil
}
