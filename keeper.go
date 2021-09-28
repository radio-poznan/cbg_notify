package cbg_notify

import (
	"fmt"
	"sync"
	"time"
)

type KeepItem struct {
	Term string
	Ts   time.Time
}

func (k *KeepItem) IsEqual(i KeepItem) bool {
	return  k.Term == i.Term && k.Ts == i.Ts
}

type ResendExec func (items []KeepItem)

type KeepToSend interface {
	Add(item KeepItem)
	Resend(timeout int, exec ResendExec)
	Remove(item KeepItem)
}

type keepToSend struct {
	mux   sync.Mutex
	Items []KeepItem
}

func NewKeepToSend() *keepToSend {
	return &keepToSend{Items: make([]KeepItem, 0)}
}

func (k *keepToSend) Add(item KeepItem) {
	k.mux.Lock()
	defer k.mux.Unlock()

	ctxStock := k.Items
	k.Items = append(ctxStock, item)
}

func (k *keepToSend) Remove(item KeepItem) {
	k.mux.Lock()
	defer k.mux.Unlock()

	max := len(k.Items)
	for index, val := range k.Items {
		if val.IsEqual(item) {
			removeAtIndex := index
			if removeAtIndex == 0 {
				k.Items = k.Items[1:]
			} else if removeAtIndex == max {
				k.Items = k.Items[:removeAtIndex-1]
			} else {
				k.Items = append(k.Items[:removeAtIndex-1], k.Items[removeAtIndex+1:]...)
			}
		}
	}
}


func (k *keepToSend) Resend(timeout int, exec ResendExec) {
	time.Sleep(time.Duration(timeout) * time.Minute)
	k.infoResend()
	exec(k.Items)
	k.Resend(timeout, exec)
}

func (k *keepToSend) infoResend() {
	fmt.Println("\nPamięć niezapisanych utworów: ", len(k.Items))
	for i, item :=range k.Items {
		fmt.Println(fmt.Sprintf("\n\t %d. %s @%d", i, item.Term, item.Ts))
	}
}