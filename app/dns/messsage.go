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

func MessageFromBytes(payload []byte) *Message {
	message := &Message{}
	payloadLen := len(payload)

	if payloadLen < 12 {
		return nil
	}

	if header := HeaderFromBytes(payload[0:12]); header != nil {
		message.Header = *header
	} else {
		return nil
	}
	payload = payload[12:]

	questionCount := message.Header.QuestionCount
	if questions, consumed := QuestionsFromBytes(payload[:payloadLen], questionCount); questions != nil {
		message.Questions = questions
		payload = payload[consumed:]
	} else {
		return nil
	}

	_ = payload

	return message
}
