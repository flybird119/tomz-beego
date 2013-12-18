package main

import (
	"container/list"
	"fmt"
	"log"
	"sync"
	"time"
)

type SessionManager struct {
	Lock    sync.Mutex
	SL      *list.List
	Expires int
}
type Session struct {
	Key     string
	Expires int
	Value   interface{}
}

var SM *SessionManager

func (sm *SessionManager) NewSession(key string, value interface{}) *Session {
	return &Session{
		Key:     key,
		Expires: sm.Expires,
		Value:   value,
	}
}
func (sm *SessionManager) Listen() {
	time.AfterFunc(time.Second, func() { sm.Listen() })
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	for e := sm.SL.Front(); e != nil; e = e.Next() {
		if e.Value.(*Session).Expires == 0 {
			sm.SL.Remove(e)
		} else {
			e.Value.(*Session).Expires--
		}
	}
}
func (sm *SessionManager) Get(key string) interface{} {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	for e := sm.SL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Session).Key {
			e.Value.(*Session).Expires = sm.Expires
			sm.SL.MoveToBack(e)
			return e.Value.(*Session).Value
		}
	}
	return nil
}
func (sm *SessionManager) Set(key string, value interface{}) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	for e := sm.SL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Session).Key {
			sm.SL.Remove(e)
			return
		}
	}
	sm.SL.PushBack(sm.NewSession(key, value))
}
func (sm *SessionManager) Del(key string) {
	sm.Lock.Lock()
	defer sm.Lock.Unlock()
	for e := sm.SL.Front(); e != nil; e = e.Next() {
		if key == e.Value.(*Session).Key {
			sm.SL.Remove(e)
			return
		}
	}
	return
}

func main() {
	SM.Set("tsing", "hello")
	SM.Set("qing", "world")
	SM.Set("AppUserCode", "usercode")
	SM.Set("PassWord", "password")
	SM.Listen()

	time.Sleep(2 * time.Second)
	log.Println(SM.Get("tsing").(string))
	time.Sleep(1 * time.Second)
	log.Println(SM.SL.Len())
	time.Sleep(4 * time.Second)

	for e := SM.SL.Front(); e != nil; e = e.Next() {
		if e.Value.(*Session).Expires == 0 {
			SM.SL.Remove(e)
		} else {
			e.Value.(*Session).Expires--
		}
		time.Sleep(3 * time.Second)
		log.Println(e.Value.(*Session).Key)
	}

}

func init() {
	SM = &SessionManager{
		SL:      list.New(),
		Expires: 15,
	}
	SM.Listen()
}
