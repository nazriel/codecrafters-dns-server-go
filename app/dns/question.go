package dns

import (
	"encoding/binary"
)

type Question struct {
	Name  Name
	Type  uint16
	Class uint16
}

func (q Question) Bytes() []byte {
	buf := q.Name.AsLabel()
	buf = binary.BigEndian.AppendUint16(buf, q.Type)
	buf = binary.BigEndian.AppendUint16(buf, q.Class)
	return buf
}
