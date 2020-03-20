package tailf

import (
	"os"
	"log"
	"fmt"
	"time"
)

type Echo struct {
	Filename string
	File *os.File
}

func (this *Echo) Write() {
	var err error
	this.File, err = os.OpenFile(this.Filename, os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < 100; i++{
		s := fmt.Sprintf("123456789%d\nasdf%d\n", i, i)
		_, err := this.File.WriteString(s)
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("---------%d---------\n", n)
		time.Sleep(time.Second)
	}
}