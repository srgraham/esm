package esm


// RecordHeader describes a file within a zip file.
// See the zip spec for details.
type RecordHeader struct {
	// Name is the name of the file.
	// It must be a relative path: it must not start with a drive
	// letter (e.g. C:) or leading slash, and only forward slashes
	// are allowed.

	_type [4]byte
	dataSize uint32
	flags uint32
	id uint32
	revision uint32
	version uint16
	unknown uint16
	data []uint8
}


const (


	recordHeaderLen = 24


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
