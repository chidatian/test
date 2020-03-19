package config

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"bytes"
)

type Config struct {
	Filename string
	Context map[string]string
}

func (this *Config) handler(line []byte) {
	line = bytes.TrimSpace(line)
	if len(line) > 1 {
		if line[0] != '#' {
			ret := bytes.Split(line, []byte(" "))
			this.Context[string(bytes.TrimSpace(ret[0]))] = string(bytes.TrimSpace(ret[1]))
		}
	}
}

func (this *Config) LoadConfig() {
	this.Context = make(map[string]string, 3)
	file, err := os.Open(this.Filename)
	if err != nil {
		log.Println(err)
	}
	fileInfo, _ := file.Stat()
	reader := bufio.NewReader(file)
	var p int64 = 0
	for {
		_, err := file.Seek(p, 0)
		line, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			break
		}
		this.handler(line)
		p += int64(len(line))
		if p >= fileInfo.Size() {
			break
		}
	}
}

func Run() {
	c := new(Config)
	c.Filename = "./logd.conf"
	c.LoadConfig()
	fmt.Println(c.Context["position"])
	fmt.Println("---")
}