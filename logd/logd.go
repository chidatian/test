package main

import (
	"fmt"
	"log"
	"time"
	"logd/tailf"
	"logd/config"
	"regexp"
)

type logd struct {
	MsgChan chan *Message
}

type Message struct {
	Src []byte
	Ip []byte
	Datetime []byte
	Server []byte
	Request []byte
}

func New() *logd {
	return &logd {
		MsgChan : make(chan *Message, 10),
	}
}

func (this *logd) ListenMsgChan() {
	for {
		select {
		case item := <- this.MsgChan : 
			go func(m *Message) {
				exp := regexp.MustCompile("([0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2})|(client: [0-9]{1,3}[.][0-9]{1,3}[.][0-9]{1,3}[.][0-9]{1,3})|(server: [^,]*)|(request: \"{1}[^\"]*\"{1})")
				ret := exp.FindAll(m.Src, 4)
				m.Datetime, m.Ip, m.Server, m.Request = ret[0],ret[1],ret[2],ret[3]
				fmt.Printf("%s---%s---%s---%s\n",string(m.Ip[8:]),string(m.Datetime),string(m.Server),string(m.Request))
			}(item)
		}
	}
}

var conf *config.Config
var cache *config.Cache

func main() {
	logd := New()
	go logd.ListenMsgChan()

	tail := new(tailf.Tailf)
	tail.Filename = conf.Get("path")
	b := tail.OpenFile()
	var n, p int64 = 0, cache.Position()
	if tail.FileInfo.Size() < p {
		p = tail.FileInfo.Size()
	}
	if b {
		for {
			if tail.IsModTime() {
				for {
					tail.Seek(p)
					res, err := tail.ReadBytes()
					if err != nil {
						log.Println(err)
					}
					n = int64(len(res))
					p += n
					cache.Write(p)
					logd.MsgChan <- &Message{Src:res}
					if p >= tail.FileInfo.Size() {
						break
					}
				} // end for
			}
			time.Sleep(time.Second)
		} // end for
	}
}

func init() {
	conf = config.Load("./config/logd.conf")
	cache = config.LoadCache("./config/cache")
}