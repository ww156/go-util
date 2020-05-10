package util

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	str := "123"
	//c04483ac5c934f62569c0fcb6fd75a1dd6b4fe5a
	fmt.Println(Sha1(str))
}

func TestSha1File(t *testing.T) {
	file := "./test.txt"
	//c04483ac5c934f62569c0fcb6fd75a1dd6b4fe5a
	fmt.Println(Sha1File(file))
}
