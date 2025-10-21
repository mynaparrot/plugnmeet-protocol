package auth

import (
	"errors"
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

// VerifyPlugNmeetAccessToken can be use to verify plugNmeet access token
// gracefulPeriod allows a token to be considered valid for this duration past its original expiration time.
// A value of 0 means no graceful period (strict expiration).
// NotBefore (nbf) validation is always strict against the current time.
func VerifyPlugNmeetAccessToken(apiKey, secret, token string, gracefulPeriod time.Duration) (*plugnmeet.PlugNmeetTokenClaims, error) {
	tok, err := jwt.ParseSigned(token, []jose.SignatureAlgorithm{jose.HS256})
	if err != nil {
		return nil, err
	}

	out := jwt.Claims{}
	claims := plugnmeet.PlugNmeetTokenClaims{}
	if err = tok.Claims([]byte(secret), &out, &claims); err != nil {
		return nil, err
	}

	// Always validate against current time for NotBefore and IssuedAt.
	// We will handle Expiry separately if a gracefulPeriod is provided.
	exp := jwt.Expected{Issuer: apiKey, Subject: claims.UserId, Time: time.Now().UTC()}

	// Perform initial validation.
	// This will check NotBefore, IssuedAt, and Expiry against time.Now().UTC().
	validationErr := out.Validate(exp)

	if validationErr != nil {
		// If there's an error, check if it's an expiration error and if a gracefulPeriod is active.
		if gracefulPeriod > 0 && errors.Is(validationErr, jwt.ErrExpired) {
			// If the token expired, but its original expiry time is within the graceful period
			// from the current time, we consider it valid.
			if out.Expiry != nil && out.Expiry.Time().After(time.Now().UTC().Add(-gracefulPeriod)) {
				// The Token is expired but within the graceful period, so it's valid.
				validationErr = nil
			}
		}
	}

	if validationErr != nil {
		return nil, validationErr
	}
	claims.UserId = out.Subject

	return &claims, nil
}
