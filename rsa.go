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

//RSA公钥私钥产生
func GenRsaKey(bits int, private_pem string, public_pem string) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}

	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}

	file, err := os.Create(private_pem)
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}

	file, err = os.Create(public_pem)
	if err != nil {
		return err
	}

	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	return nil
}

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
