package esm

import (
	"io"
	"encoding/binary"
	"fmt"
	"bytes"
	"compress/zlib"
	//"io/ioutil"
	"io/ioutil"
	"strconv"
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
	subGroups []*Group
	parentGroup *Group
	parentRecord *Record
}


// calculates the size of all the things (header and all data)
func (g *Group) Size() int64 {
	return int64(g.groupSize + groupHeaderLen)
}

func (g *Group) Root() *Root {
	return g.parentRoot
}
func (g *Group) ParentGroup() *Group {
	return g.parentGroup
}
func (g *Group) ParentRecord() *Record {
	return g.parentRecord
}


func (g *Group) Dump() string {
	str := g.String() + "\nRecord Count: " + strconv.Itoa(len(g.records)) + "\nSubGroups: " + strconv.Itoa(len(g.subGroups))
	if g.parentRecord != nil {
		str += "\nParent Record:\n" + g.parentRecord.Dump()
	}
	return str
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

	//fmt.Println(g)

	return nil
}



func (g *Group) isValid() bool {
	return binary.BigEndian.Uint32([]byte(g._type[:])) == groupHeaderSignature
}



func (g *Group) Type() (string){
	return fmt.Sprintf("%s", g.label)
}

func (g *Group) String() string {
	var outType string
	switch(g.groupType){
	case 0:
		outType = fmt.Sprintf("%s", g.Type())
	case 1:
		outType = fmt.Sprintf("WORLD_0x%x", g.label)
	case 2:
		outType = fmt.Sprintf("INSIDE_0x%x", g.label)
	case 3:
		outType = fmt.Sprintf("INSIDE_SUB_0x%x", g.label)
	case 4:
		outType = fmt.Sprintf("OUTSIDE_0x%x", g.label)
	case 5:
		outType = fmt.Sprintf("OUTSIDE_SUB_0x%x", g.label)
	case 6:
		outType = fmt.Sprintf("CELL_CHILD_0x%x", g.label)
	case 7:
		outType = fmt.Sprintf("TOPIC_CHILD_0x%x", g.label)
	case 8:
		outType = fmt.Sprintf("CELL_PERSIST_CHILD_0x%x", g.label)
	case 9:
		outType = fmt.Sprintf("CELL_TEMP_CHILD_0x%x", g.label)
	case 10:
		outType = fmt.Sprintf("CELL_VISIBLE_DIST_CHILD_0x%x", g.label)
	default:
		outType = fmt.Sprintf("%d.%d", g.groupType, g.label)
	}
	str := fmt.Sprintf("Group[%s](%d): ", outType, g.groupSize)
	for _, record := range g.records {
		str += fmt.Sprintf("%s", record.Type()) + ", "
	}

	return str
}


func (g *Group) readRecords(reader io.ReaderAt) error {

	g.records = make([]*Record, 0)

	var off int64 = g.off + groupHeaderLen

	currentReader := reader
	g.subGroups = make([]*Group, 0)

	for off < g.off + g.Size() {
		headerReader := io.NewSectionReader(reader, off, recordHeaderLen)

		record := &Record{parentGroup: g, off: off}
		err := record.readHeader(*headerReader)

		currentReader = reader

		if err == ErrRecordIsGRUP {

			var nextGroup *Group
			var err2 error
			nextGroup, off, err2 = g.Root().readNextGroup(currentReader, off)
			nextGroup.parentGroup = g

			if len(g.records) > 0 {
				nextGroup.parentRecord = g.records[len(g.records) - 1]
			}

			if err2 != nil {
				return err
			}

			g.subGroups = append(g.subGroups, nextGroup)
			continue
		}

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

		// if zlib compressed, then swap reader out with uncompressed section
		// FIXME: this looks like ass but idk how to do it correctly
		if record.isCompressed() {

			// get size of decompressed data
			bufDataDecompSize := make([]byte, 4)

			// get compressed data
			bufDataComp := make([]byte, record.dataSize - 4)
			if _, err := sr.ReadAt(bufDataDecompSize, 0); err != nil {
				return err
			}
			bDataDecompSize := readBuf(bufDataDecompSize[:])

			if _, err := sr.ReadAt(bufDataComp, 4); err != nil {
				return err
			}

			// set record size to the new decomp size
			record.dataSize = bDataDecompSize.uint32()

			bDataComp := bufDataComp[:]

			bCompReader := bytes.NewReader(bDataComp)

			readCloserDecomp, err := zlib.NewReader(bCompReader)

			if err != nil{
				panic(err)
			}

			byteDecomp, err := ioutil.ReadAll(readCloserDecomp)

			readerDecomp := bytes.NewReader(byteDecomp)

			sr = io.NewSectionReader(readerDecomp, 0, int64(record.dataSize))

			// record.readFields() skips over the record header bytes, so set it to negative that
			record.off = -1 * recordHeaderLen

			currentReader = readerDecomp

		}

		err = record.readFields(currentReader)

		if err != nil {
			return err
		}

		g.records = append(g.records, record)
	}

	return nil





}

func (g *Group) GetRecords() []*Record {
	records := make([]*Record, 0)

	for _, subGroup := range g.subGroups {
		groupRecords := subGroup.GetRecords()
		records = append(records, groupRecords...)
	}

	records = append(records, g.records...)

	return records
}
















