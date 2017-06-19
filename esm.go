package esm

type Marshaler interface {
	MarshalEsm() ([]byte, error)
}

type Unmarshaler interface {
	UnmarshalEsm([]byte) error
}

type CustomEncoder interface {
	EncodeEsm(*Encoder) error
}

type CustomDecoder interface {
	DecodeEsm(*Decoder) error
}
