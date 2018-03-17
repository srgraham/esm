package esm

import (
	"reflect"
)

const (


	recordHeaderLen = 24
	groupHeaderLen = 24
	fieldHeaderLen = 6

	stringFileMainHeaderLen     = 8
	stringFileDirectoryEntryLen = 8

	fileHeaderSignature       = 0x54455334 // TES4
	groupHeaderSignature      = 0x47525550 // GRUP
	edidSignature             = "EDID" // 0x45444944 // EDID

	directoryHeaderSignature = 0x02014b50
	directoryEndSignature    = 0x06054b50
	directory64LocSignature  = 0x07064b50
	directory64EndSignature  = 0x06064b50
	dataDescriptorSignature  = 0x08074b50 // de-facto standard; required by OS X Finder
	fileHeaderLen            = 30         // + filename + extra
	directoryHeaderLen       = 46         // + filename + extra + comment
	directoryEndLen          = 22         // + comment
	dataDescriptorLen        = 16         // four uint32: descriptor signature, crc32, compressed size, size
	dataDescriptor64Len      = 24         // descriptor with 8 byte sizes
	directory64LocLen        = 20         //
	directory64EndLen        = 56         // + extra


	// extra header id's
	zip64ExtraId = 0x0001 // zip64 Extended Information Extra Field
)

type null [0]uint8
type char byte
type char4 [4]byte
type wchar uint16
//type int8 int8
//type uint8 uint8
//type int16 int16
//type uint16 uint16
//type int32 int32
//type uint32 uint32
//type int64 int64
//type uint64 uint64
//type float32 float32
//type float64 float64

type DEPRECATED_TYPE null
type ubyte DEPRECATED_TYPE
type short DEPRECATED_TYPE
type ushort DEPRECATED_TYPE
type long DEPRECATED_TYPE
type ulong DEPRECATED_TYPE
type float DEPRECATED_TYPE

type vsval []byte

type formid uint32
type iref uint32
type hash [8]uint8
type filetime uint64
type systemtime [16]uint8
type rgb uint32

type LString string
type dlstring string
type ilstring string
type bstring string
type bzstring string
type wstring string
type wzstring string
type zstring string


func (z zstring) String() string {
	return string(z)
}

//func (z formid) GoString() string {
//	return fmt.Sprintf("0x%x08", z)
//}




func AsUint32(val interface{}) uint32 {
	ref := reflect.ValueOf(val)
	if ref.Kind() != reflect.Uint32 {
		return 0
	}
	return uint32(ref.Uint())
}
func AsUint32Arr(val interface{}) []uint32 {
	ref := reflect.ValueOf(val)
	if ref.Kind() != reflect.Slice && ref.Kind() != reflect.Array {
		return make([]uint32, 0)
	}
	out := make([]uint32, 0);
	for i:=0; i < ref.Len(); i++ {
		item := ref.Index(i)
		out = append(out, uint32(item.Uint()))
	}
	return out
}
func AsString(val interface{}) string {
	ref := reflect.ValueOf(val)
	if ref.Kind() != reflect.String {
		return ""
	}
	return string(ref.String())
}

