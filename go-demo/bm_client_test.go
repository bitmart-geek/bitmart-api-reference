// Copyright (c) 2018 BitMart Exchange. All Rights Reserved.
//

package client

import (
	"github.com/stretchr/testify/assert"

	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"
)

var api_key = "YOUR_TEST_API_KEY"
var api_secret = "YOUR_TEST_API_SECRET"
var rsa_private_key = []byte(`-----BEGIN RSA PRIVATE KEY-----
YOUR_TEST_RSA_PRIVATE_KEY
-----END RSA PRIVATE KEY-----`)

func PrivateEncrypt(priv *rsa.PrivateKey, data []byte) (enc []byte, err error) {

	k := (priv.N.BitLen() + 7) / 8
	tLen := len(data)
	// rfc2313, section 8:
	// The length of the data D shall not be more than k-11 octets
	if tLen > k-11 {
		err = fmt.Errorf("err input size")
		return
	}
	em := make([]byte, k)
	em[1] = 1
	for i := 2; i < k-tLen-1; i++ {
		em[i] = 0xff
	}
	copy(em[k-tLen:k], data)
	c := new(big.Int).SetBytes(em)
	if c.Cmp(priv.N) > 0 {
		err = fmt.Errorf("err encryption")
		return
	}
	var m *big.Int
	var ir *big.Int
	if priv.Precomputed.Dp == nil {
		m = new(big.Int).Exp(c, priv.D, priv.N)
	} else {
		// We have the precalculated values needed for the CRT.
		m = new(big.Int).Exp(c, priv.Precomputed.Dp, priv.Primes[0])
		m2 := new(big.Int).Exp(c, priv.Precomputed.Dq, priv.Primes[1])
		m.Sub(m, m2)
		if m.Sign() < 0 {
			m.Add(m, priv.Primes[0])
		}
		m.Mul(m, priv.Precomputed.Qinv)
		m.Mod(m, priv.Primes[0])
		m.Mul(m, priv.Primes[1])
		m.Add(m, m2)

		for i, values := range priv.Precomputed.CRTValues {
			prime := priv.Primes[2+i]
			m2.Exp(c, values.Exp, prime)
			m2.Sub(m2, m)
			m2.Mul(m2, values.Coeff)
			m2.Mod(m2, prime)
			if m2.Sign() < 0 {
				m2.Add(m2, prime)
			}
			m2.Mul(m2, values.R)
			m.Add(m, m2)
		}
	}

	if ir != nil {
		m.Mul(m, ir)
		m.Mod(m, priv.N)
	}
	enc = m.Bytes()
	return
}

func TestRSASign(t *testing.T) {
	c := NewBMClient(api_key, api_secret, string(rsa_private_key))
	block, _ := pem.Decode(rsa_private_key)
	assert.NotNil(t, block)
	assert.Equal(t, block.Type, "RSA PRIVATE KEY")
	rsaPrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	assert.Nil(t, err)
	signContent := c.apiKey + ":" + c.apiSecret + ":" + strconv.FormatInt(time.Now().UTC().Unix(), 10)

	encData, err := PrivateEncrypt(rsaPrivateKey.(*rsa.PrivateKey), []byte(signContent))
	privateEncryptData := base64.StdEncoding.EncodeToString(encData)

	rsaSignedData, err := c.signWithPrivateKey(signContent, c.rsaPrivateKey)
	assert.Equal(t, privateEncryptData, rsaSignedData)

}

func TestBMAccessToken(t *testing.T) {
	c := NewBMClient(api_key, api_secret, string(rsa_private_key))
	err := c.getAuthToken()
	assert.Nil(t, err)
}
