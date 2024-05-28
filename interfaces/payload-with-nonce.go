package interfaces

type payloadWithNonce interface {
	SetNonce(nonce uint64)
}

type _payloadWithNonce struct {
	Nonce uint64
}
