package utils

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

// NkeyOptionFromSeedText creates a NATS NKey option from a raw seed string.
func NkeyOptionFromSeedText(seedText string) (nats.Option, error) {
	// Create the KeyPair from the seed text.
	seedChars := []byte(seedText)
	kp, err := nkeys.FromSeed(seedChars)
	if err != nil {
		return nil, err
	}
	// Immediately wipe the raw seed text characters from memory.
	wipeSlice(seedChars)

	// Get the public key for validation.
	pub, err := kp.PublicKey()
	if err != nil {
		return nil, err
	}
	if !nkeys.IsValidPublicUserKey(pub) {
		return nil, errors.New("nats: not a valid nkey user seed")
	}

	// Create a signature callback closure that captures the KeyPair
	sigCB := func(nonce []byte) ([]byte, error) {
		return kp.Sign(nonce)
	}

	return nats.Nkey(pub, sigCB), nil
}

// wipeSlice overwrites a byte slice with 'x' to clear sensitive data from memory.
func wipeSlice(buf []byte) {
	for i := range buf {
		buf[i] = 'x'
	}
}
