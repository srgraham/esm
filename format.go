package esm

import (
	"fmt"
	//"reflect"
	"reflect"
)

var FieldsStructLookup map[string]map[string]interface{}

var UnimplementedFields map[string]map[string]bool

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


func MakeFieldStruct(label string) map[string]interface{} {
	FieldsStructLookup[label] = make(map[string]interface{})

	// also set base values

	FieldsStructLookup[label]["EDID"] = zstringZero

	return FieldsStructLookup[label]
}

func LogUnimplementedField(recordTypeStr string, fieldTypeStr string, dataBuf readBuf) {

	if _, ok := UnimplementedFields[recordTypeStr]; !ok {
		UnimplementedFields[recordTypeStr] = make(map[string]bool)
	}
	UnimplementedFields[recordTypeStr][fieldTypeStr] = true
}

func DumpUnimplementedFields() {
	fmt.Println("---Unimplemented Fields---")
	for recordType, recordFields := range UnimplementedFields {
		var fieldTypes []string
		for fieldType, _ := range recordFields {
			fieldTypes = append(fieldTypes, fieldType)
		}
		fmt.Printf("%s: %s\n", recordType, fieldTypes)
	}
}


func init() {







	UnimplementedFields = make(map[string]map[string]bool)
	FieldsStructLookup = make(map[string]map[string]interface{})

	TES4 := MakeFieldStruct("TES4")
	GMST := MakeFieldStruct("GMST")
	MESG := MakeFieldStruct("MESG")





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
		case 'u':
			return reflect.TypeOf(uint32Zero)
		case 'f':
			return reflect.TypeOf(float32Zero)
		case 's':
			return reflect.TypeOf(lstringZero)
		default:
			panic(fmt.Errorf("Couldnt figure out type for char '%c'", firstChar))
		}
		return nil
	}

	// MESG
	MESG["DESC"] = lstringZero
	MESG["INAM"] = uint32Zero
	//MESG["QNAM"] = formidZero
	MESG["DNAM"] = uint32Zero


}