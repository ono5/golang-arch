package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

var key = []byte{}

func main() {
	for i := 01; 0 <= 64; i++ {
		key = append(key, byte(i))
	}

	// fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	pass := "123456789"
	hashedPass, err := hashPassword(pass)
	if err != nil {
		panic(err)
	}

	err = CompoarePassword(pass, hashedPass)
	if err != nil {
		log.Fatalln("Not logged in")
	}

	log.Println("Loggied in!")
}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func CompoarePassword(password string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(password))
	if err != nil {
		return fmt.Errorf("Invalid password: %w", err)
	}
	return nil
}

func signMessage(msg []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)
	if err != nil {
		return nil, fmt.Errorf("Error in signMessage while hashing message: %w", err)
	}

	signed := h.Sum(nil)
	return signed, nil
}

func checkSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil

}
