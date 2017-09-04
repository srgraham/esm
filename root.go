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
		var group *Group
		var err error
		group, off, err = root.readNextGroup(reader, off)

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}

		if group == nil {
			fmt.Printf("yo")
		}

		root.groups = append(root.groups, group)
	}

	return nil
}


func (root *Root) readNextGroup(reader io.ReaderAt, off int64) (*Group, int64, error) {
	headerReader := io.NewSectionReader(reader, off, groupHeaderLen)

	group := &Group{parentRoot: root, off: off}
	err := group.readHeader(*headerReader)

	off += groupHeaderLen

	sr := io.NewSectionReader(reader, off, group.Size() - groupHeaderLen)
	group.sr = sr

	off += group.Size() - groupHeaderLen

	if err != nil {
		return group, off, err
	}

	switch group.groupType {
	case 0: // top
		if isAllowedGroupType, ok := root.allowedGroupTypes[group.Type()]; ok && isAllowedGroupType {
			err = group.readRecords(reader)
			if err != nil {
				return group, off, err
			}
		} else {
			fmt.Printf("Skipping GROUP read for %d'%s'\n", group.groupType, group.Type())
			if _, ok := FieldsStructLookup[group.Type()]; !ok {
				return group, off, fmt.Errorf("Missing definition for GROUP record type %s", group.Type())
				//fmt.Printf("Missing definition for GROUP record type %s\n\n", group.Type())
			}
		}
		return group, off, nil

	//case 1: // world children
	//case 2: // interior cell block
	//case 3: // interior cell subblock
	//case 4: // exterior cell block
	//case 5: // exterior cell subblock
	//case 6: // cell children
	//case 7: // topic children
	//case 8: // cell persistent children
	//case 9: // cell temporary children
	//case 10: // cell visible distant children
	default:
		err = group.readRecords(reader)
		if err != nil {
			return group, off, err
		}
	}
	return group, off, nil
}

func (root *Root) GetRecords() []*Record {
	records := make([]*Record, 0)

	for _, group := range root.groups {
		groupRecords := group.GetRecords()
		records = append(records, groupRecords...)
	}

	return records
}

func (root *Root) GetRecordsOfType(_type string) []*Record {
	allRecords := root.GetRecords()
	records := make([]*Record, 0)
	for _, record := range allRecords {
		if record.Type() == _type {
			records = append(records, record)
		}
	}
	return records
}


func (root *Root) GetRecordByFormId(id formid) *Record {
	if record, ok := FormIds[id].(*Record); ok {
		return record
	}
	return nil
}


func (root *Root) GetRecordByEdid(edid string) *Record {
	id := EdidIds[edid]
	return root.GetRecordByFormId(id)
}

