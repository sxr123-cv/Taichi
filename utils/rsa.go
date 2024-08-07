package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

func RsaSign(prvkey []byte, hash crypto.Hash, data []byte) ([]byte, error) {
	block, _ := pem.Decode(prvkey)
	if block == nil {
		return nil, errors.New("decode private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.SignPKCS1v15(rand.Reader, privateKey, hash, data)
}

func RsaVerifySign(pubkey []byte, hash crypto.Hash, data, sig []byte) error {
	block, _ := pem.Decode(pubkey)
	if block == nil {
		return errors.New("decode public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(pub, hash, data, sig)
}

func Hash(hash crypto.Hash, data []byte) []byte {
	var hashed []byte
	switch hash {
	case crypto.SHA224:
		h := sha256.Sum224(data)
		hashed = h[:]
	case crypto.SHA256:
		h := sha256.Sum256(data)
		hashed = h[:]
	case crypto.SHA384:
		h := sha512.Sum384(data)
		hashed = h[:]
	case crypto.SHA512:
		h := sha512.Sum512(data)
		hashed = h[:]
	}
	return hashed
}
