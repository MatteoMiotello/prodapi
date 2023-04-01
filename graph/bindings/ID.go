package bindings

import "encoding/binary"

type ID int64

func UnmarshalID(b []byte, v *ID) error {
	*v = ID(binary.BigEndian.Uint64(b))

	return nil
}

func MarshalID(v *ID) ([]byte, error) {
	byte := new([]byte)

	return binary.BigEndian.AppendUint64(*byte, uint64(*v)), nil
}
