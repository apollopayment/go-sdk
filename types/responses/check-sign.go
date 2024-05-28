package responses

type CheckSignResponse struct {
	Errors               []string `json:"errors"`
	CheckSignatureResult bool     `json:"checkSignatureResult"`
	Signature            string   `json:"signature"`
	ReceivedBody         string   `json:"receivedBody"`
}
