package dns

import (
	"encoding/binary"
)

type HeaderFlags struct {
	QueryResponseIndiciator uint8 // 1bit
	Opcode                  uint8 // 4bit
	AuthoritativeAnswer     uint8 // 1 bit
	Truncation              uint8 // 1 bit
	RecursionDesired        uint8 // 1 bit
	RecursionAvailable      uint8 // 1 bit
	Reserved                uint8 // 3 bit
	ResponseCode            uint8 // 4 bit
}

func UnpackFlags(input uint16) HeaderFlags {
	flags := HeaderFlags{}

	flags.QueryResponseIndiciator = uint8(input >> 15)
	flags.Opcode = uint8(input>>11) & 0b1111
	flags.AuthoritativeAnswer = uint8(input>>10) & 0b1
	flags.Truncation = uint8(input>>9) & 0b1
	flags.RecursionDesired = uint8(input>>8) & 0b1
	flags.RecursionAvailable = uint8(input>>7) & 0b1
	flags.Reserved = uint8(input>>4) & 0b111
	flags.ResponseCode = uint8(input) & 0b111

	return flags
}

func (hf HeaderFlags) Packed() uint16 {
	flags := uint16(hf.QueryResponseIndiciator) << 15
	flags += uint16(hf.Opcode) << 11
	flags += uint16(hf.AuthoritativeAnswer) << 10
	flags += uint16(hf.Truncation) << 9
	flags += uint16(hf.RecursionDesired) << 8
	flags += uint16(hf.RecursionAvailable) << 7
	flags += uint16(hf.Reserved) << 4
	flags += uint16(hf.ResponseCode)
	return flags
}

type Header struct {
	PacketID              uint16
	HeaderFlags           HeaderFlags
	QuestionCount         uint16
	AnswerRecordCount     uint16
	AuthorityRecordCount  uint16
	AdditionalRecordCount uint16
}

func HeaderFromBytes(payload []byte) *Header {
	header := &Header{}

	if len(payload) != 12 {
		return nil
	}

	header.PacketID = binary.BigEndian.Uint16(payload[0:2])
	header.HeaderFlags = UnpackFlags(binary.BigEndian.Uint16(payload[2:4]))
	header.QuestionCount = binary.BigEndian.Uint16(payload[4:6])
	header.AnswerRecordCount = binary.BigEndian.Uint16(payload[6:8])
	header.AuthorityRecordCount = binary.BigEndian.Uint16(payload[8:10])
	header.AdditionalRecordCount = binary.BigEndian.Uint16(payload[10:12])
	return header
}

func (header Header) Bytes() []byte {
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[:2], header.PacketID)
	binary.BigEndian.PutUint16(buf[2:4], header.HeaderFlags.Packed())
	binary.BigEndian.PutUint16(buf[4:6], header.QuestionCount)
	binary.BigEndian.PutUint16(buf[6:8], header.AnswerRecordCount)
	binary.BigEndian.PutUint16(buf[8:10], header.AuthorityRecordCount)
	binary.BigEndian.PutUint16(buf[10:], header.AdditionalRecordCount)

	return buf
}
