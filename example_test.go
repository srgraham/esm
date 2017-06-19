package esm_test

import (
	"bytes"
	"fmt"

	"../esm"
)

func ExampleMarshal() {
	type Item struct {
		Foo string
	}

	b, err := esm.Marshal(&Item{Foo: "bar"})
	if err != nil {
		panic(err)
	}

	var item Item
	err = esm.Unmarshal(b, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}

func ExampleMarshal_mapStringInterface() {
	in := map[string]interface{}{"foo": 1, "hello": "world"}
	b, err := esm.Marshal(in)
	if err != nil {
		panic(err)
	}

	var out map[string]interface{}
	err = esm.Unmarshal(b, &out)
	if err != nil {
		panic(err)
	}

	fmt.Println("foo =", out["foo"])
	fmt.Println("hello =", out["hello"])

	// Output:
	// foo = 1
	// hello = world
}

func ExampleDecoder_SetDecodeMapFunc() {
	buf := new(bytes.Buffer)

	enc := esm.NewEncoder(buf)
	in := map[string]string{"hello": "world"}
	err := enc.Encode(in)
	if err != nil {
		panic(err)
	}

	dec := esm.NewDecoder(buf)
	dec.SetDecodeMapFunc(func(d *esm.Decoder) (interface{}, error) {
		n, err := d.DecodeMapLen()
		if err != nil {
			return nil, err
		}

		m := make(map[string]string, n)
		for i := 0; i < n; i++ {
			mk, err := d.DecodeString()
			if err != nil {
				return nil, err
			}

			mv, err := d.DecodeString()
			if err != nil {
				return nil, err
			}

			m[mk] = mv
		}
		return m, nil
	})

	out, err := dec.DecodeInterface()
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: map[hello:world]
}

func ExampleDecoder_Query() {
	b, err := esm.Marshal([]map[string]interface{}{
		{"id": 1, "attrs": map[string]interface{}{"phone": 12345}},
		{"id": 2, "attrs": map[string]interface{}{"phone": 54321}},
	})
	if err != nil {
		panic(err)
	}

	dec := esm.NewDecoder(bytes.NewBuffer(b))
	values, err := dec.Query("*.attrs.phone")
	if err != nil {
		panic(err)
	}
	fmt.Println("phones are", values)

	dec.Reset(bytes.NewBuffer(b))
	values, err = dec.Query("1.attrs.phone")
	if err != nil {
		panic(err)
	}
	fmt.Println("2nd phone is", values[0])
	// Output: phones are [12345 54321]
	// 2nd phone is 54321
}

func ExampleEncoder_StructAsArray() {
	type Item struct {
		Foo string
		Bar string
	}

	var buf bytes.Buffer
	enc := esm.NewEncoder(&buf).StructAsArray(true)
	err := enc.Encode(&Item{Foo: "foo", Bar: "bar"})
	if err != nil {
		panic(err)
	}

	dec := esm.NewDecoder(&buf)
	v, err := dec.DecodeInterface()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	// Output: [foo bar]
}

func ExampleMarshal_asArray() {
	type Item struct {
		_esm struct{} `esm:",asArray"`
		Foo      string
		Bar      string
	}

	var buf bytes.Buffer
	enc := esm.NewEncoder(&buf)
	err := enc.Encode(&Item{Foo: "foo", Bar: "bar"})
	if err != nil {
		panic(err)
	}

	dec := esm.NewDecoder(&buf)
	v, err := dec.DecodeInterface()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
	// Output: [foo bar]
}
