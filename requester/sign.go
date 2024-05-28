package requester

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func (r *Requester) MakeSignFromStruct(payload interface{}) (string, []byte, error) {
	bytesPayload, err := json.Marshal(payload)

	if err != nil {
		return "", nil, err
	}

	sign := r.MakeSignFromBytes(bytesPayload)

	return sign, bytesPayload, nil
}

func (r *Requester) MakeSignFromBytes(payload []byte) string {
	h := hmac.New(sha256.New, []byte(r.secretKey))

	h.Write(payload)

	return hex.EncodeToString(h.Sum(nil))
}
