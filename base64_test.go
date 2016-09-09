package sample

import (
	"log"
	"testing"
)

func TestBase64(t *testing.T) {
	// encode
	testData := "hello world"
	debyte := base64Encode([]byte(testData))

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		log.Fatal(err.Error())
	}

	if testData != string(enbyte) {
		log.Fatal(testData + " is not equal to " + string(enbyte))
		t.Fail()
	}
}

func TestBase64_2(t *testing.T) {
	// encode
	testData := "abc123!?$*&()'-=@~"
	debyte := base64Encode([]byte(testData))

	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		log.Fatal(err.Error())
	}

	if testData != string(enbyte) {
		//log.Fatal(testData + " is not equal to " + string(enbyte))
		log.Println(testData + " is not equal to " + string(enbyte))
	}
}

func TestBase64_Std(t *testing.T) {
	// encode
	testData := "hello world"
	debyte := base64StdEncode([]byte(testData))

	// decode
	enbyte, err := base64StdDecode(debyte)
	if err != nil {
		log.Fatal(err.Error())
	}

	if testData != string(enbyte) {
		log.Fatal(testData + " is not equal to " + string(enbyte))
		//t.Fail()
	}
}

func TestBase64_Std_2(t *testing.T) {
	// encode
	testData := "abc123!?$*&()'-=@~"
	debyte := base64StdEncode([]byte(testData))

	// decode
	enbyte, err := base64StdDecode(debyte)
	if err != nil {
		log.Fatal(err.Error())
	}

	if testData != string(enbyte) {
		log.Fatal(testData + " is not equal to " + string(enbyte))
		//t.Fail()
	}
}

func TestBase64_Compare(t *testing.T) {
	// encode
	testData := "hello world"
	debyte := base64Encode([]byte(testData))

	debyte1 := base64StdEncode([]byte(testData))

	if string(debyte) != string(debyte1) {
		log.Println(string(debyte) + " is not equal to " + string(debyte1))
	}
}
