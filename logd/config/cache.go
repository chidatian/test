package config

import (
	"os"
	"io/ioutil"
	"log"
	"strconv"
	// "fmt"
)

type Cache struct {
	filename string
	file *os.File
	position int64
}

func (this *Cache) Position() int64 {
	return this.position
}

func (this *Cache) read() {
	file, err := os.OpenFile(this.filename, os.O_CREATE|os.O_RDONLY, 0666)
	if err != nil {
		log.Println(err)
	}
	ret, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	this.position, _ = strconv.ParseInt(string(ret), 2, 64)
	if file.Close() != nil {
		log.Println(err)
	}
}

func (this *Cache) open() {
	var err error
	this.file, err = os.OpenFile(this.filename, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Println(err)
	}
}

func (this *Cache) Write(p int64) {
	ret := strconv.FormatInt(p, 2)
	this.file.Seek(0, 0)
	_, err := this.file.WriteString(ret)
	if err != nil {
		log.Println(err)
	}
}

func LoadCache(file string) *Cache {
	c := new(Cache)
	c.filename = file
	c.read()
	c.open()
	return c
}