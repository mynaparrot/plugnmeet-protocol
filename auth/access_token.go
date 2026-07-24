package auth

import (
	"time"

	"github.com/go-jose/go-jose/v4"
	"github.com/go-jose/go-jose/v4/jwt"
	"github.com/livekit/protocol/auth"
	"github.com/mynaparrot/plugnmeet-protocol/plugnmeet"
)

func GeneratePlugNmeetJWTAccessToken(apiKey, secret, userId string, tokenValidity time.Duration, c *plugnmeet.PlugNmeetTokenClaims) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(secret)},
		(&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		return "", err
	}

	cl := &jwt.Claims{
		Issuer:    apiKey,
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(tokenValidity)),
		Subject:   userId,
	}
	return jwt.Signed(sig).Claims(cl).Claims(c).Serialize()
}

func GenerateLivekitAccessToken(apiKey, secret string, tokenValidity time.Duration, c *plugnmeet.PlugNmeetTokenClaims, canPublish, canSubscribe bool, originalUserId string) (string, error) {
	at := auth.NewAccessToken(apiKey, secret)
	grant := &auth.VideoGrant{
		RoomJoin:     true,
		Room:         c.RoomId,
		RoomAdmin:    c.IsAdmin,
		Hidden:       c.IsHidden,
		CanPublish:   &canPublish,
		CanSubscribe: &canSubscribe,
	}

	at.SetVideoGrant(grant).
		SetIdentity(c.UserId).
		SetName(c.Name).
		SetValidFor(tokenValidity)

	// If originalUserId is not empty we are minting a token for the native twin
	// (identity [userId]-native). Tag the participant with an attribute so
	// webhook handlers can identify the twin without relying on the identity
	// suffix alone. This attribute is NOT part of the plugNmeet JWT claims.
	if originalUserId != "" {
		at.SetAttributes(map[string]string{
			"original_user_id": originalUserId,
		})
	}

	return at.ToJWT()
}

// GenerateTokenForDownloadRecording will generate token
// path format: sub_path/roomSid/filename
func GenerateTokenForDownloadRecording(path, apiKey, apiSecret string, tokenValidity time.Duration) (string, error) {
	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(apiSecret)}, (&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		return "", err
	}

	cl := jwt.Claims{
		Issuer:    apiKey,
		NotBefore: jwt.NewNumericDate(time.Now().UTC()),
		Expiry:    jwt.NewNumericDate(time.Now().UTC().Add(tokenValidity)),
		// format: sub_path/roomSid/filename
		Subject: path,
	}

	return jwt.Signed(sig).Claims(cl).Serialize()
}
