package src

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

// 字符串SHA1
func Sha1(str string) string {
	sha1Ctx := sha1.New()
	sha1Ctx.Write([]byte(str))
	data := sha1Ctx.Sum(nil)
	return hex.EncodeToString(data[:])
}

// 文件SHA1
func Sha1File(filePath string) (string, error) {
	f, err := os.Open("test.txt")
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	data := h.Sum(nil)
	return hex.EncodeToString(data[:]), nil
}
