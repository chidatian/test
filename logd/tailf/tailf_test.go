package tailf

import (
	"testing"
	"fmt"
	"log"
	"time"
	// "io"
)

func TestTailf(t *testing.T) {
	e := new(Echo)
	e.Filename = "./errors.log"
	go e.Write()

	fmt.Println("-----------test")

	tail := new(Tailf)
	tail.Filename = "./errors.log"
	b := tail.OpenFile()

	var n, p int64 = 0, 0
	if b {
		for {
			if tail.IsModTime() {
				for {
					tail.Seek(p)
					res, err := tail.ReadBytes()
					// fmt.Println(err == io.EOF)
					if err != nil {
						log.Println(err)
					}
					n = int64(len(res))
					p += n
					fmt.Println(string(res))
					if p >= tail.FileInfo.Size() {
						break
					}
				} // for
			}
			time.Sleep(time.Second)
		} // for
	}
}
