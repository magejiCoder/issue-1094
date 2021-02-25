package main

import "fmt"

func main() {
	f := Foo()
	fmt.Printf("f.bat: %s", f.bat)
}

type T = struct {
	bar bool
	bat string
}

func Foo() *T {
	return &foo{
		bar: false,
		bat: "issue-1094",
	}
}

type foo = struct {
	bar bool
	bat string
}