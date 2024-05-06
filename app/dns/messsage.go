package dns

type Message struct {
	Header   Header
	Question []Question
}

func (message Message) Bytes() []byte {
	bytes := message.Header.Bytes()
	for _, q := range message.Question {
		bytes = append(bytes, q.Bytes()...)
	}
	return bytes
}
