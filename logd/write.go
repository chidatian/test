package main

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

	for i := 0; i < 10; i++{
		s := fmt.Sprintf("this is a log log test---------------%d\nasdf%d\n", i, i)
		_, err := this.File.WriteString(s)
		if err != nil {
			log.Println(err)
		}
		// fmt.Printf("---------%d---------\n", n)
		time.Sleep(time.Second)
	}
}

func main() {
	e := new(Echo)
	e.Filename = "./errors.log"
	e.Write()
}