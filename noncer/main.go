package noncer

import (
	"sync"
	"time"
)

type Noncer struct {
	mx *sync.Mutex

	nonce int64
}

func New() *Noncer {
	initialNonce := time.Now().UnixMilli()

	return &Noncer{
		mx:    &sync.Mutex{},
		nonce: initialNonce,
	}
}

func (n *Noncer) GetNonce() int64 {
	n.mx.Lock()
	defer n.mx.Unlock()

	n.nonce = n.nonce + 1

	return n.nonce
}
