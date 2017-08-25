package esm

import (
	"fmt"
	"io"
)

//import (
//	"fmt"
//	"log"
//)
//
//func main(){
//	fmt.Printf("yoooo")
//	r, err := OpenReader("./ShellRain.esp")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer r.Close()
//}


func Peek(obj interface{}, off int64) {
	var out string;
	switch obj.(type) {
	case readBuf:
		objReadBuf := obj.(readBuf)
		out = (objReadBuf).Human()
		break
	case []byte: //, []uint8:
		objReadBuf := readBuf(obj.([]byte)[:])
		Peek(objReadBuf, off)
		return
	case io.ReaderAt:
		dumpBuff := make([]byte, 80)
		if _, err := obj.(io.ReaderAt).ReadAt(dumpBuff, off); err != nil {
			panic(err)
		}
		Peek(dumpBuff, off)
		return
	case io.SectionReader:
		sr := obj.(io.SectionReader)
		dumpBuff := make([]byte, 80)
		sr.ReadAt(dumpBuff, off)
		Peek(dumpBuff, 0)
		return
	}
	fmt.Printf("PEEK: %s\n", out)
	return
}