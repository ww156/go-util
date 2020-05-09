package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

// Md5 字符串MD5
func Md5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	data := md5Ctx.Sum(nil)
	return hex.EncodeToString(data[:])
}

// Md5File 文件MD5
func Md5File(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	md5Ctx := md5.New()
	if _, err := io.Copy(md5Ctx, f); err != nil {
		return "", err
	}
	data := md5Ctx.Sum(nil)
	return hex.EncodeToString(data[:]), nil
}
