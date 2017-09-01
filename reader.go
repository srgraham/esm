package esm


import (
"encoding/binary"
"errors"
_ "fmt"
"io"
"os"

)

var (
	ErrFormat = errors.New("esm: not a valid esm/esp file")
	ErrNilType = errors.New("cannot decode nil type")
	ErrDecodeUnknownType = errors.New("cannot decode %s type")
	ErrUnimplementedField = errors.New("Unimplemented field")
	ErrRecordIsGRUP = errors.New("Record is being read as a GRUP")
)

type Reader struct {
	r             io.ReaderAt
	Record        *Record
	Comment       string
	//decompressors map[uint16]Decompressor
}

type ReadCloser struct {
	f *os.File
	Reader
}



// OpenReader will open the esm/esp file specified by name and return a ReadCloser.
func OpenReader(name string, allowedGroups []string) (*ReadCloser, *Root, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	r := new(ReadCloser)

	var root *Root
	var err2 error

	if root, err2 = r.init(f, fi.Size(), allowedGroups); err2 != nil {
		f.Close()
		return nil, root, err2
	}
	r.f = f
	return r, root, nil
}

// NewReader returns a new Reader reading from r, which is assumed to
// have the given size in bytes.
func NewReader(readerAt io.ReaderAt, size int64, allowedGroups []string) (*Reader, *Root, error) {
	var root *Root
	var err error

	reader := new(Reader)

	if root, err = reader.init(readerAt, size, allowedGroups); err != nil {
		return nil, root, err
	}
	return reader, root, nil
}

func (z *Reader) init(reader io.ReaderAt, size int64, allowedGroups []string) (*Root, error) {
	if size == 0 {
		size = 1<<63-1
	}
	//end, err := readDirectoryEnd(r, size)
	//if err != nil {
	//	return err
	//}
	//if end.directoryRecords > uint64(size)/fileHeaderLen {
	//	return fmt.Errorf("archive/zip: TOC declares impossible %d files in %d byte zip", end.directoryRecords, size)
	//}
	//z.r = reader

	off := int64(0)

	sr := io.NewSectionReader(reader, off, size)
	//if _, err = sr.Seek(int64(end.directoryOffset), io.SeekStart); err != nil {
	//	return err
	//}
	//reader := bufio.NewReader(sr)

	rootRecord := &Record{off: off}

	err := rootRecord.readHeader(*sr)

	if err != nil {
		return nil, err
	}

	// check its a TES4 parentRecord
	if binary.BigEndian.Uint32([]byte(rootRecord._type[:])) != fileHeaderSignature {
		return nil, ErrFormat
	}
	// make sure its the right form id (0)
	if rootRecord.formid != 0 {
		return nil, ErrFormat
	}

	// read the fields of the parentRoot parentRecord
	err = rootRecord.readFields(reader)

	if err != nil {
		return nil, err
	}

	// we have the TES4 data. now lets grab the groups
	root := &Root{rootRecord : rootRecord, readerAt: reader, readerSize: size, off: off}

	if allowedGroups != nil {
		for _, allowedGroupType := range allowedGroups {
			root.AllowGroup(allowedGroupType)
		}
	} else {
		root.AllowAllGroups()
	}

	//root.AllowAllGroups()
	//root.DisallowGroup("WRLD")
	//root.DisallowGroup("CELL")

	err = root.readGroups(reader)
	if err != nil {
		return root, err
	}

	//DumpUnimplementedFields()
	//DumpFormIds()

	return root, nil

}


// Close closes the Zip file, rendering it unusable for I/O.
func (rc *ReadCloser) Close() error {
	return rc.f.Close()
}



