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
	readerAt         io.ReaderAt
	zipsize      int64
	//headerOffset int64
	records []*Record
}


func (g *Group) readHeader(r io.Reader) error {
	var buf [recordHeaderLen]byte
	if _, err := io.ReadFull(r, buf[:]); err != nil {
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


















