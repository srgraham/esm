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

	var zero_zstring zstring

	FieldsStructLookup = make(map[string]map[string]interface{})
	//FieldsStructLookup = map[string]interface{} {
	FieldsStructLookup["TES4"] = make(map[string]interface{})
	FieldsStructLookup["TES4"]["HEDR"] = struct {
		Version      float32
		NumRecords   int32
		NextObjectId uint32
	}{}

	FieldsStructLookup["TES4"]["CNAM"] = zero_zstring
	FieldsStructLookup["TES4"]["MAST"] = zero_zstring

	_ = 9
}

