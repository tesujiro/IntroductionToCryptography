package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "tesujiro"

	h := sha1.New()

	// バイト配列を渡す
	h.Write([]byte(s))

	// 結果をバイト配列で得る
	bs := h.Sum(nil)

	// 16進数で表示
	fmt.Printf("%x\n", bs)
}
