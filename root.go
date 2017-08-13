package esm

import (
	"io"

)

type Root struct {
	rootRecord *Record
	sr *io.SectionReader
	r *io.Reader
	off int64
	readerAt         io.ReaderAt
	readerSize      int64
	groups []*Group
}


func (root *Root) Size() int64 {
	return root.readerSize
}


func (root *Root) readGroups(reader io.ReaderAt) error {

	// read from the start of the file + recordSize depth to start reading the groups
	groupsSr := io.NewSectionReader(root.readerAt, root.rootRecord.Size(), 1<<63-1)

	if _, err := groupsSr.Seek(0, io.SeekStart); err != nil {
		return err
	}

	root.groups = make([]*Group, 0)

	var off int64 = root.off + root.rootRecord.Size()

	for {
		headerReader := io.NewSectionReader(reader, off, groupHeaderLen)

		group := &Group{root: root, off: off}
		err := group.readHeader(*headerReader)

		off += groupHeaderLen

		sr := io.NewSectionReader(reader, off, group.Size() - groupHeaderLen)
		group.sr = sr

		off += group.Size() - groupHeaderLen

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		group.readRecords(reader)

		root.groups = append(root.groups, group)
	}

	return nil
}