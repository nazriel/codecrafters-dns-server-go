package dns

type Message struct {
	Header
}

func (message Message) Bytes() []byte {
	bytes := message.Header.Bytes()
	return bytes
}
