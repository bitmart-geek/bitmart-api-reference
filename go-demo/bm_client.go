// Copyright (c) 2018 BitMart Exchange. All Rights Reserved.
//

package client

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	baseUrl  = "https://api.bitmart.com/v2"
	tokenUrl = baseUrl + "/token"
)

type BMClient struct {
	apiKey    string
	apiSecret string

	// RSA private key with PKCS8 format
	// Example:
	//
	// "
	// -----BEGIN RSA PRIVATE KEY-----
	//	MIIE****
	// -----END RSA PRIVATE KEY-----
	// "
	rsaPrivateKey string
	accessToken   string
}

type accessToken struct {
	AccessToken string `json:"access_token"`
}

func NewBMClient(apikey, apisecret, privateKey string) *BMClient {
	return &BMClient{
		apiKey:        apikey,
		apiSecret:     apisecret,
		rsaPrivateKey: privateKey,
	}
}
func (c *BMClient) getSignContent() string {
	return c.apiKey + ":" + c.apiSecret + ":" + strconv.FormatInt(time.Now().UTC().Unix()*1000, 10)
}

func (c *BMClient) signWithPrivateKey(message, privateKey string) (string, error) {

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", fmt.Errorf("no key found")
	}

	if block.Type != "RSA PRIVATE KEY" {
		return "", fmt.Errorf("unsupported key type %q", block.Type)
	}

	rsaPrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Sign secret with rsa with PKCS 1.5 as the padding algorithm
	// The result should be exactly same as "openssl rsautl -sign -inkey "YOUR_RSA_PRIVATE_KEY" -in "YOUR_PLAIN_TEXT""
	signer, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey.(*rsa.PrivateKey), crypto.Hash(0), []byte(message))
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signer), nil
}

func (c *BMClient) getAuthToken() error {
	clientSecret, err := c.signWithPrivateKey(c.getSignContent(), c.rsaPrivateKey)
	if err != nil {
		log.Printf("Fail to sign client secret. Error: %+v \n", err)
		return err
	}

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", c.apiKey)
	data.Set("client_secret", clientSecret)

	req, err := http.NewRequest("POST", tokenUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Printf("Fail to create http request. Error: %+v \n", err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Fail to do http call. Error: %+v \n", err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Printf("Fail to read response body for access token. Error: %+v \n", err)
		return err
	}

	token := accessToken{}
	if err := json.Unmarshal(body, &token); err != nil {
		log.Printf("Fail to unmarshal json for access token. Error: %+v \n", err)
		return err
	}

	if token.AccessToken == "" {
		log.Println("Error: fail to get access token.")
		return fmt.Errorf("fail to get access token")
	}

	fmt.Printf("This is just for test, please delete this line. Your access token: %s\n", token.AccessToken)
	c.accessToken = token.AccessToken
	return nil
}
