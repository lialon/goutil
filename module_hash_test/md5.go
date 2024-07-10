package ha

import (
	"crypto/md5"
	"fmt"
)

func Test1() {
	block := md5.New()
	fmt.Println(block.Sum([]byte("1234")))
}

func Test2() {
	fmt.Println("pl")
}
