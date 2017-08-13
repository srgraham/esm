package esm

var FieldsStructLookup map[string]map[string]interface{}

type thing struct {
	Version      float32
	NumRecords   int32
	NextObjectId uint32
}

type thing2 struct {
	HEDR thing
	CNAM float32
}


func init() {

	//var uint8Zero uint8
	//var uint16Zero uint16
	var uint32Zero uint32
	var uint64Zero uint64
	//var int8Zero int8
	//var int16Zero int16
	//var int32Zero int32
	//var int64Zero int64
	//var float32Zero float32
	//var float64Zero float64

	//var nullZero null
	//var charZero char
	//var char4Zero char4
	//var wcharZero wchar
	//
	//var vsvalZero vsval
	//
	//var formidZero formid
	//var irefZero iref
	//var hashZero hash
	//var filetimeZero filetime
	//var systemtimeZero systemtime
	//var rgbZero rgb
	//
	//var lstringZero lstring
	//var dlstringZero dlstring
	//var ilstringZero ilstring
	//var bstringZero bstring
	//var bzstringZero bzstring
	//var wstringZero wstring
	//var wzstringZero wzstring
	var zstringZero zstring








	FieldsStructLookup = make(map[string]map[string]interface{})
	//FieldsStructLookup = map[string]interface{} {
	FieldsStructLookup["TES4"] = make(map[string]interface{})
	FieldsStructLookup["TES4"]["HEDR"] = struct {
		Version      float32
		NumRecords   int32
		NextObjectId uint32
	}{}

	FieldsStructLookup["TES4"]["CNAM"] = zstringZero
	FieldsStructLookup["TES4"]["SNAM"] = zstringZero
	FieldsStructLookup["TES4"]["MAST"] = zstringZero
	FieldsStructLookup["TES4"]["DATA"] = uint64Zero
	FieldsStructLookup["TES4"]["INTV"] = uint32Zero

	_ = 9
}

