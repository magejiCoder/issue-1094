package main

import (
	"fmt"

	"github.com/magejiCoder/issue-1094/proto/foo"
)

func main() {
	fb := &foo.FooBar{
		Name: "bat",
	}
	fmt.Printf("f.bat: %s", fb.GetName())
}
