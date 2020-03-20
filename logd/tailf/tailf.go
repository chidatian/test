package tailf

import (
	"log"
	"os"
	"bufio"
	// "fmt"
)

type Tailf struct {
	Filename string
	File *os.File
	FileInfo os.FileInfo
	Reader *bufio.Reader
	ModTime int64
}

func (this *Tailf) Seek(offset int64) int64 {
	ret, err := this.File.Seek(offset, 0)
	if err != nil {
		log.Println(err)
	}

	this.Reader = bufio.NewReader(this.File)
	return ret
}

func (this *Tailf) ReadBytes() ( []byte, error ) {
	line, err := this.Reader.ReadBytes('\n')
	return line, err
}

func (this *Tailf) OpenFile() bool {
	var err error
	this.File, err = os.Open(this.Filename)
	if err != nil {
		log.Println(err)
		return false
	}
	this.Stat()
	this.ModTime = this.FileInfo.ModTime().Unix()
	return true
}

func (this *Tailf) IsModTime() bool {
	this.Stat()
	t := this.FileInfo.ModTime().Unix()
	if this.ModTime == t {
		return false
	}
	this.ModTime = t
	return true
}

func (this *Tailf) Stat() {
	var err error
	this.FileInfo, err = this.File.Stat()
	if err != nil {
		log.Println(err)
	}
}