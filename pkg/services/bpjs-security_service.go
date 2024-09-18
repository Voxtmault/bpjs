package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	daku "github.com/daku10/go-lz-string"
	"github.com/rotisserie/eris"
	"github.com/voxtmault/bpjs-rs-module/config"
	"github.com/voxtmault/bpjs-rs-module/pkg/interfaces"
)

type BPJSSecurityService struct {
}

var _ interfaces.BPJSSecurity = &BPJSSecurityService{}

func (s *BPJSSecurityService) CreateSignature(ctx context.Context, timestamp int64) (string, error) {
	// Explanations
	// 1. Signature is created using HMAC-SHA256
	// 2. Signature requires ConsumerID, Timestamp, and ConsumerSecret

	cfg := config.GetConfig().BPJSConfig

	h := hmac.New(sha256.New, []byte(cfg.ConsumerSecret))
	message := fmt.Sprintf("%s&%d", cfg.ConsumerID, timestamp)

	h.Write([]byte(message))

	signature := h.Sum(nil)

	// Base64 encode the signature
	encodedSignature := base64.StdEncoding.EncodeToString(signature)

	// URL encode the base64-encoded signature (if necessary)
	// urlEncodedSignature := url.QueryEscape(encodedSignature)

	return encodedSignature, nil
}

func (s *BPJSSecurityService) DecryptResponse(ctx context.Context, timestamp int64, message string) (string, error) {
	// Logic
	// 1. Decrypt the message using AES256 (CBC Mode) - SHA256
	// 2. Decompress using Lz-String method decompressFromEncodedURIComponent

	cfg := config.GetConfig().BPJSConfig

	key := sha256.Sum256([]byte(fmt.Sprintf("%s%s%d", cfg.ConsumerID, cfg.ConsumerSecret, timestamp)))

	log.Println("Message", message)
	log.Println("Key", fmt.Sprintf("%x", key))

	// Base64-decode the URL-decoded message
	encryptedData, err := base64.StdEncoding.DecodeString(strings.ReplaceAll(message, `\/`, "/"))
	// encryptedData, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	// Extract the IV from the encrypted message
	if len(encryptedData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := key[:aes.BlockSize]
	// ciphertext := key[aes.BlockSize:]

	// Create the AES cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", eris.Wrap(err, "failed to create AES cipher")
	}

	// Decrypt the message using AES-256 in CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(encryptedData))
	mode.CryptBlocks(plainText, encryptedData)

	// // Remove padding
	// plaintext, err := pkcs7Unpad(ciphertext, aes.BlockSize)
	// if err != nil {
	// 	return "", eris.Wrap(err, "failed to remove padding")
	// }

	// encodedPlainText := base64.StdEncoding.EncodeToString(plaintext)

	// data, err := lzstring.Decompress(encodedPlainText, "")
	// if err != nil {
	// 	return "", eris.Wrap(err, "failed to decompress data")
	// }

	data, err := daku.DecompressFromEncodedURIComponent(string(plainText))
	if err != nil {
		return "", eris.Wrap(err, "failed to decompress data")
	}

	log.Println("Data: ", data)

	return data, nil
}

// pkcs7Unpad removes padding from the decrypted message
func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("invalid padding size")
	}
	padding := int(data[length-1])
	if padding > blockSize || padding == 0 {
		return nil, errors.New("invalid padding size")
	}
	for i := 0; i < padding; i++ {
		if data[length-1-i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}
	return data[:length-padding], nil
}
