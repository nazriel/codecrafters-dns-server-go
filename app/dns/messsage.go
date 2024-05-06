package dns

type Message struct {
	Header    Header
	Questions []Question
	Answers   []Answer
}

func (message Message) Bytes() []byte {
	bytes := message.Header.Bytes()
	for _, q := range message.Questions {
		bytes = append(bytes, q.Bytes()...)
	}
	for _, a := range message.Answers {
		bytes = append(bytes, a.Bytes()...)
	}

	return bytes
}
