package esm

import (
	"fmt"
	//"reflect"
	"reflect"
	"sort"
)

var FieldsStructLookup map[string]map[string]interface{}

var UnimplementedFields map[string]map[string][]readBuf

var FormIds map[formid]interface{}

// make an interface that no other types will match
type Skip bool
type Unknown bool

// no * here so that GoString() works right
func (s Skip) GoString() string {
	return "SKIP"
}
// no * here so that GoString() works right
func (u Unknown) GoString() string {
	return "UNKNOWN"
}

var SkipZero Skip
var UnknownZero Unknown

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

type OBND struct {
	X1 int16
	Y1 int16
	Z1 int16
	X2 int16
	Y2 int16
	Z2 int16
}

type PosRot struct{
	Position struct {
		X float32
		Y float32
		Z float32
	}
	Rotation struct {
		X float32
		Y float32
		Z float32
	}
}


func MakeFieldStruct(label string) map[string]interface{} {
	FieldsStructLookup[label] = make(map[string]interface{})

	// also set base values

	FieldsStructLookup[label]["EDID"] = zstringZero
	FieldsStructLookup[label]["FULL"] = lstringZero

	FieldsStructLookup[label]["KSIZ"] = uint32Zero

	FieldsStructLookup[label]["KWDA"] = func (b readBuf, record Record) interface{} {

		fieldKSIZ := record.fieldsByType("KSIZ")[0]

		var count uint32
		var ok bool

		if count, ok = fieldKSIZ.data.(uint32); !ok {
			panic(fmt.Errorf("Couldnt read count of KSIZ for KWDA"))
		}

		kwdas := make([]formid, count)

		for i := uint32(0); i < count; i += 1 {
			kwdas[i] = b.formid()
		}

		return kwdas
	}


	FieldsStructLookup[label]["MODL"] = zstringZero

	FieldsStructLookup[label]["OBND"] = OBND{}








	return FieldsStructLookup[label]
}

func LogUnimplementedField(recordTypeStr string, fieldTypeStr string, dataBuf readBuf) {

	if _, ok := UnimplementedFields[recordTypeStr]; !ok {
		UnimplementedFields[recordTypeStr] = make(map[string][]readBuf)
	}
	if _, ok := UnimplementedFields[recordTypeStr][fieldTypeStr]; !ok {
		UnimplementedFields[recordTypeStr][fieldTypeStr] = make([]readBuf, 0)
	}
	UnimplementedFields[recordTypeStr][fieldTypeStr] = append(UnimplementedFields[recordTypeStr][fieldTypeStr], dataBuf)
}

func DumpUnimplementedFields() {
	fmt.Println("---Unimplemented Fields---")

	keys := make([]string, 0)
	for k, _ := range UnimplementedFields{
		keys = append(keys, k)
	}
	sort.Strings(keys)



	for _, recordType := range keys {
		recordFields := UnimplementedFields[recordType]

		var fieldTypes []string
		for fieldType, fieldBufs := range recordFields {
			fieldInfo := fmt.Sprintf("%s(%d)", fieldType, len(fieldBufs))
			fieldTypes = append(fieldTypes, fieldInfo)
		}
		fmt.Printf("%s: %s\n", recordType, fieldTypes)
	}
}

func dumpAndCrash(b readBuf, record Record) reflect.Type {
	fmt.Printf(
		"### Dump and crash :D ###\nBuffer(%d): %s\n\n%s\n\n",
		b.Size(),
		b.Human(),
		record.String(),
	)
	panic(":D")
}

func DumpFormIds() {
	fmt.Println("---Form IDs---")
	for formId, ref := range FormIds {
		fmt.Printf("%08x: %s\n", formId, ref)
	}
}


func init() {







	UnimplementedFields = make(map[string]map[string][]readBuf)
	FieldsStructLookup = make(map[string]map[string]interface{})
	FormIds = make(map[formid]interface{})




	/* TES4 */
	TES4 := MakeFieldStruct("TES4")
	TES4["HEDR"] = struct {
		Version      float32
		NumRecords   int32
		NextObjectId uint32
	}{}
	TES4["CNAM"] = zstringZero
	TES4["SNAM"] = zstringZero
	TES4["MAST"] = zstringZero
	TES4["DATA"] = uint64Zero
	TES4["INTV"] = UnknownZero
	TES4["INCC"] = UnknownZero
	TES4["TNAM"] = UnknownZero


	/* AACT */
	AACT := MakeFieldStruct("AACT")
	_ = AACT

	/* ACTI */
	ACTI := MakeFieldStruct("ACTI")
	_ = ACTI

	/* ADDN */
	ADDN := MakeFieldStruct("ADDN")
	_ = ADDN

	/* AECH */
	AECH := MakeFieldStruct("AECH")
	_ = AECH

	/* ALCH */
	ALCH := MakeFieldStruct("ALCH")
	_ = ALCH

	/* AMDL */
	AMDL := MakeFieldStruct("AMDL")
	_ = AMDL

	/* AMMO */
	AMMO := MakeFieldStruct("AMMO")
	_ = AMMO

	/* ANIO */
	ANIO := MakeFieldStruct("ANIO")
	_ = ANIO

	/* AORU */
	AORU := MakeFieldStruct("AORU")
	_ = AORU

	/* ARMA */
	ARMA := MakeFieldStruct("ARMA")
	_ = ARMA

	/* ARMO */
	ARMO := MakeFieldStruct("ARMO")
	_ = ARMO

	/* ARTO */
	ARTO := MakeFieldStruct("ARTO")
	_ = ARTO

	/* ASPC */
	ASPC := MakeFieldStruct("ASPC")
	_ = ASPC

	/* ASTP */
	ASTP := MakeFieldStruct("ASTP")
	_ = ASTP

	/* AVIF */
	AVIF := MakeFieldStruct("AVIF")
	_ = AVIF

	/* BNDS */
	BNDS := MakeFieldStruct("BNDS")
	_ = BNDS

	/* BOOK */
	BOOK := MakeFieldStruct("BOOK")
	_ = BOOK

	/* BPTD */
	BPTD := MakeFieldStruct("BPTD")
	_ = BPTD

	/* CAMS */
	CAMS := MakeFieldStruct("CAMS")
	_ = CAMS

	/* CELL */
	CELL := MakeFieldStruct("CELL")
	_ = CELL

	/* CLAS */
	CLAS := MakeFieldStruct("CLAS")
	_ = CLAS

	/* CLFM */
	CLFM := MakeFieldStruct("CLFM")
	_ = CLFM

	/* CLMT */
	CLMT := MakeFieldStruct("CLMT")
	_ = CLMT

	/* CMPO */
	CMPO := MakeFieldStruct("CMPO")
	_ = CMPO

	/* COBJ */
	COBJ := MakeFieldStruct("COBJ")
	_ = COBJ

	/* COLL */
	COLL := MakeFieldStruct("COLL")
	_ = COLL

	/* CONT */
	CONT := MakeFieldStruct("CONT")
	_ = CONT

	/* CPTH */
	CPTH := MakeFieldStruct("CPTH")
	_ = CPTH

	/* CSTY */
	CSTY := MakeFieldStruct("CSTY")
	_ = CSTY

	/* DEBR */
	DEBR := MakeFieldStruct("DEBR")
	_ = DEBR

	/* DFOB */
	DFOB := MakeFieldStruct("DFOB")
	_ = DFOB

	/* DLVW */
	DLVW := MakeFieldStruct("DLVW")
	_ = DLVW

	/* DMGT */
	DMGT := MakeFieldStruct("DMGT")
	_ = DMGT

	/* DOBJ */
	DOBJ := MakeFieldStruct("DOBJ")
	_ = DOBJ

	/* DOOR */
	DOOR := MakeFieldStruct("DOOR")
	_ = DOOR

	/* ECZN */
	ECZN := MakeFieldStruct("ECZN")
	_ = ECZN

	/* EFSH */
	EFSH := MakeFieldStruct("EFSH")
	_ = EFSH

	/* ENCH */
	ENCH := MakeFieldStruct("ENCH")
	_ = ENCH

	/* EQUP */
	EQUP := MakeFieldStruct("EQUP")
	_ = EQUP

	/* EXPL */
	EXPL := MakeFieldStruct("EXPL")
	_ = EXPL

	/* FACT */
	FACT := MakeFieldStruct("FACT")
	_ = FACT

	/* FLOR */
	FLOR := MakeFieldStruct("FLOR")
	_ = FLOR

	/* FLST */
	FLST := MakeFieldStruct("FLST")
	_ = FLST

	/* FSTP */
	FSTP := MakeFieldStruct("FSTP")
	_ = FSTP

	/* FSTS */
	FSTS := MakeFieldStruct("FSTS")
	_ = FSTS

	/* FURN */
	FURN := MakeFieldStruct("FURN")
	FURN["FULL"] = lstringZero
	FURN["FULL"] = lstringZero




	/* GDRY */
	GDRY := MakeFieldStruct("GDRY")
	_ = GDRY

	/* GLOB */
	GLOB := MakeFieldStruct("GLOB")
	_ = GLOB

	/* GMST */
	GMST := MakeFieldStruct("GMST")
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

	/* GRAS */
	GRAS := MakeFieldStruct("GRAS")
	_ = GRAS


	/* HAZD */
	HAZD := MakeFieldStruct("HAZD")
	_ = HAZD

	/* HDPT */
	HDPT := MakeFieldStruct("HDPT")
	_ = HDPT

	/* IDLE */
	IDLE := MakeFieldStruct("IDLE")
	_ = IDLE

	/* IDLM */
	IDLM := MakeFieldStruct("IDLM")
	_ = IDLM

	/* IMAD */
	IMAD := MakeFieldStruct("IMAD")
	_ = IMAD

	/* IMGS */
	IMGS := MakeFieldStruct("IMGS")
	_ = IMGS

	/* INGR */
	INGR := MakeFieldStruct("INGR")
	_ = INGR

	/* INNR */
	INNR := MakeFieldStruct("INNR")
	_ = INNR

	/* IPCT */
	IPCT := MakeFieldStruct("IPCT")
	_ = IPCT

	/* IPDS */
	IPDS := MakeFieldStruct("IPDS")
	_ = IPDS

	/* KEYM */
	KEYM := MakeFieldStruct("KEYM")
	_ = KEYM

	/* KSSM */
	KSSM := MakeFieldStruct("KSSM")
	_ = KSSM

	/* KYWD */
	KYWD := MakeFieldStruct("KYWD")
	_ = KYWD

	/* LAYR */
	LAYR := MakeFieldStruct("LAYR")
	_ = LAYR

	/* LAYZ */
	LAYZ := MakeFieldStruct("LAYZ")
	_ = LAYZ

	/* LCRT */
	LCRT := MakeFieldStruct("LCRT")
	_ = LCRT

	/* LCTN */
	LCTN := MakeFieldStruct("LCTN")
	_ = LCTN

	/* LENS */
	LENS := MakeFieldStruct("LENS")
	_ = LENS

	/* LGTM */
	LGTM := MakeFieldStruct("LGTM")
	_ = LGTM

	/* LIGH */
	LIGH := MakeFieldStruct("LIGH")
	_ = LIGH

	/* LSCR */
	LSCR := MakeFieldStruct("LSCR")
	_ = LSCR

	/* LTEX */
	LTEX := MakeFieldStruct("LTEX")
	_ = LTEX

	/* LVLI */
	LVLI := MakeFieldStruct("LVLI")
	_ = LVLI

	/* LVLN */
	LVLN := MakeFieldStruct("LVLN")
	_ = LVLN

	/* MATO */
	MATO := MakeFieldStruct("MATO")
	_ = MATO

	/* MATT */
	MATT := MakeFieldStruct("MATT")
	_ = MATT

	/* MESG */
	MESG := MakeFieldStruct("MESG")
	MESG["DESC"] = lstringZero
	MESG["INAM"] = uint32Zero
	//MESG["QNAM"] = formidZero
	MESG["DNAM"] = uint32Zero

	/* MGEF */
	MGEF := MakeFieldStruct("MGEF")
	_ = MGEF

	/* MISC */
	MISC := MakeFieldStruct("MISC")
	_ = MISC

	/* MOVT */
	MOVT := MakeFieldStruct("MOVT")
	_ = MOVT

	/* MSTT */
	MSTT := MakeFieldStruct("MSTT")
	_ = MSTT

	/* MSWP */
	MSWP := MakeFieldStruct("MSWP")
	_ = MSWP

	/* MUSC */
	MUSC := MakeFieldStruct("MUSC")
	_ = MUSC

	/* MUST */
	MUST := MakeFieldStruct("MUST")
	_ = MUST

	/* NAVI */
	NAVI := MakeFieldStruct("NAVI")
	_ = NAVI

	/* NOCM */
	NOCM := MakeFieldStruct("NOCM")
	_ = NOCM

	/* NOTE */
	NOTE := MakeFieldStruct("NOTE")
	_ = NOTE

	/* NPC_ */
	NPC_ := MakeFieldStruct("NPC_")
	_ = NPC_

	/* OMOD */
	OMOD := MakeFieldStruct("OMOD")
	_ = OMOD

	/* OTFT */
	OTFT := MakeFieldStruct("OTFT")
	_ = OTFT

	/* OVIS */
	OVIS := MakeFieldStruct("OVIS")
	_ = OVIS

	/* PACK */
	PACK := MakeFieldStruct("PACK")
	_ = PACK

	/* PERK */
	PERK := MakeFieldStruct("PERK")
	_ = PERK

	/* PGRE */
	PGRE := MakeFieldStruct("PGRE")
	_ = PGRE

	/* PHZD */
	PHZD := MakeFieldStruct("PHZD")
	_ = PHZD

	/* PKIN */
	PKIN := MakeFieldStruct("PKIN")
	_ = PKIN

	/* PROJ */
	PROJ := MakeFieldStruct("PROJ")
	_ = PROJ

	/* QUST */
	QUST := MakeFieldStruct("QUST")
	_ = QUST

	/* RACE */
	RACE := MakeFieldStruct("RACE")
	_ = RACE

	/* REFR */
	REFR := MakeFieldStruct("REFR")
	REFR["NAME"] = formidZero
	REFR["XPRM"] = struct {
		X float32
		Y float32
		Z float32
		Unknown float32
		Unknown2 uint32
	}{}
	REFR["XLRM"] = formidZero

	type structXWCU struct {
		X float32
		Y float32
		Z float32
		Unknown float32
	}

	REFR["XWCU"] = struct {
		LinearVelocity structXWCU
		AngularVelocity structXWCU
		Empty structXWCU
	}{}

	REFR["DATA"] = PosRot{}

	REFR["XLIG"] = SkipZero
	REFR["XLYR"] = formidZero // GECK layer info (house floor grouping) :D
	REFR["XTEL"] = struct {
		Door formid
		PosRot
		Flags uint32
	}{}

	REFR["XNDP"] = struct {
		NavMesh formid
		TeleportMarkerTriangle int16
		Unused [2]byte
	}{}

	REFR["XSCL"] = float32Zero

	//REFR["XLOC"] = struct {
	//	Level byte
	//	Flags1 [3]byte
	//	KEYM formid
	//	Flags2 byte
	//	Flags3 [3]byte
	//	Flags4 [8]byte
	//
	//}{}
	//REFR["XLOC"] = dumpAndCrash
	_ = REFR

	/* REGN */
	REGN := MakeFieldStruct("REGN")
	_ = REGN

	/* RELA */
	RELA := MakeFieldStruct("RELA")
	_ = RELA

	/* REVB */
	REVB := MakeFieldStruct("REVB")
	_ = REVB

	/* RFCT */
	RFCT := MakeFieldStruct("RFCT")
	_ = RFCT

	/* RFGP */
	RFGP := MakeFieldStruct("RFGP")
	_ = RFGP

	/* SCCO */
	SCCO := MakeFieldStruct("SCCO")
	_ = SCCO

	/* SCOL */
	SCOL := MakeFieldStruct("SCOL")
	_ = SCOL

	/* SCSN */
	SCSN := MakeFieldStruct("SCSN")
	_ = SCSN

	/* SMBN */
	SMBN := MakeFieldStruct("SMBN")
	_ = SMBN

	/* SMEN */
	SMEN := MakeFieldStruct("SMEN")
	_ = SMEN

	/* SMQN */
	SMQN := MakeFieldStruct("SMQN")
	_ = SMQN

	/* SNCT */
	SNCT := MakeFieldStruct("SNCT")
	_ = SNCT

	/* SNDR */
	SNDR := MakeFieldStruct("SNDR")
	_ = SNDR

	/* SOPM */
	SOPM := MakeFieldStruct("SOPM")
	_ = SOPM

	/* SOUN */
	SOUN := MakeFieldStruct("SOUN")
	_ = SOUN

	/* SPEL */
	SPEL := MakeFieldStruct("SPEL")
	_ = SPEL

	/* SPGD */
	SPGD := MakeFieldStruct("SPGD")
	_ = SPGD

	/* STAG */
	STAG := MakeFieldStruct("STAG")
	_ = STAG

	/* STAT */
	STAT := MakeFieldStruct("STAT")
	_ = STAT

	/* TACT */
	TACT := MakeFieldStruct("TACT")
	_ = TACT

	/* TERM */
	TERM := MakeFieldStruct("TERM")
	_ = TERM

	/* TREE */
	TREE := MakeFieldStruct("TREE")
	_ = TREE

	/* TRNS */
	TRNS := MakeFieldStruct("TRNS")
	_ = TRNS

	/* TXST */
	TXST := MakeFieldStruct("TXST")
	_ = TXST

	/* VTYP */
	VTYP := MakeFieldStruct("VTYP")
	_ = VTYP

	/* WATR */
	WATR := MakeFieldStruct("WATR")
	_ = WATR

	/* WEAP */
	WEAP := MakeFieldStruct("WEAP")
	_ = WEAP

	/* WRLD */
	WRLD := MakeFieldStruct("WRLD")
	WRLD["ICON"] = zstringZero
	WRLD["XWEM"] = zstringZero
	_ = WRLD

	/* WTHR */
	WTHR := MakeFieldStruct("WTHR")
	_ = WTHR

	/* ZOOM */
	ZOOM := MakeFieldStruct("ZOOM")
	_ = ZOOM





}