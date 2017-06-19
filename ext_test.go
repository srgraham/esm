package esm_test

import (
	"reflect"
	"testing"

	"../esm"
	"../esm/codes"
)

func init() {
	esm.RegisterExt(9, (*ExtTest)(nil))
}

func TestRegisterExtPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("panic expected")
		}
		got := r.(error).Error()
		wanted := "esm: ext with id=9 is already registered"
		if got != wanted {
			t.Fatalf("got %q, wanted %q", got, wanted)
		}
	}()
	esm.RegisterExt(9, (*ExtTest)(nil))
}

type ExtTest struct {
	S string
}

var _ esm.CustomEncoder = (*ExtTest)(nil)
var _ esm.CustomDecoder = (*ExtTest)(nil)

func (ext ExtTest) EncodeEsm(e *esm.Encoder) error {
	return e.EncodeString("hello " + ext.S)
}

func (ext *ExtTest) DecodeEsm(d *esm.Decoder) error {
	var err error
	ext.S, err = d.DecodeString()
	return err
}

func TestExt(t *testing.T) {
	for _, v := range []interface{}{ExtTest{"world"}, &ExtTest{"world"}} {
		b, err := esm.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}

		var dst interface{}
		err = esm.Unmarshal(b, &dst)
		if err != nil {
			t.Fatal(err)
		}

		v, ok := dst.(ExtTest)
		if !ok {
			t.Fatalf("got %#v, wanted ExtTest", dst)
		}

		wanted := "hello world"
		if v.S != wanted {
			t.Fatalf("got %q, wanted %q", v.S, wanted)
		}

		var ext ExtTest
		err = esm.Unmarshal(b, &ext)
		if err != nil {
			t.Fatal(err)
		}
		if ext.S != wanted {
			t.Fatalf("got %q, wanted %q", ext.S, wanted)
		}
	}
}

func TestUnknownExt(t *testing.T) {
	b := []byte{codes.FixExt1, 1, 0}

	var dst interface{}
	err := esm.Unmarshal(b, &dst)
	if err == nil {
		t.Fatalf("got nil, wanted error")
	}
	got := err.Error()
	wanted := "esm: unregistered ext id=1"
	if got != wanted {
		t.Fatalf("got %q, wanted %q", got, wanted)
	}
}

func TestDecodeExtWithMap(t *testing.T) {
	type S struct {
		I int
	}
	esm.RegisterExt(2, S{})

	b, err := esm.Marshal(&S{I: 42})
	if err != nil {
		t.Fatal(err)
	}

	var got map[string]interface{}
	if err := esm.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	wanted := map[string]interface{}{"I": uint64(42)}
	if !reflect.DeepEqual(got, wanted) {
		t.Fatalf("got %#v, but wanted %#v", got, wanted)
	}
}
