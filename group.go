package esm

import "io"
import (
	"encoding/binary"
	"fmt"
)

type GroupHeader struct {
	// Name is the name of the file.
	// It must be a relative path: it must not start with a drive
	// letter (e.g. C:) or leading slash, and only forward slashes
	// are allowed.

	_type char4
	groupSize uint32
	label uint32
	groupType uint32
	stamp uint16
	unknown uint16
	version uint16
	unknown2 uint16
	data []uint8
}

type Group struct {
	GroupHeader
	root *Root
	sr *io.SectionReader
	off int64
	readerAt io.ReaderAt
	readerSize int64
	//headerOffset int64
	records []*Record
}


// calculates the size of all the things (header and all data)
func (g *Group) Size() int64 {
	return int64(g.groupSize)
}


func (g *Group) readHeader(sr io.SectionReader) error {

	buf := make([]byte, groupHeaderLen)
	if _, err := sr.Read(buf); err != nil {
		return err
	}
	b := readBuf(buf[:])

	// TODO: validate signature is in the list of allowed Record header types?

	g._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}
	g.groupSize = b.uint32()
	g.label = b.uint32()
	g.groupType = b.uint32()
	g.stamp = b.uint16()
	g.unknown = b.uint16()
	g.version = b.uint16()
	g.unknown2 = b.uint16()

	fmt.Println(g)

	return nil
}



func (g *Group) isValid() bool {
	return binary.BigEndian.Uint32([]byte(g._type[:])) == groupHeaderSignature
}



func (g *Group) Type() (string){
	return fmt.Sprintf("%s", g._type)
}

func (g *Group) String() string {
	str := fmt.Sprintf("Group[%s](%d): ", g.Type(), g.groupSize)
	for _, record := range g.records {
		str += fmt.Sprintf("%s", record.Type()) + ", "
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


func (g *Group) readRecords(reader io.ReaderAt) error {


	// read from the start of the file + recordSize depth to start reading the groups
	//recordsSr := io.NewSectionReader(g.readerAt, 0, g.Size())
	//if _, err := recordsSr.Seek(0, io.SeekStart); err != nil {
	//	return err
	//}

	g.records = make([]*Record, 0)

	var off int64 = g.off

	for {
		headerReader := io.NewSectionReader(reader, off, recordHeaderLen)

		record := &Record{group: g}
		err := record.readHeader(*headerReader)

		off += recordHeaderLen

		sr := io.NewSectionReader(reader, off, record.Size())
		record.sr = sr

		off += record.Size()

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		record.readFields(reader)

		g.records = append(g.records, record)
	}

	return nil






	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//rs := io.NewSectionReader(g.readerAt, recordHeaderLen, int64(g.groupSize))
	//if _, err := rs.Seek(0, io.SeekStart); err != nil {
	//	return err
	//}
	//
	//reader := bufio.NewReader(rs)
	//
	//g.records = make([]*Record, 0)
	//
	//// process all field headers
	//for {
	//
	//	record := &Record{group: g, readerAt: rs, readerSize: int64(g.groupSize)}
	//	err := record.readHeader(reader)
	//
	//	if err == ErrFormat || err == io.ErrUnexpectedEOF {
	//		break
	//	}
	//	if err != nil {
	//		return err
	//	}
	//
	//	//record.dataSectionReader := io.NewSectionReader(g.readerAt, rs.Seek(), 1000) //int64(record.dataSize) + int64(fieldHeaderLen))
	//	//
	//	////reader := bufio.NewReader(rs)
	//	//
	//	//record.data = make([]byte, int64(record.dataSize + 20))
	//	//
	//	////var buf [recordHeaderLen]byte
	//	//if _, err := record.dataSectionReader.Read(record.data); err != nil {
	//	//	return err
	//	//}
	//
	//	//fmt.Println(record.data)
	//
	//	record.readFields()
	//
	//	g.records = append(g.records, record)
	//	// skip past rest of data
	//
	//	//buf.Discard(int(record.dataSize))
	//
	//	//fmt.Println(record)
	//
	//	//buf = buf[record.dataSize:]
	//}
	//
	//// at this point, only the headers for each record has been grabbed
	//
	//// now process the data for each record
	//
	////for record := range g.records {
	////	record.readData()
	////}
	//
	//return nil
}


















