package esm

import (
	"io"

	"fmt"
)

type Root struct {
	rootRecord *Record
	sr *io.SectionReader
	r *io.Reader
	off int64
	readerAt         io.ReaderAt
	readerSize      int64
	groups []*Group
	allowedGroupTypes map[string]bool
}


func (root *Root) Size() int64 {
	return root.readerSize
}

func (root *Root) IsLocalized() bool {
	return root.rootRecord.isLocalized()
}


func (root *Root) AllowGroup(groupType string) {
	if root.allowedGroupTypes == nil {
		root.allowedGroupTypes = make(map[string]bool)
	}
	root.allowedGroupTypes[groupType] = true
}
func (root *Root) DisallowGroup(groupType string) {
	if root.allowedGroupTypes == nil {
		root.allowedGroupTypes = make(map[string]bool)
	}
	root.allowedGroupTypes[groupType] = false
}
func (root *Root) AllowAllGroups() {
	if root.allowedGroupTypes == nil {
		root.allowedGroupTypes = make(map[string]bool)
	}
	for groupType, _ := range FieldsStructLookup {
		root.allowedGroupTypes[groupType] = true
	}
}
func (root *Root) DisallowAllGroups() {
	if root.allowedGroupTypes == nil {
		root.allowedGroupTypes = make(map[string]bool)
	}
	for groupType, _ := range FieldsStructLookup {
		root.allowedGroupTypes[groupType] = false
	}
}


func (root *Root) readGroups(reader io.ReaderAt) error {

	// read from the start of the file + recordSize depth to start reading the groups
	groupsSr := io.NewSectionReader(root.readerAt, root.rootRecord.Size(), root.Size())

	if _, err := groupsSr.Seek(0, io.SeekStart); err != nil {
		return err
	}

	root.groups = make([]*Group, 0)

	var off int64 = root.off + root.rootRecord.Size()

	for off < root.off + root.Size() {
		headerReader := io.NewSectionReader(reader, off, groupHeaderLen)

		group := &Group{parentRoot: root, off: off}
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

		if isAllowedGroupType, ok := root.allowedGroupTypes[group.Type()]; ok && isAllowedGroupType {
			err = group.readRecords(reader)
			if err != nil {
				return err
			}
		} else {
			fmt.Printf("Skipping GROUP read for '%s'\n", group.Type())
			if _, ok := FieldsStructLookup[group.Type()]; !ok {
				return fmt.Errorf("Missing definition for GROUP record type %s", group.Type())
				//fmt.Printf("Missing definition for GROUP record type %s\n\n", group.Type())
			}
		}

		root.groups = append(root.groups, group)
	}

	return nil
}