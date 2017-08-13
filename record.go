package esm

import (
	"io"
	"fmt"
	"bufio"
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

type Record struct {
	RecordHeader
	readerAt         io.ReaderAt
	zipsize      int64
	//headerOffset int64
	fields []*Field
}


func (r *Record) readHeader(reader io.Reader) error {
	var buf [recordHeaderLen]byte
	if _, err := io.ReadFull(reader, buf[:]); err != nil {
		return err
	}
	b := readBuf(buf[:])

	// TODO: validate signature is in the list of allowed Record header types?

	r._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}

	r.dataSize = b.uint32()
	r.flags = b.uint32()
	r.id = b.uint32()
	r.revision = b.uint32()
	r.version = b.uint16()
	r.unknown = b.uint16()

	fmt.Println(r)

	return nil
}


func (r *Record) Type() (string){
	return fmt.Sprintf("%s", r._type)
}

func (r *Record) String() string {
	str := fmt.Sprintf("Record[%s](%d): ", r.Type(), r.dataSize)
	for _, field := range r.fields {
		str += fmt.Sprintf("%s", field.Type()) + ", "
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

func (r *Record) isMaster() bool {
	return r.flags & 0x1 != 0
}
func (r *Record) isConstant() bool {
	return r.flags & 0x40 != 0
}
func (r *Record) isCompressed() bool {
	return r.flags & 0x40000 != 0
}
func (r *Record) isMarker() bool {
	return r.flags & 0x800000 != 0
}


func (r *Record) readFields() error {

	rs := io.NewSectionReader(r.readerAt, recordHeaderLen, int64(r.dataSize) + int64(recordHeaderLen))
	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return err
	}

	reader := bufio.NewReader(rs)

	r.fields = make([]*Field, 0)

	// process all field headers
	for {

		field := &Field{record: r, readerAt: rs, zipsize: int64(r.dataSize)}
		err := field.readHeader(reader)

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		//field.dataSectionReader := io.NewSectionReader(r.readerAt, rs.Seek(), 1000) //int64(field.dataSize) + int64(fieldHeaderLen))
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

		r.fields = append(r.fields, field)
		// skip past rest of data

		//buf.Discard(int(field.dataSize))

		//fmt.Println(field)

		//buf = buf[field.dataSize:]
	}

	// at this point, only the headers for each field has been grabbed

	// now process the data for each field

	//for field := range r.fields {
	//	field.readData()
	//}

	return nil
}






