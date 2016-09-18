package sample

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	fPublicPem, errOpen := os.Open("./rsa_public.pem")
	if errOpen != nil {
		return nil, errOpen
	}
	defer fPublicPem.Close()

	bytePublicPem, errReadAll := ioutil.ReadAll(fPublicPem)
	if errReadAll != nil {
		return nil, errReadAll
	}

	block, _ := pem.Decode(bytePublicPem)
	pubInterface, errParse := x509.ParsePKIXPublicKey(block.Bytes)
	if errParse != nil {
		return nil, errParse
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pubInterface.(*rsa.PublicKey), origData)
}

// 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	fPrivatePem, errOpen := os.Open("./rsa_private.pem")
	if errOpen != nil {
		return nil, errOpen
	}
	defer fPrivatePem.Close()

	bytePrivatePem, errReadAll := ioutil.ReadAll(fPrivatePem)
	if errReadAll != nil {
		return nil, errReadAll
	}

	block, _ := pem.Decode(bytePrivatePem)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	priv, errParse := x509.ParsePKCS1PrivateKey(block.Bytes)
	if errParse != nil {
		return nil, errParse
	}

	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
