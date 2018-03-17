package esm

import (
	"io"
)


type StringFileHeader struct {
	count uint32
	dataSize uint32
}

type StringFile struct {
	StringFileHeader
	sr          *io.SectionReader
	off         int64

	//directoryEntries []*StringFileDirectoryEntry
	data []uint8
	dataBuf []byte

	// map[id from mod file] = offset in data buffer
	directoryEntries map[uint32]uint32
}

type StringFileDirectoryEntry struct {
	id uint32
	offset uint32
}

// calculates the size of all the things (header and all data)
func (sf *StringFile) Size() int64 {
	return int64(stringFileMainHeaderLen) + sf.DirectoryEntrySize() + int64(sf.dataSize)
}
func (sf *StringFile) DirectoryEntrySize() int64 {
	return int64(sf.count * stringFileDirectoryEntryLen)
}

func (sf *StringFile) read(sr io.SectionReader) error {
	buf := make([]byte, stringFileMainHeaderLen)

	if _, err := sr.Read(buf); err != nil {
		return err
	}
	b := readBuf(buf[:])

	sf.count = b.uint32()
	sf.dataSize = b.uint32()

	sf.readDirectoryEntries()

	sf.dataBuf = make([]byte, sf.dataSize)

	sr.ReadAt(sf.dataBuf, int64(stringFileMainHeaderLen) + int64(sf.DirectoryEntrySize()))

	sf.data = readBuf(sf.dataBuf[:])

	return nil
}



func (sf *StringFile) readDirectoryEntries() error {

	off := int64(stringFileMainHeaderLen)

	buf := make([]byte, sf.DirectoryEntrySize())


	if _, err := sf.sr.Read(buf); err != nil {
		return err
	}

	b := readBuf(buf[:])

	sf.directoryEntries = map[uint32]uint32{}

	for off < sf.DirectoryEntrySize() {
		id := b.uint32()
		offset := b.uint32()
		off += stringFileDirectoryEntryLen

		sf.directoryEntries[id] = offset
	}

	return nil
}

func (sf *StringFile) GetStringForId(id uint32) string {
	offset := sf.directoryEntries[id]

	i := offset
	for (sf.data)[i] != 0 {
		i += 1
	}
	x := zstring(sf.data[offset:i])
	return AsString(x)
}
