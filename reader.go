package esm


import (
"bufio"
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





//func (record *Record) hasDataDescriptor() bool {
//	return f.Flags&0x8 != 0
//}

// OpenReader will open the esm/esp file specified by name and return a ReadCloser.
func OpenReader(name string) (*ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		f.Close()
		return nil, err
	}
	r := new(ReadCloser)
	if err := r.init(f, fi.Size()); err != nil {
		f.Close()
		return nil, err
	}
	r.f = f
	return r, nil
}

// NewReader returns a new Reader reading from r, which is assumed to
// have the given size in bytes.
func NewReader(r io.ReaderAt, size int64) (*Reader, error) {
	zr := new(Reader)
	if err := zr.init(r, size); err != nil {
		return nil, err
	}
	return zr, nil
}

func (z *Reader) init(r io.ReaderAt, size int64) error {
	//end, err := readDirectoryEnd(r, size)
	//if err != nil {
	//	return err
	//}
	//if end.directoryRecords > uint64(size)/fileHeaderLen {
	//	return fmt.Errorf("archive/zip: TOC declares impossible %d files in %d byte zip", end.directoryRecords, size)
	//}
	z.r = r

	//z.Record = readRecordHeader(r)
	//z.Comment = end.comment
	rs := io.NewSectionReader(r, 0, size)
	//if _, err = rs.Seek(int64(end.directoryOffset), io.SeekStart); err != nil {
	//	return err
	//}
	buf := bufio.NewReader(rs)

	rootRecord := &Record{readerAt: r, zipsize: size}

	err := rootRecord.readHeader(buf)

	if err != nil {
		return err
	}

	// check its a TES4 record
	if binary.BigEndian.Uint32([]byte(rootRecord._type[:])) != fileHeaderSignature {
		return ErrFormat
	}
	// make sure its the right form id (0)
	if rootRecord.id != 0 {
		return ErrFormat
	}

	// read the fields of the root record
	rootRecord.readFields()

	// we have the TES4 data. now lets grab the groups
	root := &Root{rootRecord : rootRecord}


	//root.
	_ = root

	return nil

}


// Close closes the Zip file, rendering it unusable for I/O.
func (rc *ReadCloser) Close() error {
	return rc.f.Close()
}



