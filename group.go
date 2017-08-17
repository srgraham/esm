package esm

import (
	"io"
	"encoding/binary"
	"fmt"
)

type GroupHeader struct {
	_type char4
	groupSize uint32
	label char4
	groupType uint32
	stamp uint16
	unknown uint16
	version uint16
	unknown2 uint16
	data []uint8
}

type Group struct {
	GroupHeader
	parentRoot *Root
	sr         *io.SectionReader
	off        int64
	records []*Record
}


// calculates the size of all the things (header and all data)
func (g *Group) Size() int64 {
	return int64(g.groupSize + groupHeaderLen)
}


func (g *Group) readHeader(sr io.SectionReader) error {

	buf := make([]byte, groupHeaderLen)
	if _, err := sr.Read(buf); err != nil {
		return err
	}
	b := readBuf(buf[:])

	// TODO: validate signature is in the list of allowed Record header types?

	g._type = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}
	g.groupSize = b.uint32() - groupHeaderLen
	g.label = char4{byte(b.char()), byte(b.char()), byte(b.char()), byte(b.char())}
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
	return fmt.Sprintf("%s", g.label)
}

func (g *Group) String() string {
	str := fmt.Sprintf("Group[%s](%d): ", g.Type(), g.groupSize)
	for _, record := range g.records {
		str += fmt.Sprintf("%s", record.Type()) + ", "
	}

	return str
}


func (g *Group) readRecords(reader io.ReaderAt) error {

	g.records = make([]*Record, 0)

	var off int64 = g.off + groupHeaderLen

	for off < g.off + g.Size() {
		headerReader := io.NewSectionReader(reader, off, recordHeaderLen)

		record := &Record{parentGroup: g, off: off}
		err := record.readHeader(*headerReader)

		off += recordHeaderLen

		sr := io.NewSectionReader(reader, off, record.Size() - recordHeaderLen)
		record.sr = sr



		off += record.Size() - recordHeaderLen

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		err = record.readFields(reader)
		if err != nil {
			return err
		}

		g.records = append(g.records, record)
	}

	return nil





}


















