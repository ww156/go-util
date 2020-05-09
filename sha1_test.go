package util

import (
	"fmt"
	"testing"
)

func TestSha1(t *testing.T) {
	//c04483ac5c934f62569c0fcb6fd75a1dd6b4fe5a
	fmt.Println(Sha1(STR))
}

func TestSha1File(t *testing.T) {
	//c04483ac5c934f62569c0fcb6fd75a1dd6b4fe5a
	fmt.Println(Sha1File(FILE))
}
