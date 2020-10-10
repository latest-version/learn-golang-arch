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

	for i := 1; i <= 64; i++ {
		key = append(key, byte(i))
	}

	// fmt.Println("Key :", key)
	// // fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))
	// password := "12345678"
	// hashedPass, err := hashPassword(password)

	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(hashedPass)
	// err = comparePassword(password, hashedPass)
	// if err != nil {
	// 	log.Fatalln("Not logged in")
	// }

	// log.Println("Logged in!")

	valid := false
	bs, err := signMessage([]byte("Hello"))
	if err != nil {
		log.Fatalln("Signature Invalid")
	}
	fmt.Println("Message :", bs)

	valid, err = chechSig([]byte("Hello"), bs)
	if err != nil {
		log.Fatalln("Signature Compare Invalid")
	}

	fmt.Println(valid)

}

func hashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("Error while generating bcrypt hash from password: %w", err)
	}
	return bs, nil
}

func comparePassword(password string, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
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

	signature := h.Sum(nil)
	return signature, nil
}

func chechSig(msg, sig []byte) (bool, error) {
	newSig, err := signMessage(msg)
	if err != nil {
		return false, fmt.Errorf("Error in checkSig while getting signature message: %w", err)
	}

	same := hmac.Equal(newSig, sig)
	return same, nil
}
