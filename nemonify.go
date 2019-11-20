package main

import (
	"encoding/base64"

	mnemonics "github.com/RTradeLtd/entropy-mnemonics"
)

// Encode is used to encode a piece of data
func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Decode is used to decode a piece of encoded data
func Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// ToMnemonic takes a string an turns it into a mnemonic
func ToMnemonic(msg string) (string, error) {
	phrase, err := mnemonics.ToPhrase([]byte(msg), mnemonics.English)
	if err != nil {
		return "", err
	}
	return phrase.String(), nil
}

// FromMnemonic takes a mnemonic phrase and returns the underlying string
func FromMnemonic(phrase string) (string, error) {
	mnemonicBytes, err := mnemonics.FromString(phrase, mnemonics.English)
	if err != nil {
		return "", err
	}
	return string(mnemonicBytes), nil
}
