package esm

import (
	"fmt"
	//"reflect"
	"reflect"
)

var FieldsStructLookup map[string]map[string]interface{}

func init() {

	var boolZero bool
	var uint8Zero uint8
	var uint16Zero uint16
	var uint32Zero uint32
	var uint64Zero uint64
	var int8Zero int8
	var int16Zero int16
	var int32Zero int32
	var int64Zero int64
	var float32Zero float32
	var float64Zero float64

	var nullZero null
	var charZero char
	var char4Zero char4
	var wcharZero wchar

	var vsvalZero vsval

	var formidZero formid
	var irefZero iref
	var hashZero hash
	var filetimeZero filetime
	var systemtimeZero systemtime
	var rgbZero rgb

	var lstringZero lstring
	var dlstringZero dlstring
	var ilstringZero ilstring
	var bstringZero bstring
	var bzstringZero bzstring
	var wstringZero wstring
	var wzstringZero wzstring
	var zstringZero zstring



	_ = boolZero
	_ = uint8Zero
	_ = uint16Zero
	_ = uint32Zero
	_ = uint64Zero
	_ = int8Zero
	_ = int16Zero
	_ = int32Zero
	_ = int64Zero
	_ = float32Zero
	_ = float64Zero
	_ = nullZero
	_ = charZero
	_ = char4Zero
	_ = wcharZero
	_ = vsvalZero
	_ = formidZero
	_ = irefZero
	_ = hashZero
	_ = filetimeZero
	_ = systemtimeZero
	_ = rgbZero
	_ = lstringZero
	_ = dlstringZero
	_ = ilstringZero
	_ = bstringZero
	_ = bzstringZero
	_ = wstringZero
	_ = wzstringZero
	_ = zstringZero








	FieldsStructLookup = make(map[string]map[string]interface{})


	// TES4
	FieldsStructLookup["TES4"] = make(map[string]interface{})
	TES4 := FieldsStructLookup["TES4"]

	TES4["HEDR"] = struct {
		Version      float32
		NumRecords   int32
		NextObjectId uint32
	}{}
	TES4["CNAM"] = zstringZero
	TES4["SNAM"] = zstringZero
	TES4["MAST"] = zstringZero
	TES4["DATA"] = uint64Zero
	TES4["INTV"] = uint32Zero


	// GMST
	FieldsStructLookup["GMST"] = make(map[string]interface{})
	GMST := FieldsStructLookup["GMST"]

	GMST["EDID"] = zstringZero

	GMST["DATA"] = func (b readBuf, record Record) reflect.Type {
		var firstChar byte

		fieldEDID := record.fieldsByType("EDID")[0]

		var str zstring
		var ok bool

		if str, ok = fieldEDID.data.(zstring); !ok {
			panic(fmt.Errorf("Couldnt read zstring of EDID"))
		}

		firstChar = str[0]

		switch firstChar {
		case 'b':
			return reflect.TypeOf(uint32Zero)
		case 'i':
			return reflect.TypeOf(uint32Zero)
		case 'f':
			return reflect.TypeOf(float32Zero)
		case 'l':
			return reflect.TypeOf(lstringZero)
		default:
			panic(fmt.Errorf("Couldnt figure out type for char '%c'", firstChar))
		}


		return nil
	}

}