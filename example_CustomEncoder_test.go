package esm_test

import (
	"fmt"

	"../esm"
)

type customStruct struct {
	S string
	N int
}

var _ esm.CustomEncoder = (*customStruct)(nil)
var _ esm.CustomDecoder = (*customStruct)(nil)

func (s *customStruct) EncodeEsm(enc *esm.Encoder) error {
	return enc.Encode(s.S, s.N)
}

func (s *customStruct) DecodeEsm(dec *esm.Decoder) error {
	return dec.Decode(&s.S, &s.N)
}

func ExampleCustomEncoder() {
	b, err := esm.Marshal(&customStruct{S: "hello", N: 42})
	if err != nil {
		panic(err)
	}

	var v customStruct
	err = esm.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", v)
	// Output: esm_test.customStruct{S:"hello", N:42}
}
