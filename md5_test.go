package sample

import (
	"log"
	"testing"
)

func TestMd5_SumLowercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Sum_Lowercase32(testData)
	expectData := "d1d03d8643963150f05a5f9ffee0c449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}

func TestMd5_SumUppercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Sum_Uppercase32(testData)
	expectData := "D1D03D8643963150F05A5F9FFEE0C449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}

func TestMd5_EncodeLowercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Encode_Lowercase32(testData)
	expectData := "d1d03d8643963150f05a5f9ffee0c449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}

func TestMd5_EncodeUppercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Encode_Uppercase32(testData)
	expectData := "D1D03D8643963150F05A5F9FFEE0C449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}

func TestMd5_BeegoLowercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Beego_Lowercase32(testData)
	expectData := "d1d03d8643963150f05a5f9ffee0c449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}

func TestMd5_BeegoUppercase32(t *testing.T) {
	// encode
	testData := "github.com/niugo/sample"
	codeData := MD5Beego_Uppercase32(testData)
	expectData := "D1D03D8643963150F05A5F9FFEE0C449"
	if expectData != string(codeData) {
		log.Println(expectData + " is not equal to " + string(codeData))
	}
}
