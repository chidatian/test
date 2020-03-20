package config

import (
	"testing"
	"fmt"
)

// func TestConfig(t *testing.T) {
	// Load("./logd.conf")
// }

func TestCache(t *testing.T) {
	c := LoadCache("./cache")
	fmt.Println(c.Position())
}