package main

import (
	f "fmt"
	os "os"
	run "runtime"
)

const tool string = "OS"

func box() string {
	x, err := os.Hostname()
	if err != nil {
		f.Println(err)
		os.Exit(0)
	}
	return x
}

func info() string {
	y := run.GOOS
	return y
}

func main() {

	f.Printf("%s = %s", tool, info())
	f.Println("")
	f.Printf("Hostname = %s", box())

}
