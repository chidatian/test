package main

import (
	"os"
	"log"
	// "fmt"
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

	for i := 0; i < 2; i++{
		// s := fmt.Sprintf("this is a log log test---------------%d\nasdf%d\n", i, i)
		s := `2020/03/22 12:37:41 [error] 131328#131400: *1 FastCGI sent in stderr: "PHP Parse error:  syntax error, unexpected '?' in E:\www\you\vendor\laravel\framework\src\Illuminate\Foundation\helpers.php on line 500" while reading response header from upstream, client: 127.0.0.1, server: local.you.com, request: "GET /test/vendor/chidatian HTTP/1.1", upstream: "fastcgi://127.0.0.1:9000", host: "local.you.com"`
		s += "\n"
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