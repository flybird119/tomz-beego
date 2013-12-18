package main

import (
	"fmt"
	"time"
)

func main() {

	if err := test(); err != nil {
		fmt.Println(err)

	}
}

type BaseError struct {
	When      time.Time
	ErrorCode string
	Message   string
}

func (e BaseError) Error() string {
	return fmt.Sprintf("Time:\t%v\nErrorCode:\t%v\nMessage:\t%v", e.When, e.ErrorCode, e.Message)
}

func test() error {
	// return BaseError{When: time.Date(1990, 1, 10, 10, 10, 10, 0, time.UTC), ErrorCode: "101", Message: "testksjdfkjsdlfjsldjkfsd"}
	return BaseError{When: time.Now(), ErrorCode: "101", Message: "testksjdfkjsdlfjsldjkfsd"}
}
