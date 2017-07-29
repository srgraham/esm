package esm

import (
	"encoding/binary"

)


// RecordHeader describes a file within a zip file.
// See the zip spec for details.
type RecordHeader struct {
	// Name is the name of the file.
	// It must be a relative path: it must not start with a drive
	// letter (e.g. C:) or leading slash, and only forward slashes
	// are allowed.

	_type char4
	dataSize uint32
	flags uint32
	id uint32
	revision uint32
	version uint16
	unknown uint16
	data []uint8
}

type FieldHeader struct {
	_type char4
	dataSize uint16
}


const (


	recordHeaderLen = 24
	fieldHeaderLen = 6


	fileHeaderSignature      = 0x54455334 // TES4

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

type lstring string
type dlstring string
type ilstring string
type bstring string
type bzstring string
type wstring string
type wzstring string
type zstring string



type readBuf []byte






//func (b *readBuf) uint16() uint16 {
//	v := binary.LittleEndian.Uint16(*b)
//	*b = (*b)[2:]
//	return v
//}
//
//func (b *readBuf) uint32() uint32 {
//	v := binary.LittleEndian.Uint32(*b)
//	*b = (*b)[4:]
//	return v
//}
//
//func (b *readBuf) uint64() uint64 {
//	v := binary.LittleEndian.Uint64(*b)
//	*b = (*b)[8:]
//	return v
//}
//
//func (b *readBuf) byte() byte {
//	v := (*b)[0]
//	*b = (*b)[1:]
//	return v
//}
//
//func (b *readBuf) skip(size int64) {
//	*b = (*b)[size:]
//	return
//}




func (b *readBuf) char() char { return char(b.uint8()) }

func (b *readBuf) wchar() {panic("Unimplemented readBuf type")}


//func (b *readBuf) bool() bool {
//	x := (*b)[0]
//	*b = (*b)[1:]
//	return x != 0
//}

func (b *readBuf) uint8() uint8 {
	x := (*b)[0]
	*b = (*b)[1:]
	return x
}

func (b *readBuf) uint16() uint16 {
	x := binary.LittleEndian.Uint16((*b)[0:2])
	*b = (*b)[2:]
	return x
}

func (b *readBuf) uint32() uint32 {
	x := binary.LittleEndian.Uint32((*b)[0:4])
	*b = (*b)[4:]
	return x
}

func (b *readBuf) uint64() uint64 {
	x := binary.LittleEndian.Uint64((*b)[0:8])
	*b = (*b)[8:]
	return x
}

func (b *readBuf) int8() int8 { return int8(b.uint8()) }

func (b *readBuf) int16() int16 { return int16(b.uint16()) }

func (b *readBuf) int32() int32 { return int32(b.uint32()) }

func (b *readBuf) int64() int64 { return int64(b.uint64()) }

func (b *readBuf) float32() float32 { return float32(b.uint32()) }

func (b *readBuf) float64() float64 { return float64(b.uint64()) }




// DEPRECATED TYPES
func (b *readBuf) ubyte() {panic("Deprecated readBuf type")}
func (b *readBuf) short() {panic("Deprecated readBuf type")}
func (b *readBuf) ushort() {panic("Deprecated readBuf type")}
func (b *readBuf) long() {panic("Deprecated readBuf type")}
func (b *readBuf) ulong() {panic("Deprecated readBuf type")}
func (b *readBuf) float() {panic("Deprecated readBuf type")}


//// todo: write special parser for this
//type vsval []byte
//
//type formid uint32
//type iref uint32
//type hash [8]uint8
//type filetime uint64
//type systemtime [16]uint8
//type rgb uint32
//
//type lstring string
//type dlstring string
//type ilstring string
//type bstring string
//type bzstring string
//type wstring string
//type wzstring string
//type zstring string












