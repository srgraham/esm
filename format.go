package esm

var FieldsStructLookup map[string]map[string]interface{}

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

}

