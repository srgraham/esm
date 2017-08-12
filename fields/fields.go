package fields

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


//type things map[string]map[string]interface{} {
//	"TES4": map[string]interface{} {
//		"HEDR": thing
//	}
//}


func init() {

	FieldsStructLookup = make(map[string]map[string]interface{})
	//FieldsStructLookup = map[string]interface{} {
	FieldsStructLookup["TES4"] = make(map[string]interface{})
	FieldsStructLookup["TES4"]["HEDR"] = struct {
		Version      float32
		NumRecords   int32
		NextObjectId uint32
	}{}

	_ = 9
}

//FieldsStructLookup map[string]interface{} {
//	TES4 struct {
//		HEDR struct {
//			Version float32
//			NumRecords int32
//			NextObjectId uint32
//		}
//	}
//}


