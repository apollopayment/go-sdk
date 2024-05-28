package noncer

import (
	"math/rand"
	"sync"
)

type Noncer struct {
	mx *sync.Mutex

	nonce uint32
}

func New() *Noncer {
	initialNonce := rand.Uint32()

	return &Noncer{
		mx:    &sync.Mutex{},
		nonce: initialNonce,
	}
}

func (n *Noncer) GetNonce() uint32 {
	n.mx.Lock()
	defer n.mx.Unlock()

	n.nonce = n.nonce + 1

	return n.nonce
}
