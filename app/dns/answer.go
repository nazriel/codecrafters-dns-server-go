package dns

import "encoding/binary"

type Answer struct {
	Name   Name
	Type   uint16
	Class  uint16
	TTL    uint32
	Length uint16
	Data   uint32
}

func (a Answer) Bytes() []byte {
	buf := a.Name.AsLabel()
	buf = binary.BigEndian.AppendUint16(buf, a.Type)
	buf = binary.BigEndian.AppendUint16(buf, a.Class)
	buf = binary.BigEndian.AppendUint32(buf, a.TTL)
	buf = binary.BigEndian.AppendUint16(buf, a.Length)
	buf = binary.BigEndian.AppendUint32(buf, a.Data)
	return buf
}
