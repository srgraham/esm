# MessagePack encoding for Golang

[![Build Status](https://travis-ci.org/vmihailenco/esm.svg?branch=v2)](https://travis-ci.org/vmihailenco/esm)
[![GoDoc](https://godoc.org/../esm?status.svg)](https://godoc.org/../esm)

Supports:
- Primitives, arrays, maps, structs, time.Time and interface{}.
- Appengine *datastore.Key and datastore.Cursor.
- [CustomEncoder](https://godoc.org/../esm#example-CustomEncoder)/CustomDecoder interfaces for custom encoding.
- [Extensions](https://godoc.org/../esm#example-RegisterExt) to encode type information.
- Renaming fields via `esm:"my_field_name"`.
- Omitting empty fields via `esm:",omitempty"`.
- [Map keys sorting](https://godoc.org/../esm#Encoder.SortMapKeys).
- Encoding/decoding all [structs as arrays](https://godoc.org/../esm#Encoder.StructAsArray) or [individual structs](https://godoc.org/../esm#example-Marshal--AsArray).
- Simple but very fast and efficient [queries](https://godoc.org/../esm#example-Decoder-Query).

API docs: https://godoc.org/../esm.
Examples: https://godoc.org/../esm#pkg-examples.

## Installation

Install:

```shell
go get -u ../esm
```

## Quickstart

```go
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
```

## Benchmark

```
BenchmarkStructVmihailencoEsm-4   	  200000	     12814 ns/op	    2128 B/op	      26 allocs/op
BenchmarkStructUgorjiGoEsm-4      	  100000	     17678 ns/op	    3616 B/op	      70 allocs/op
BenchmarkStructUgorjiGoCodec-4        	  100000	     19053 ns/op	    7346 B/op	      23 allocs/op
BenchmarkStructJSON-4                 	   20000	     69438 ns/op	    7864 B/op	      26 allocs/op
BenchmarkStructGOB-4                  	   10000	    104331 ns/op	   14664 B/op	     278 allocs/op
```

## Howto

Please go through [examples](https://godoc.org/../esm#pkg-examples) to get an idea how to use this package.

## See also

- [Golang PostgreSQL ORM](https://github.com/go-pg/pg)
- [Golang message task queue](https://github.com/go-msgqueue/msgqueue)
