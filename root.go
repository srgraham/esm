package esm

import (
	"io"
	"bufio"
)

type Root struct {
	rootRecord *Record
	groups []*Group
}



func (r *Record) readGroups() error {

	rs := io.NewSectionReader(r.readerAt, groupHeaderLen, int64(r.dataSize) + int64(recordHeaderLen))
	if _, err := rs.Seek(0, io.SeekStart); err != nil {
		return err
	}

	reader := bufio.NewReader(rs)

	r.fields = make([]*Field, 0)

	// process all field headers
	for {

		field := &Field{record: r, readerAt: rs, zipsize: int64(r.dataSize)}
		err := field.readHeader(reader)

		if err == ErrFormat || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			return err
		}

		//field.dataSectionReader := io.NewSectionReader(r.readerAt, rs.Seek(), 1000) //int64(field.dataSize) + int64(fieldHeaderLen))
		//
		////reader := bufio.NewReader(rs)
		//
		//field.data = make([]byte, int64(field.dataSize + 20))
		//
		////var buf [recordHeaderLen]byte
		//if _, err := field.dataSectionReader.Read(field.data); err != nil {
		//	return err
		//}

		//fmt.Println(field.data)

		field.readData()

		r.fields = append(r.fields, field)
		// skip past rest of data

		//buf.Discard(int(field.dataSize))

		//fmt.Println(field)

		//buf = buf[field.dataSize:]
	}

	// at this point, only the headers for each field has been grabbed

	// now process the data for each field

	//for field := range r.fields {
	//	field.readData()
	//}

	return nil
}