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

)

var (
	ErrFormat = errors.New("esm: not a valid esm/esp file")
	ErrNilType = errors.New("cannot decode nil type")
	ErrDecodeUnknownType = errors.New("cannot decode %s type")
	ErrUnimplementedField = errors.New("Unimplemented field")
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
	if err != nil && err != ErrUnimplementedField {
		return err
	}
	fmt.Println("field_data", data)

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

		var rv reflect.Value

		switch v.Interface().(type) {

		case char:
			rv = reflect.ValueOf(b.char())

		case wchar:
			rv = reflect.ValueOf(b.wchar())

		case uint8:
			rv = reflect.ValueOf(b.uint8())

		case uint16:
			rv = reflect.ValueOf(b.uint16())

		case uint32:
			rv = reflect.ValueOf(b.uint32())

		case uint64:
			rv = reflect.ValueOf(b.uint64())

		case int8:
			rv = reflect.ValueOf(b.int8())

		case int16:
			rv = reflect.ValueOf(b.int16())

		case int32:
			rv = reflect.ValueOf(b.int32())

		case int64:
			rv = reflect.ValueOf(b.int64())

		case float32:
			rv = reflect.ValueOf(b.float32())
			//subVRv := reflect.ValueOf(subV)
			//subRrv = reflect.Value(subVRv)

		case float64:
			rv = reflect.ValueOf(b.float64())

		case zstring:
			rv = reflect.ValueOf(b.zstring())

		default:
			panic(fmt.Errorf("cannot decode type '%s'", v.Type()))
		}

		fmt.Println(rv, reflect.Indirect(rv))

		v.Set(reflect.Indirect(rv))

		return nil
	}
	return nil
}


func (f *Field) getFieldStructure() (out interface{}, err error) {

	_type := fmt.Sprintf("%s", f._type)

	recordTypeStr := fmt.Sprintf("%s", f.record._type)
	fieldTypeStr := _type

	zeroValue := FieldsStructLookup[recordTypeStr][fieldTypeStr]

	t := reflect.TypeOf(zeroValue)

	if t == nil {
		err = ErrUnimplementedField
		msg := fmt.Sprintf("Unimplemented field %s.%s", recordTypeStr, fieldTypeStr)
		return msg, err
	}

	v := reflect.New(t)
	f.dataBuf.readType(t, v.Elem())

	return v.Elem(), nil
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

}


// Close closes the Zip file, rendering it unusable for I/O.
func (rc *ReadCloser) Close() error {
	return rc.f.Close()
}


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
