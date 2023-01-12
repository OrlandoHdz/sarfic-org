package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"time"

	random "math/rand"

	"github.com/golang-jwt/jwt/v4"
)

func Encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Sencrypt(plaintext string, passphrase string) (chipertext string, err error) {
	block, _ := aes.NewCipher([]byte(passphrase))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	ciphertextByte := gcm.Seal(
		nonce,
		nonce,
		[]byte(plaintext),
		nil)
	chipertext = base64.StdEncoding.EncodeToString(ciphertextByte)
	return
}

func Sdecrypt(cipherText, key string) (plainText string, err error) {
	// prepare cipher
	keyByte := []byte(key)
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonceSize := gcm.NonceSize()
	//
	// process ciphertext
	fmt.Println("cipherText:", cipherText)
	ciphertextByte, _ := base64.StdEncoding.DecodeString(cipherText)
	nonce, ciphertextByteClean := ciphertextByte[:nonceSize], ciphertextByte[nonceSize:]
	plaintextByte, err := gcm.Open(
		nil,
		nonce,
		ciphertextByteClean,
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	plainText = string(plaintextByte)
	//
	return
}

type SmxClaims struct {
	Smx string `json:"smx"`
	Rdn int    `json:"rdn"`
	jwt.StandardClaims
}

func GeneraJWT(nombreUsuario string) (string, error) {

	firma := []byte("facturacomerciaal14")

	random.Seed(time.Now().UnixNano())
	num := random.Intn(1100)

	claims := SmxClaims{
		"smx",
		num,
		jwt.StandardClaims{
			ExpiresAt: 11000,
			Issuer:    nombreUsuario,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(firma)

	return ss, err
}

func Saludar() string {
	return "Te saludo desde el mod"
}
