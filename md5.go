package sample

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

//MD5即Message-Digest Algorithm 5（信息-摘要算法5），用于确保信息传输完整一致。

// md5 加密的第一种方法
func MD5Sum_Lowercase32(data string) string {
	srcData := []byte(data)
	result := md5.Sum(srcData)
	return hex.EncodeToString(result[:])
}

func MD5Sum_Uppercase32(data string) string {
	return strings.ToUpper(MD5Sum_Lowercase32(data))
}

//md5 加密的第二种方法
func MD5Encode_Lowercase32(data string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(data))
	cipherStr := md5Ctx.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func MD5Encode_Uppercase32(data string) string {
	return strings.ToUpper(MD5Encode_Lowercase32(data))
}

// md5 加密的第三种方法(beego)
func MD5Beego_Lowercase32(data string) string {
	md5Ctx := md5.New()
	io.WriteString(md5Ctx, data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5Beego_Uppercase32(data string) string {
	return strings.ToUpper(MD5Beego_Lowercase32(data))
}
