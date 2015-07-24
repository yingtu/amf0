package amf0

import (
	"io"
)

type NullType struct{}

var _ AmfType = &NullType{}

func NewNull() *NullType {
	return &NullType{}
}

// Decode implements AmfType.Decode
func (n *NullType) Decode(r io.Reader) error {
	return nil
}

// DecodeFrom implements AmfType.DecodeFrom
func (n *NullType) DecodeFrom(slice []byte, pos int) (int, error) {
	return 0, nil
}

// Encode implements AmfType.Encode
func (n *NullType) Encode(w io.Writer) (int, error) {
	return w.Write(n.EncodeBytes())
}

// EncodeTo implements AmfType.EncodeTo
func (n *NullType) EncodeTo(slice []byte, pos int) {
	slice[pos] = MARKER_NULL
}

// EncodeBytes implements AmfType.EncodeBytes
func (n *NullType) EncodeBytes() []byte {
	return []byte{MARKER_NULL}
}

// Implements AmfType.Marker
func (n *NullType) Marker() byte {
	return MARKER_NULL
}
