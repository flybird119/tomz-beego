package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argLen := len(os.Args)

	if argLen < 2 {
		fmt.Printf("Plase give the path")
		return
	}

	path := os.Args[1]
	var mode os.FileMode = 777
	if argLen > 2 {
		modeInt, _ := strconv.ParseUint(os.Args[2], 10, 64)
		mode = os.FileMode(modeInt)
	}

	if e := os.Mkdir(path, mode); e != nil {
		fmt.Println(e)
	}
}
