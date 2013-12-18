package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
//	"crypto/tls"
	"errors"
	"fmt"
)

var decrypted string

func main() {
	var data []byte
	var err error
	data, err = RsaEncrypt([]byte("polaris@studygolang.com"))
	if err != nil {
		panic(err)
	}

	fmt.Println("rsa encrypt base64:" + base64.StdEncoding.EncodeToString(data))
	origData, err := RsaDecrypt(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("string: ",string(origData))

//	block, b2 := pem.Decode(publicKey)
//	println(block.Type,b2)

//	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
//	println(privateKey)

}

var publicKey = []byte(`
-----BEGIN CERTIFICATE-----
MIIDhzCCAm+gAwIBAgIJAMYy9hClC02WMA0GCSqGSIb3DQEBBQUAMFoxCzAJBgNV
BAYTAkNOMQ8wDQYDVQQIDAZIZWliZWkxEDAOBgNVBAcMB0JlaWppbmcxDjAMBgNV
BAoMBXptIGhsMQswCQYDVQQLDAJ6bTELMAkGA1UEAwwCem0wHhcNMTMxMjA1MTAw
ODMxWhcNMTQxMjA1MTAwODMxWjBaMQswCQYDVQQGEwJDTjEPMA0GA1UECAwGSGVp
YmVpMRAwDgYDVQQHDAdCZWlqaW5nMQ4wDAYDVQQKDAV6bSBobDELMAkGA1UECwwC
em0xCzAJBgNVBAMMAnptMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA
5MbEFDSspkYM65jj8KCH33JJaAf6N2SudKN99k1+CRkVod2yyDy+lROsi8LuSLuh
hsWJQ5RDNhHdbo9RhnrIfAp25D598XbrV1YtijzlijUWGPg81hFpgml4JHKuJUCh
OPXcAoh6Hp6KveHnzRuQNa8PM7FBNgjuOh5qn13lFPEOSiLPSAVnh+7zM0ZaZ3eX
Ichpjz3lpcvgOj21VnYfyzeT1pJMUtyV6F6kdliMl8hW629C+8e/Fu7bMbrOXzRx
nXu2eYLJXU9Hv2fUa991AlW8Sn/epZ38ez2qrVq5ymc+nUWcCtjA3W0Tyzd4nDxd
LK/oilGTuA2YK2r/lnMIUQIDAQABo1AwTjAdBgNVHQ4EFgQUUD9xr07VhMni/h5h
3l2dW4PAJcswHwYDVR0jBBgwFoAUUD9xr07VhMni/h5h3l2dW4PAJcswDAYDVR0T
BAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEA3D+PVcp8Aps1AQ8c3grIel5g1xYI
XkIaK9sD3spJuFEXXYHni/P3pQ4vvppacB+eoK1jPYPtm0buK//NTB6lVHh0Az3d
kIba6VMBu6glEsl31Hm68mAyGDG1ysMpofRp2PZvgstD5bWKtDHkL7rNtqRGvr+o
04CIdNyh0/yKzIQEVkGXdt0OIPQYcUzRDKqV5l3K5078n6Rx+9pm3DMXDxLTQTD+
I29LsvEGL8YyvBrPNsqDIdom1l4co6VEqitzUZXAMtfwkOXkkLuzj0zLAuhvgg3s
dJJVfQWJKjVnq5dzH3MnnxYbZ5bAabiGUqFq6E4bgaM7iOmYn7sVC6911A==
-----END CERTIFICATE-----
`)

//var publicKey = []byte(`
//-----BEGIN PUBLIC KEY-----
//MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
//ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
//wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
//AUeJ6PeW+DAkmJWF6QIDAQAB
//-----END PUBLIC KEY-----
//`)

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)


// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

