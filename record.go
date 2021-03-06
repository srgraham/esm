package esm

import (
	"io"
	"fmt"
	"strconv"
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
	formid formid
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

	if r.Type() == "GRUP" {
		return ErrRecordIsGRUP
	}

	r.dataSize = b.uint32()
	r.flags = b.uint32()
	r.formid = b.formid()
	r.revision = b.uint32()
	r.version = b.uint16()
	r.unknown = b.uint16()

	FormIds[r.formid] = r

	if r.formid == 0xa06e6 {
		fmt.Print(999999)
	}

	//fmt.Println(r)

	return nil
}


func (r *Record) Type() (string){
	return fmt.Sprintf("%s", r._type)
}

func (r *Record) String() string {
	str := fmt.Sprintf("Record[%s]{%d}(%d): ", r.Type(), r.FormId(), r.dataSize)
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

func (r *Record) Dump() string {
	str := r.String()
	str += "\nformId: " + strconv.FormatUint(uint64(r.FormId()), 10)
	for _, field := range r.fields {
		str += "\n" + field.String()
	}
	return str + "\n"
}

func (r *Record) DebugHeader() {
	fmt.Printf("_type: %#v\ndataSize: %#v\nflags: %#v\nformid: %#v\nrevision: %#v\nversion: %#v\nunknown: %#v\ndata: %#v\n", r._type, r.dataSize, r.flags, r.formid, r.revision, r.version, r.unknown, r.data)
}

func (r *Record) FormId() uint32 {
	return uint32(r.formid)
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

func (r *Record) ParentGroup() *Group {
	return r.parentGroup
}

func (r *Record) NearestParentRecord() *Record {
	if r.ParentGroup() != nil {
		return r.ParentGroup().NearestParentRecord()
	}
	return nil
}


func (r *Record) readFields(reader io.ReaderAt) error {

	//fmt.Println(r.String())

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


func (r *Record) GetOneFieldForType(field_type string) *Field {
	for _, field := range r.fields {
		if field.Type() == field_type {
			return field
		}
	}
	return nil
}



func (r *Record) GetFieldDataForType(field_type string) interface{} {
	field := r.GetOneFieldForType(field_type)
	if field == nil {
		return nil
	}
	return field.data
}



