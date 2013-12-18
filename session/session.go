package main

import (
	"time"
)

type SessionManager struct {
	Lock    sync.Mutex
	Expires int
}

type Session struct {
	userId      int64
	userName    string
	userCode    string
	password    string
	language    string
	requestTime time.Time
}
