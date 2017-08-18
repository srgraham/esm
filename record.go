package esm

import (
	"io"
	"fmt"
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
	parentGroup *Group
	sr          *io.SectionReader
	off         int64
	fields      []*Field
}

// calculates the size of all the things (header and all data)
func (r *Record) Size() int64 {
	return int64(recordHeaderLen) + int64(r.dataSize)
}

func (r *Record) readHeader(sr io.SectionReader) error {
	buf := make([]byte, recordHeaderLen)
	//fmt.Println(sr.Size())
	if _, err := sr.Read(buf); err != nil {
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
//func (parentRecord *Record) hasDataDescriptor() bool {
//	return f.Flags&0x8 != 0
//}
func (r *Record) isConstant() bool {
	return r.flags & 0x40 != 0
}
func (r *Record) isLocalized() bool {
	return r.flags & 0x80 != 0
}
func (r *Record) isCompressed() bool {
	return r.flags & 0x40000 != 0
}
func (r *Record) isMarker() bool {
	return r.flags & 0x800000 != 0
}


func (r *Record) Root() *Root {
	return r.parentGroup.Root()
}


func (r *Record) readFields(reader io.ReaderAt) error {

	//rs := io.NewSectionReader(r.readerAt, recordHeaderLen, int64(r.dataSize) + int64(recordHeaderLen))
	//if _, err := rs.Seek(0, io.SeekStart); err != nil {
	//	return err
	//}

	//reader := bufio.NewReader(rs)

	off := r.off + recordHeaderLen

	//rs := io.NewSectionReader(reader, off, int64(r.dataSize))

	r.fields = make([]*Field, 0)

	// process all field headers
	for off < r.off + r.Size() {

		headerReader := io.NewSectionReader(reader, off, fieldHeaderLen)

		field := &Field{parentRecord: r, off: off}
		err := field.readHeader(*headerReader)

		off += fieldHeaderLen

		sr := io.NewSectionReader(reader, off, field.Size() - fieldHeaderLen)
		field.sr = sr

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		err = field.readData(reader)
		if err != nil {
			return err
		}


		off += int64(field.dataSize)

		r.fields = append(r.fields, field)


	}

	return nil
}

func (r *Record) fieldsByType(field_type string) []*Field {
	fields := make([]*Field, 0)
	for _, field := range r.fields {
		if field.Type() == field_type {
			fields = append(fields, field)
		}
	}
	return fields
}




