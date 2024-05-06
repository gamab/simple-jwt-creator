package main

import (
	"bufio"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func readPrivateKey(filename string) (any, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := bufio.NewReader(file)
	raw, err := io.ReadAll(b)
	if err != nil {
		return nil, err
	}

	// parse PEM ECDSA private key
	block, _ := pem.Decode(raw)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the private key")
	}

	return x509.ParseECPrivateKey(block.Bytes)
}

func readJWTAssertion(filename string) (map[string]any, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b := bufio.NewReader(file)
	raw, err := io.ReadAll(b)
	if err != nil {
		return nil, err
	}

	out := map[string]any{}
	err = json.Unmarshal(raw, &out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func createToken(claims map[string]any, key any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims(claims))

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func main() {
	keyFile := flag.String("private-ec-key-file", "", "path to the private key file")
	assertionFile := flag.String("assertion-file", "", "path to the JWT assertion file")
	flag.Parse()

	if keyFile == nil || *keyFile == "" {
		panic("missing required flag: -private-ec-key-file")
	}
	if assertionFile == nil || *assertionFile == "" {
		panic("missing required flag: -assertion-file")
	}

	key, err := readPrivateKey(*keyFile)
	if err != nil {
		panic(err)
	}

	// read JWT assertion
	jwt, err := readJWTAssertion(*assertionFile)
	if err != nil {
		panic(err)
	}

	// sign JWT assertion
	token, err := createToken(jwt, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
