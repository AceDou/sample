package sample

import (
	"log"
	"testing"
)

func TestGenRsa(t *testing.T) {
	if err := GenRsaKey(1024, "private.pem", "public.pem"); err != nil {
		log.Fatal("private.pem/public.pem 生成失败")
	}
}

func TestRsa(t *testing.T) {
	srcRsaData := "douxiance@xiaoniu66.com"
	data, err := RsaEncrypt([]byte(srcRsaData))
	if err != nil {
		log.Println("err:", err)
	}

	if calcRsaData, err := RsaDecrypt(data); err != nil {
		log.Println(err)
	} else {
		if srcRsaData != string(calcRsaData) {
			log.Fatal(srcRsaData + " 不等于 " + string(calcRsaData))
		}
	}
}
