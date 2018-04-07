package tests

import (
	_ "fmt"
	_ "io/ioutil"
	"testing"
	"log"
	"fmt"
	esm ".."
	"encoding/json"
	"os"
)

type LODType struct {
	LOD_1 string
	LOD_2 string
	LOD_3 string
	LOD_4 string
}

type RefrStruct struct {
	FormId uint32 `json:"fid"`
	StatFormId uint32 `json:"statFid"`
	//Model string `json:"model"`
	Scale float32 `json:"scale"`
	PosX float32 `json:"posX"`
	PosY float32 `json:"posY"`
	PosZ float32 `json:"posZ"`
	RotX float32 `json:"rotX"`
	RotY float32 `json:"rotY"`
	RotZ float32 `json:"rotZ"`
	//BoundsX1 int16 `json:"boundsX1"`
	//BoundsY1 int16 `json:"boundsY1"`
	//BoundsZ1 int16 `json:"boundsZ1"`
	//BoundsX2 int16 `json:"boundsX2"`
	//BoundsY2 int16 `json:"boundsY2"`
	//BoundsZ2 int16 `json:"boundsZ2"`

	CellFid uint32 `json:"cellFid"`
}

type StatStruct struct {
	FormId uint32 `json:"fid"`
	//StatFormId uint32 `json:"statFid"`
	Model string `json:"model"`
	//Scale float32 `json:"scale"`
	//PosX float32 `json:"posX"`
	//PosY float32 `json:"posY"`
	//PosZ float32 `json:"posZ"`
	//RotX float32 `json:"rotX"`
	//RotY float32 `json:"rotY"`
	//RotZ float32 `json:"rotZ"`
	BoundsX1 int16 `json:"boundsX1"`
	BoundsY1 int16 `json:"boundsY1"`
	BoundsZ1 int16 `json:"boundsZ1"`
	BoundsX2 int16 `json:"boundsX2"`
	BoundsY2 int16 `json:"boundsY2"`
	BoundsZ2 int16 `json:"boundsZ2"`
	Lod1 string `json:"lod1"`
	Lod2 string `json:"lod2"`
	Lod3 string `json:"lod3"`
	Lod4 string `json:"lod4"`
}
type CellStruct struct {
	FormId uint32 `json:"fid"`
	EditorId string `json:"editor_id"`
	//NameLid string `json:"nameLid"`
	IsInterior bool `json:"isInterior"`
	GridX int32 `json:"grid_x"`
	GridY int32 `json:"grid_y"`
	WrldFid uint32 `json:"wrldFid"`
}
type WrldStruct struct {
	FormId uint32 `json:"fid"`
	//NameLid string `json:"nameLid"`
}
type LctnStruct struct {
	FormId uint32 `json:"fid"`
	MarkerRefrFid uint32 `json:"marker_refr_fid"`
}

func TestXxx(t *testing.T) {
	fmt.Printf("yoooo")
	//r, err := esm.OpenReader("./SkjAlert_All_DLC.esp")
	//r, err := esm.OpenReader("./ShellRain.esp")

	allowedGroupTypes := []string{

		//"AACT",
		//"ACTI",
		//"ADDN",
		//"AECH",
		//"ALCH",
		//"AMDL",
		//"AMMO",
		//"ANIO",
		//"AORU",
		//"ARMA",
		//"ARMO",
		//"ARTO",
		//"ASPC",
		//"ASTP",
		//"AVIF",
		//"BNDS",
		//"BOOK",
		//"BPTD",
		//"CAMS",
		"CELL",
		//"CLAS",
		//"CLFM",
		//"CLMT",
		//"CMPO",
		//"COBJ",
		//"COLL",
		//"CONT",
		//"CPTH",
		//"CSTY",
		//"DEBR",
		//"DFOB",
		//"DLVW",
		//"DMGT",
		//"DOBJ",
		//"DOOR",
		//"ECZN",
		//"EFSH",
		//"ENCH",
		//"EQUP",
		//"EXPL",
		//"FACT",
		//"FLOR",
		//"FLST",
		//"FSTP",
		//"FSTS",
		//"GDRY",
		//"GLOB",
		//"GRAS",
		//"GMST",
		//"HAZD",
		//"HDPT",
		//"IDLE",
		//"IDLM",
		//"IMAD",
		//"IMGS",
		//"INGR",
		//"INNR",
		//"IPCT",
		//"IPDS",
		//"KEYM",
		//"KSSM",
		//"KYWD",
		//"LAYR",
		//"LAYZ",
		//"LCRT",
		"LCTN",
		//"LENS",
		//"LGTM",
		//"LIGH",
		//"LSCR",
		//"LTEX",
		//"LVLI",
		//"LVLN",
		//"MATO",
		//"MATT",
		//"MGEF",
		//"MISC",
		//"MOVT",
		//"MSTT",
		//"MSWP",
		//"MUSC",
		//"MUST",
		//"NAVI",
		//"NOCM",
		//"NOTE",
		//"NPC_",
		//"OMOD",
		//"OTFT",
		//"OVIS",
		//"PACK",
		//"PERK",
		//"PGRE",
		//"PHZD",
		//"PKIN",
		//"PROJ",
		//"QUST",
		//"RACE",
		"REFR",
		//"REGN",
		//"RELA",
		//"REVB",
		//"RFCT",
		//"RFGP",
		//"SCCO",
		//"SCOL",
		//"SCSN",
		//"SMBN",
		//"SMEN",
		//"SMQN",
		//"SNCT",
		//"SNDR",
		//"SOPM",
		//"SOUN",
		//"SPEL",
		//"SPGD",
		//"STAG",
		"STAT",
		//"TACT",
		//"TERM",
		//"TREE",
		//"TRNS",
		//"TXST",
		//"VTYP",
		//"WATR",
		//"WEAP",
		"WRLD",
		//"WTHR",
		//"ZOOM",
	}

	//allowedGroupTypes = []string{"LCTN"}

	//allowedGroupTypes = nil

	//specificType := ""

	//specificType = "REFR"

	//if specificType != "" {
	//	allowedGroupTypes = []string{
	//		specificType,
	//	}
	//}

	//allowedGroupTypes = []string{"CELL", "WRLD"}

	r, root, err := esm.OpenReader("C:/Program Files (x86)/Steam/steamapps/common/Fallout 4/Data/Fallout4.esm", allowedGroupTypes)
	defer r.Close()

	if err != nil {
		log.Fatal(err)
	}

	esm.DumpUnimplementedFields()

	//IndustrialMachine48 := root.GetRecordByEdid("DiamondBulkheadWall12")
	//IndustrialMachine48 := root.GetRecordByFormId(0xA06E6)
	//IndustrialMachine48 := root.GetRecordByFormId(0x249c04)
	//IndustrialMachine48 := root.GetRecordByEdid("IndustrialMachine48")
	//
	//if IndustrialMachine48 == nil {
	//	esm.DumpEdidIds()
	//	panic("nil yo")
	//}

	//fmt.Printf("hello: %s", IndustrialMachine48.Dump())

	refrs := root.GetRecordsOfType("REFR")
	stats := root.GetRecordsOfType("STAT")
	cells := root.GetRecordsOfType("CELL")
	wrlds := root.GetRecordsOfType("WRLD")
	lctns := root.GetRecordsOfType("LCTN")

	refrRows := make(map[uint32]RefrStruct)
	statRows := make(map[uint32]StatStruct)
	cellRows := make(map[uint32]CellStruct)
	wrldRows := make(map[uint32]WrldStruct)
	lctnRows := make(map[uint32]LctnStruct)

	for _, refr := range refrs {

		statFormId := esm.AsUint32(refr.GetOneFieldForType("NAME").Data())

		//cellFormId := refr.Root()

		if statFormId == 0 {
			fmt.Println("Skip failed .Data() %s", statFormId)
			continue
		}

		//stat := root.GetRecordByFormIdUint32(statFormId)
		//
		//if stat == nil {
		//	fmt.Printf("nil refr for %d\n", statFormId)
		//	continue
		//}
		//
		//if stat.Type() != "STAT" {
		//	fmt.Printf("skipping non-STAT %s\n", stat.Type())
		//	continue
		//}

		//fmt.Println(refr.Dump())
		//fmt.Println(stat.Dump())



		refrDATA, _ := refr.GetFieldDataForType("DATA").(esm.PosRot)
		//statOBND, _ := stat.GetFieldDataForType("OBND").(esm.OBND)
		refrXSCL, _ := refr.GetFieldDataForType("XSCL").(float32)
		//statMODL := esm.AsString(stat.GetFieldDataForType("MODL"))
		//statMNAM, _ := stat.GetFieldDataForType("MNAM").(esm.StructLod4)



		//fmt.Println(statMNAM)
		formId := refr.FormId()

		//if formId == 761249 {
		//
		//	cellRecord := refr.NearestParentRecord()
		//	fmt.Println(cellRecord.Dump())
		//
		//	//cellFid :=
		//
		//	fmt.Println(refr.Dump())
		//	fmt.Println(refr.ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().ParentGroup().ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().Dump())
		//	fmt.Println(refr.ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().ParentGroup().Dump())
		//	//fmt.Println(refr.ParentGroup())
		//}


		//model := statMODL
		scale := refrXSCL
		posX := refrDATA.Position.X
		posY := refrDATA.Position.Y
		posZ := refrDATA.Position.Z
		rotX := refrDATA.Rotation.X
		rotY := refrDATA.Rotation.Y
		rotZ := refrDATA.Rotation.Z
		var cellFid uint32

		cellRecord := refr.NearestParentRecord()
		if cellRecord != nil {
			cellFid = cellRecord.FormId()
		}

		refrRowJson := RefrStruct{
			FormId: formId,
			StatFormId: statFormId,
			Scale: scale,
			PosX: posX,
			PosY: posY,
			PosZ: posZ,
			RotX: rotX,
			RotY: rotY,
			RotZ: rotZ,
			CellFid: cellFid,
		}

		refrRows[formId] = refrRowJson
	}




	for _, stat := range stats {

		//refrDATA, _ := refr.GetFieldDataForType("DATA").(esm.PosRot)
		statOBND, _ := stat.GetFieldDataForType("OBND").(esm.OBND)
		//refrXSCL, _ := refr.GetFieldDataForType("XSCL").(float32)
		statMODL := esm.AsString(stat.GetFieldDataForType("MODL"))
		statMNAM, _ := stat.GetFieldDataForType("MNAM").(esm.StructLod4)

		formId := stat.FormId()
		model := statMODL
		boundsX1 := statOBND.X1
		boundsY1 := statOBND.Y1
		boundsZ1 := statOBND.Z1
		boundsX2 := statOBND.X2
		boundsY2 := statOBND.Y2
		boundsZ2 := statOBND.Z2
		lod1 := esm.AsString(statMNAM.LOD_1)
		lod2 := esm.AsString(statMNAM.LOD_2)
		lod3 := esm.AsString(statMNAM.LOD_3)
		lod4 := esm.AsString(statMNAM.LOD_4)

		statRowJson := StatStruct{
			FormId: formId,
			Model: model,
			BoundsX1: boundsX1,
			BoundsY1: boundsY1,
			BoundsZ1: boundsZ1,
			BoundsX2: boundsX2,
			BoundsY2: boundsY2,
			BoundsZ2: boundsZ2,
			Lod1: lod1,
			Lod2: lod2,
			Lod3: lod3,
			Lod4: lod4,
		}

		statRows[formId] = statRowJson
	}



	for _, cell := range cells {

		//fmt.Println(cell.Dump())

		//refrDATA, _ := refr.GetFieldDataForType("DATA").(esm.PosRot)
		//cellCELL, _ := cell.GetFieldDataForType("CELL").(esm.OBND)
		////refrXSCL, _ := refr.GetFieldDataForType("XSCL").(float32)
		//cellMODL := esm.AsString(cell.GetFieldDataForType("MODL"))
		cellXCLC, _ := cell.GetFieldDataForType("XCLC").(esm.GridXY)
		//cellFULL := esm.AsString(cell.GetFieldDataForType("FULL"))
		cellDATA, _ := cell.GetFieldDataForType("DATA").(uint16)
		cellEDID := esm.AsString(cell.GetFieldDataForType("EDID"))

		var wrldFid uint32

		wrldRecord := cell.NearestParentRecord()
		if wrldRecord != nil {
			wrldFid = wrldRecord.FormId()
		}

		formId := cell.FormId()
		editorId := cellEDID
		gridX := cellXCLC.X
		gridY := cellXCLC.Y
		//nameLid := cellFULL
		isInterior := (cellDATA & 0x1) != 0
		//
		cellRowJson := CellStruct{
			FormId: formId,
			EditorId: editorId,
			//NameLid: nameLid,
			IsInterior: isInterior,
			GridX: gridX,
			GridY: gridY,
			WrldFid: wrldFid,
		}
		
		cellRows[formId] = cellRowJson
	}
	for _, wrld := range wrlds {

		formId := wrld.FormId()
		//nameLid := esm.AsString(wrld.GetFieldDataForType("FULL"))

		wrldRowJson := WrldStruct{
			FormId: formId,
			//NameLid: nameLid,
		}

		wrldRows[formId] = wrldRowJson
	}

	for _, lctn := range lctns {
		formId := lctn.FormId()
		markerRefrFid := esm.AsUint32(lctn.GetFieldDataForType("MNAM"))

		lctnRowJson := LctnStruct{
			FormId: formId,
			MarkerRefrFid: markerRefrFid,
		}

		lctnRows[formId] = lctnRowJson
	}





	// REFR
	refrFullOut, _ := json.Marshal(refrRows)
	fileRefr, err1 := os.Create("refr.json")
	if err1 != nil {
		panic(err1)
	}
	defer fileRefr.Close()

	_, err = fileRefr.Write(refrFullOut)
	if err != nil {
		panic(err)
	}
	fileRefr.Sync()



	// STAT
	statFullOut, _ := json.Marshal(statRows)
	fileStat, err2 := os.Create("stat.json")
	if err2 != nil {
		panic(err2)
	}
	defer fileStat.Close()

	_, err = fileStat.Write(statFullOut)
	if err != nil {
		panic(err)
	}
	fileStat.Sync()

	// CELL
	cellFullOut, _ := json.Marshal(cellRows)
	fileCell, err3 := os.Create("cell.json")
	if err3 != nil {
		panic(err3)
	}
	defer fileCell.Close()

	_, err = fileCell.Write(cellFullOut)
	if err != nil {
		panic(err)
	}
	fileCell.Sync()

	// WRLD
	wrldFullOut, _ := json.Marshal(wrldRows)
	fileWrld, err4 := os.Create("wrld.json")
	if err4 != nil {
		panic(err4)
	}
	defer fileWrld.Close()

	_, err = fileWrld.Write(wrldFullOut)
	if err != nil {
		panic(err)
	}
	fileWrld.Sync()

	// LCTN
	lctnFullOut, _ := json.Marshal(lctnRows)
	fileLctn, err5 := os.Create("lctn.json")
	if err5 != nil {
		panic(err5)
	}
	defer fileLctn.Close()

	_, err = fileLctn.Write(lctnFullOut)
	if err != nil {
		panic(err)
	}
	fileLctn.Sync()


	//fmt.Println(string(jsonFullOut))
	fmt.Printf("REFR: %d records\n", len(refrs))
	fmt.Printf("STAT: %d records\n", len(stats))
	fmt.Printf("CELL: %d records\n", len(cells))
	fmt.Printf("WRLD: %d records\n", len(wrlds))
	fmt.Printf("LCTN: %d records\n", len(lctns))

}

