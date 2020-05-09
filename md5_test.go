/*
# @Author ww
# @Time 2019/8/26 11:01
# @File md5_test.go.go
*/
package src

import (
	"testing"
)

func TestMd5(t *testing.T) {
	if Md5("Hello gopher!") == "8eceef269f2eca417f90d763565d22ff" {
		t.Log()
	} else {
		t.Fatal()
	}
}

func TestMd5File(t *testing.T) {
	str, err := Md5File("./test.txt")
	if err != nil {
		t.Fatal(err)
	} else if str == "8eceef269f2eca417f90d763565d22ff" {
		t.Log()
	} else {
		t.Fatal()
	}
}
