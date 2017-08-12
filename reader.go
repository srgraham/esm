package esm


import (
"bufio"
"encoding/binary"
"errors"
_ "fmt"
"io"
"os"
	"fmt"
	//"bytes"
	_ "reflect"
	"reflect"

	fields "./fields"
)

var (
	ErrFormat = errors.New("esm: not a valid esm/esp file")
	ErrNilType = errors.New("cannot decode nil type")
	ErrDecodeUnknownType = errors.New("cannot decode %s type")
)

type Reader struct {
	r             io.ReaderAt
	Record        *Record
	Comment       string
	//decompressors map[uint16]Decompressor
}

type ReadCloser struct {
	f *os.File
	Reader
}

type Field struct {
	FieldHeader
	record *Record
	zip *Reader
	zipr io.ReaderAt
	zipsize int64
	dataSectionReader *io.SectionReader
	//headerOffset int64
	//data []byte
	//dataBuf []byte
	dataBuf readBuf
	data interface{}
}

type Record struct {
	RecordHeader
	zip          *Reader
	zipr         io.ReaderAt
	zipsize      int64
	//headerOffset int64
	fields []*Field
}

func (f *Field) String() string {
	//str := fmt.Sprintf("Field[%s](%d): buff: %s", f._type, f.dataSize, f.dataBuf)
	str := fmt.Sprintf("Field[%s](%d): data: %s", f._type, f.dataSize, f.data)
	return str
}

func (f *Field) readData() error {

	rs := io.NewSectionReader(f.zipr, recordHeaderLen, int64(f.dataSize) + int64(recordHeaderLen))
	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return err
	}

	//buf := bufio.NewReader(rs)

	// if type is HEDR, read struct
	data, err := f.getFieldStructure()
	if err != nil {
		return err
	}
	fmt.Println(111, data)

	f.data = data

	return nil
}

func (b *readBuf) readType(t reflect.Type, v reflect.Value) (error) {
	fmt.Println(t.Kind())
	switch t.Kind() {
	case reflect.Map:
		v.Set(reflect.MakeMap(t))
	case reflect.Struct:
		fieldCount := v.NumField()
		for i := 0; i < fieldCount; i+=1 {
			structFieldType := t.Field(i)
			structFieldValue := v.Field(i)
			b.readType(structFieldType.Type, structFieldValue)
		}
	default:

		switch v.Interface().(type) {

		case char:
			v.Set(reflect.ValueOf(b.char()))

		case wchar:
			v.Set(reflect.ValueOf(b.wchar()))

		case uint8:
			v.Set(reflect.ValueOf(b.uint8()))

		case uint16:
			v.Set(reflect.ValueOf(b.uint16()))

		case uint32:
			v.Set(reflect.ValueOf(b.uint32()))

		case uint64:
			v.Set(reflect.ValueOf(b.uint64()))

		case int8:
			v.Set(reflect.ValueOf(b.int8()))

		case int16:
			v.Set(reflect.ValueOf(b.int16()))

		case int32:
			v.Set(reflect.ValueOf(b.int32()))

		case int64:
			v.Set(reflect.ValueOf(b.int64()))

		case float32:
			v.Set(reflect.ValueOf(b.float32()))
			//subVRv := reflect.ValueOf(subV)
			//subRv.Set(reflect.Value(subVRv))

		case float64:
			v.Set(reflect.ValueOf(b.float64()))

		default:
			panic(fmt.Errorf("cannot decode type '%s'", v.Type()))
		}

		return nil
	}
	return nil
}

// FIXME: figure out how to pass by ref correctly
func (b *readBuf) decode(newValuePtr interface{}) (error) {

	//var refIv2 interface{}
	//refIv2 = &iv
	////var refIv interface{}
	////refIv = &iv

	switch v := newValuePtr.(type) {
	case nil:
		err := ErrNilType
		v = nil
		return err

	//case *zstring:
	//	*v = b.decodeZstring()
	case *float32:
		*v = 0

	case *interface{}:
		return nil
		//rv := reflect.ValueOf(iv)
		//rv := reflect.ValueOf(refIv)
		//if rv.Kind() == reflect.Struct {
		//	//out := b.decodeStructByAddr(refIv)
		//	//out := b.decodeStructRv(rv)
		//	out := b.decodeStruct(*v, iv)
		//	//*v = out
		//	*v = out
		//	//iv = out
		//	//refIv = &out
		//
		//	return nil
		//}
		//*v = nil
	case interface{}:
		//*v = nil
		panic(fmt.Errorf("wf"))
	default:
		//rv := reflect.ValueOf(refIv)
		//fmt.Println("unknown kind!", rv.Kind())
		//
		//// check for struct
		//
		//if rv.Kind() == reflect.Struct {
		//	//out := b.decodeStructByAddr(&iv)
		//	//out := b.decodeStructRv(rv)
		//	//*v = out
		//	//v = out
		//	//refIv = &out
		//
		//	return nil
		//}
		//
		//fmt.Println("not struct", rv.Kind())
		//
		//if rv.Kind() != reflect.Ptr {
		//	fmt.Println(rv.Kind())
		//	//err := ErrNilType
		//	//return err
		//}
		//rv = rv.Elem()
		//
		//if rv.IsValid() {
		//	rv.Set(reflect.Zero(rv.Type()))
		//}
		err := fmt.Errorf("cannot decode %s type", v)
		return err
		//err := ErrDecodeUnknownType

	}
	return nil
}

func (b *readBuf) decodeStruct(pointerTypeRefIv interface{}, iv interface{}) interface{} {
	//pointerTypeRefIv = 9.9
	ps := reflect.ValueOf(iv)
	//ps := reflect.ValueOf(pointerTypeRefIv)

	fmt.Println(4444, ps)

	s := ps.Elem()
	//s:=ps

	f:=s.FieldByName("Version")

		//f := s.Field(0)
		f.SetFloat(9.9)

	//fmt.Println(f)
	//fmt.Println(pv)
	return pointerTypeRefIv
}


func (b *readBuf) decodeStructRv(rv reflect.Value) interface{} {

	//structValues := make([]interface{}, rv.NumField())
	//structValues := interface{}

	fmt.Println("looping struct")

	indir := reflect.Indirect(rv)

	for i := 0; i < rv.NumField(); i++ {
		var subV interface{}
		subRv := rv.Field(i)

		a := rv.FieldByName("Version").Float()
		_ = a

		rv.Field(i).SetFloat(9.9)

		fmt.Println("reading at", b)


		indir.SetFloat(9.9)

		// check for structs
		if subRv.Kind() == reflect.Struct {
			subV = b.decodeStructRv(subRv)
		} else {
			// check various types
			//v := rv.Elem()

			switch subRv.Interface().(type) {

			case char:
				subV = b.char()

			case wchar:
				subV = b.wchar()

			case uint8:
				subV = b.uint8()

			case uint16:
				subV = b.uint16()

			case uint32:
				subV = b.uint32()

			case uint64:
				subV = b.uint64()

			case int8:
				subV = b.int8()

			case int16:
				subV = b.int16()

			case int32:
				subV = b.int32()

			case int64:
				subV = b.int64()

			case float32:
				subV = b.float32()
				//subVRv := reflect.ValueOf(subV)
				//subRv.Set(reflect.Value(subVRv))

			case float64:
				subV = b.float64()

			default:
				panic(fmt.Errorf("cannot decode type '%s'", subRv.Type()))
			}


		}

		fmt.Println("decoded", subV, subRv, b)

		p := float32(9)

		f := reflect.ValueOf(p)

		//structValues["poo"] = subV
		subRv.Set(reflect.Value(reflect.ValueOf(f)))




	}

	fmt.Println("out", rv)

	return rv.Interface()
}

func (f *Field) getFieldStructure() (out interface{}, err error) {
	err = nil

	_type := fmt.Sprintf("%s", f._type)

	type S struct {
		Version float32
		NumRecords int32
		NextObjectId uint32
	}

	//out := interface

	switch _type {

	case "HEDR":
		out = S{}
	default:
		out = nil
		// TODO: set err
	}

	//fmt.Println(fields.FieldsStructLookup)

	recordType := fmt.Sprintf("%s", f.record._type)
	fieldTypeStr := _type

	zeroValue := fields.FieldsStructLookup[recordType][fieldTypeStr]

	t := reflect.TypeOf(zeroValue)

	if t == nil {
		err = fmt.Errorf("Unimplemented type %s.%s", recordType, fieldTypeStr)
		return err, err
	}


	v := reflect.New(t)
	f.dataBuf.readType(t, v.Elem())

	return v, err
}



func (record *Record) String() string {
	str := fmt.Sprintf("Record[%s](%d): ", record._type, record.dataSize)
	for _, field := range record.fields {
		str += fmt.Sprintf("%s", field._type) + ", "
	}
	//_type [4]byte
	//dataSize uint32
	//flags uint32
	//id uint32
	//revision uint32
	//version uint16
	//unknown uint16
	//data []uint8

	return str
}

func (record *Record) isMaster() bool {
	return record.flags & 0x1 != 0
}
func (record *Record) isConstant() bool {
	return record.flags & 0x40 != 0
}
func (record *Record) isCompressed() bool {
	return record.flags & 0x40000 != 0
}
func (record *Record) isMarker() bool {
	return record.flags & 0x800000 != 0
}


func (record *Record) readFields() error {

	rs := io.NewSectionReader(record.zipr, recordHeaderLen, int64(record.dataSize) + int64(recordHeaderLen))
	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return err
	}

	reader := bufio.NewReader(rs)

	record.fields = make([]*Field, 0)

	// process all field headers
	for {

		field := &Field{record: record, zip: record.zip, zipr: rs, zipsize: int64(record.dataSize)}
		err := readFieldHeader(field, reader)

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		//field.dataSectionReader := io.NewSectionReader(record.zipr, rs.Seek(), 1000) //int64(field.dataSize) + int64(fieldHeaderLen))
		//
		////reader := bufio.NewReader(rs)
		//
		//field.data = make([]byte, int64(field.dataSize + 20))
		//
		////var buf [recordHeaderLen]byte
		//if _, err := field.dataSectionReader.Read(field.data); err != nil {
		//	return err
		//}

		//fmt.Println(field.data)

		field.readData()

		record.fields = append(record.fields, field)
		// skip past rest of data

		//buf.Discard(int(field.dataSize))

		//fmt.Println(field)

		//buf = buf[field.dataSize:]
	}

	// at this point, only the headers for each field has been grabbed

	// now process the data for each field

	//for field := range record.fields {
	//	field.readData()
	//}

	return nil
}

//func (record *Record) hasDataDescriptor() bool {
//	return f.Flags&0x8 != 0
//}

// OpenReader will open the esm/esp file specified by name and return a ReadCloser.
func OpenReader(name string) (*ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}
	r := new(ReadCloser)
	if err := r.init(f, fi.Size()); err != nil {
		f.Close()
		return nil, err
	}
	r.f = f
	return r, nil
}

// NewReader returns a new Reader reading from r, which is assumed to
// have the given size in bytes.
func NewReader(r io.ReaderAt, size int64) (*Reader, error) {
	zr := new(Reader)
	if err := zr.init(r, size); err != nil {
		return nil, err
	}
	return zr, nil
}

func (z *Reader) init(r io.ReaderAt, size int64) error {
	//end, err := readDirectoryEnd(r, size)
	//if err != nil {
	//	return err
	//}
	//if end.directoryRecords > uint64(size)/fileHeaderLen {
	//	return fmt.Errorf("archive/zip: TOC declares impossible %d files in %d byte zip", end.directoryRecords, size)
	//}
	z.r = r

	//z.Record = readRecordHeader(r)
	//z.Comment = end.comment
	rs := io.NewSectionReader(r, 0, size)
	//if _, err = rs.Seek(int64(end.directoryOffset), io.SeekStart); err != nil {
	//	return err
	//}
	buf := bufio.NewReader(rs)

	rootRecord := &Record{zip:z, zipr: r, zipsize: size}
	err := readRecordHeader(rootRecord, buf)

	if err != nil {
		return err
	}

	// check its a TES4 record
	if binary.BigEndian.Uint32([]byte(rootRecord._type[:])) != fileHeaderSignature {
		return ErrFormat
	}
	// make sure its the right form id (0)
	if rootRecord.id != 0 {
		return ErrFormat
	}

	// read the fields of the root record
	rootRecord.readFields()

	return nil

	//// The count of files inside a zip is truncated to fit in a uint16.
	//// Gloss over this by reading headers until we encounter
	//// a bad one, and then only report a ErrFormat or UnexpectedEOF if
	//// the file count modulo 65536 is incorrect.
	//for {
	//	f := &Record{zip: z, zipr: r, zipsize: size}
	//	err = readDirectoryHeader(f, buf)
	//	if err == ErrFormat || err == io.ErrUnexpectedEOF {
	//		break
	//	}
	//	if err != nil {
	//		return err
	//	}
	//	z.Record = append(z.Record, f)
	//}
	//if uint16(len(z.Record)) != uint16(end.directoryRecords) { // only compare 16 bits here
	//	// Return the readDirectoryHeader error if we read
	//	// the wrong number of directory entries.
	//	return err
	//}
	//return nil
}

// RegisterDecompressor registers or overrides a custom decompressor for a
// specific method ID. If a decompressor for a given method is not found,
// Reader will default to looking up the decompressor at the package level.
//func (z *Reader) RegisterDecompressor(method uint16, dcomp Decompressor) {
//	if z.decompressors == nil {
//		z.decompressors = make(map[uint16]Decompressor)
//	}
//	z.decompressors[method] = dcomp
//}
//
//func (z *Reader) decompressor(method uint16) Decompressor {
//	dcomp := z.decompressors[method]
//	if dcomp == nil {
//		dcomp = decompressor(method)
//	}
//	return dcomp
//}

// Close closes the Zip file, rendering it unusable for I/O.
func (rc *ReadCloser) Close() error {
	return rc.f.Close()
}

// DataOffset returns the offset of the file's possibly-compressed
// data, relative to the beginning of the zip file.
//
// Most callers should instead use Open, which transparently
// decompresses data and verifies checksums.
//func (record *Record) DataOffset() (offset int64, err error) {
//	bodyOffset, err := record.findBodyOffset()
//	if err != nil {
//		return
//	}
//	return record.headerOffset + bodyOffset, nil
//}

// Open returns a ReadCloser that provides access to the Record's contents.
// Multiple files may be read concurrently.
//func (f *Record) Open() (io.ReadCloser, error) {
//	bodyOffset, err := f.findBodyOffset()
//	if err != nil {
//		return nil, err
//	}
//	size := int64(f.dataSize)
//	r := io.NewSectionReader(f.zipr, f.headerOffset+bodyOffset, size)
//
//	dcomp := f.zip.decompressor(f.Method)
//	if dcomp == nil {
//		return nil, ErrAlgorithm
//	}
//	var rc io.ReadCloser = dcomp(r)
//	var desr io.Reader
//	if f.hasDataDescriptor() {
//		desr = io.NewSectionReader(f.zipr, f.headerOffset+bodyOffset+size, dataDescriptorLen)
//	}
//	rc = &checksumReader{
//		rc:   rc,
//		hash: crc32.NewIEEE(),
//		f:    f,
//		desr: desr,
//	}
//	return rc, nil
//}


// findBodyOffset does the minimum work to verify the file has a header
// and returns the file body offset.
//func (record *Record) findBodyOffset() (int64, error) {
//	var buf [fileHeaderLen]byte
//	if _, err := record.zipr.ReadAt(buf[:], record.headerOffset); err != nil {
//		return 0, err
//	}
//	b := readBuf(buf[:])
//	if sig := b.uint32(); sig != fileHeaderSignature {
//		return 0, ErrFormat
//	}
//	b = b[22:] // skip over most of the header
//	filenameLen := int(b.uint16())
//	extraLen := int(b.uint16())
//	return int64(fileHeaderLen + filenameLen + extraLen), nil
//}

func readRecordHeader(record *Record, r io.Reader) error {
	var buf [recordHeaderLen]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
		return err
	}
	b := readBuf(buf[:])

	// TODO: validate signature is in the list of allowed Record header types?

	record._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}

	record.dataSize = b.uint32()
	record.flags = b.uint32()
	record.id = b.uint32()
	record.revision = b.uint32()
	record.version = b.uint16()
	record.unknown = b.uint16()

	fmt.Println(record)

	return nil
}
func readFieldHeader(field *Field, r io.Reader) error {
	var bufHeader [fieldHeaderLen]byte
	if _, err := io.ReadFull(r, bufHeader[:]); err != nil {
		return err
	}
	b := readBuf(bufHeader[:])

	// TODO: validate signature is in the list of allowed Record header types?

	field._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}

	field.dataSize = b.uint16()

	field.dataBuf = make([]byte, field.dataSize)

	if _, err := io.ReadFull(r, field.dataBuf); err != nil {
		return err
	}

	return nil
}

// readDirectoryHeader attempts to read a directory header from r.
// It returns io.ErrUnexpectedEOF if it cannot read a complete header,
// and ErrFormat if it doesn't find a valid header signature.
//func readDirectoryHeader(record *Record, r io.Reader) error {
//	var buf [directoryHeaderLen]byte
//	if _, err := io.ReadFull(r, buf[:]); err != nil {
//		return err
//	}
//	b := readBuf(buf[:])
//	if sig := b.uint32(); sig != directoryHeaderSignature {
//		return ErrFormat
//	}
//	f.CreatorVersion = b.uint16()
//	f.ReaderVersion = b.uint16()
//	f.Flags = b.uint16()
//	f.Method = b.uint16()
//	f.ModifiedTime = b.uint16()
//	f.ModifiedDate = b.uint16()
//	f.CRC32 = b.uint32()
//	f.CompressedSize = b.uint32()
//	f.UncompressedSize = b.uint32()
//	f.CompressedSize64 = uint64(f.CompressedSize)
//	f.UncompressedSize64 = uint64(f.UncompressedSize)
//	filenameLen := int(b.uint16())
//	extraLen := int(b.uint16())
//	commentLen := int(b.uint16())
//	b = b[4:] // skipped start disk number and internal attributes (2x uint16)
//	f.ExternalAttrs = b.uint32()
//	f.headerOffset = int64(b.uint32())
//	d := make([]byte, filenameLen+extraLen+commentLen)
//	if _, err := io.ReadFull(r, d); err != nil {
//		return err
//	}
//	f.Name = string(d[:filenameLen])
//	f.Extra = d[filenameLen : filenameLen+extraLen]
//	f.Comment = string(d[filenameLen+extraLen:])
//
//	needUSize := f.UncompressedSize == ^uint32(0)
//	needCSize := f.CompressedSize == ^uint32(0)
//	needHeaderOffset := f.headerOffset == int64(^uint32(0))
//
//	if len(f.Extra) > 0 {
//		// Best effort to find what we need.
//		// Other zip authors might not even follow the basic format,
//		// and we'll just ignore the Extra content in that case.
//		b := readBuf(f.Extra)
//		for len(b) >= 4 { // need at least tag and size
//			tag := b.uint16()
//			size := b.uint16()
//			if int(size) > len(b) {
//				break
//			}
//			if tag == zip64ExtraId {
//				// update directory values from the zip64 extra block.
//				// They should only be consulted if the sizes read earlier
//				// are maxed out.
//				// See golang.org/issue/13367.
//				eb := readBuf(b[:size])
//
//				if needUSize {
//					needUSize = false
//					if len(eb) < 8 {
//						return ErrFormat
//					}
//					f.UncompressedSize64 = eb.uint64()
//				}
//				if needCSize {
//					needCSize = false
//					if len(eb) < 8 {
//						return ErrFormat
//					}
//					f.CompressedSize64 = eb.uint64()
//				}
//				if needHeaderOffset {
//					needHeaderOffset = false
//					if len(eb) < 8 {
//						return ErrFormat
//					}
//					f.headerOffset = int64(eb.uint64())
//				}
//				break
//			}
//			b = b[size:]
//		}
//	}
//
//	// Assume that uncompressed size 2³²-1 could plausibly happen in
//	// an old zip32 file that was sharding inputs into the largest chunks
//	// possible (or is just malicious; search the web for 42.zip).
//	// If needUSize is true still, it means we didn't see a zip64 extension.
//	// As long as the compressed size is not also 2³²-1 (implausible)
//	// and the header is not also 2³²-1 (equally implausible),
//	// accept the uncompressed size 2³²-1 as valid.
//	// If nothing else, this keeps archive/zip working with 42.zip.
//	_ = needUSize
//
//	if needCSize || needHeaderOffset {
//		return ErrFormat
//	}
//
//	return nil
//}

//func readDataDescriptor(r io.Reader, f *Record) error {
//	var buf [dataDescriptorLen]byte
//
//	// The spec says: "Although not originally assigned a
//	// signature, the value 0x08074b50 has commonly been adopted
//	// as a signature value for the data descriptor record.
//	// Implementers should be aware that ZIP files may be
//	// encountered with or without this signature marking data
//	// descriptors and should account for either case when reading
//	// ZIP files to ensure compatibility."
//	//
//	// dataDescriptorLen includes the size of the signature but
//	// first read just those 4 bytes to see if it exists.
//	if _, err := io.ReadFull(r, buf[:4]); err != nil {
//		return err
//	}
//	off := 0
//	maybeSig := readBuf(buf[:4])
//	if maybeSig.uint32() != dataDescriptorSignature {
//		// No data descriptor signature. Keep these four
//		// bytes.
//		off += 4
//	}
//	if _, err := io.ReadFull(r, buf[off:12]); err != nil {
//		return err
//	}
//	b := readBuf(buf[:12])
//	if b.uint32() != f.CRC32 {
//		return ErrChecksum
//	}
//
//	// The two sizes that follow here can be either 32 bits or 64 bits
//	// but the spec is not very clear on this and different
//	// interpretations has been made causing incompatibilities. We
//	// already have the sizes from the central directory so we can
//	// just ignore these.
//
//	return nil
//}


//func readDirectoryEnd(r io.ReaderAt, size int64) (dir *directoryEnd, err error) {
//	// look for directoryEndSignature in the last 1k, then in the last 65k
//	var buf []byte
//	var directoryEndOffset int64
//	for i, bLen := range []int64{1024, 65 * 1024} {
//		if bLen > size {
//			bLen = size
//		}
//		buf = make([]byte, int(bLen))
//		if _, err := r.ReadAt(buf, size-bLen); err != nil && err != io.EOF {
//			return nil, err
//		}
//		if p := findSignatureInBlock(buf); p >= 0 {
//			buf = buf[p:]
//			directoryEndOffset = size - bLen + int64(p)
//			break
//		}
//		if i == 1 || bLen == size {
//			return nil, ErrFormat
//		}
//	}
//
//	// read header into struct
//	b := readBuf(buf[4:]) // skip signature
//	d := &directoryEnd{
//		diskNbr:            uint32(b.uint16()),
//		dirDiskNbr:         uint32(b.uint16()),
//		dirRecordsThisDisk: uint64(b.uint16()),
//		directoryRecords:   uint64(b.uint16()),
//		directorySize:      uint64(b.uint32()),
//		directoryOffset:    uint64(b.uint32()),
//		commentLen:         b.uint16(),
//	}
//	l := int(d.commentLen)
//	if l > len(b) {
//		return nil, errors.New("zip: invalid comment length")
//	}
//	d.comment = string(b[:l])
//
//	// These values mean that the file can be a zip64 file
//	if d.directoryRecords == 0xffff || d.directorySize == 0xffff || d.directoryOffset == 0xffffffff {
//		p, err := findDirectory64End(r, directoryEndOffset)
//		if err == nil && p >= 0 {
//			err = readDirectory64End(r, p, d)
//		}
//		if err != nil {
//			return nil, err
//		}
//	}
//	// Make sure directoryOffset points to somewhere in our file.
//	if o := int64(d.directoryOffset); o < 0 || o >= size {
//		return nil, ErrFormat
//	}
//	return d, nil
//}

// findDirectory64End tries to read the zip64 locator just before the
// directory end and returns the offset of the zip64 directory end if
// found.
//func findDirectory64End(r io.ReaderAt, directoryEndOffset int64) (int64, error) {
//	locOffset := directoryEndOffset - directory64LocLen
//	if locOffset < 0 {
//		return -1, nil // no need to look for a header outside the file
//	}
//	buf := make([]byte, directory64LocLen)
//	if _, err := r.ReadAt(buf, locOffset); err != nil {
//		return -1, err
//	}
//	b := readBuf(buf)
//	if sig := b.uint32(); sig != directory64LocSignature {
//		return -1, nil
//	}
//	if b.uint32() != 0 { // number of the disk with the start of the zip64 end of central directory
//		return -1, nil // the file is not a valid zip64-file
//	}
//	p := b.uint64()      // relative offset of the zip64 end of central directory record
//	if b.uint32() != 1 { // total number of disks
//		return -1, nil // the file is not a valid zip64-file
//	}
//	return int64(p), nil
//}

//// readDirectory64End reads the zip64 directory end and updates the
//// directory end with the zip64 directory end values.
//func readDirectory64End(r io.ReaderAt, offset int64, d *directoryEnd) (err error) {
//	buf := make([]byte, directory64EndLen)
//	if _, err := r.ReadAt(buf, offset); err != nil {
//		return err
//	}
//
//	b := readBuf(buf)
//	if sig := b.uint32(); sig != directory64EndSignature {
//		return ErrFormat
//	}
//
//	b = b[12:]                        // skip dir size, version and version needed (uint64 + 2x uint16)
//	d.diskNbr = b.uint32()            // number of this disk
//	d.dirDiskNbr = b.uint32()         // number of the disk with the start of the central directory
//	d.dirRecordsThisDisk = b.uint64() // total number of entries in the central directory on this disk
//	d.directoryRecords = b.uint64()   // total number of entries in the central directory
//	d.directorySize = b.uint64()      // size of the central directory
//	d.directoryOffset = b.uint64()    // offset of start of central directory with respect to the starting disk number
//
//	return nil
//}

func findSignatureInBlock(b []byte) int {
	for i := len(b) - directoryEndLen; i >= 0; i-- {
		// defined from directoryEndSignature in struct.go
		if b[i] == 'P' && b[i+1] == 'K' && b[i+2] == 0x05 && b[i+3] == 0x06 {
			// n is length of comment
			n := int(b[i+directoryEndLen-2]) | int(b[i+directoryEndLen-1])<<8
			if n+directoryEndLen+i <= len(b) {
				return i
			}
		}
	}
	return -1
}

