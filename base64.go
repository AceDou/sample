package sample

import "encoding/base64"

const (
	base64Table = "123QRSTUabcdVWXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func base64StdEncode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64StdDecode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}
