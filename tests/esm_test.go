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
		//"LCTN",
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
		//"WRLD",
		//"WTHR",
		//"ZOOM",
	}

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

	records := root.GetRecordsOfType("REFR")
	fmt.Printf("%d records\n", len(records))

	//refrRows := make([]*RefrStruct, 0)
	//statRows := make([]*StatStruct, 0)

	refrRows := make(map[uint32]RefrStruct)
	statRows := make(map[uint32]StatStruct)

	for _, refr := range records {

		statFormId := esm.AsUint32(refr.GetOneFieldForType("NAME").Data())

		if statFormId == 0 {
			fmt.Println("Skip failed .Data() %s", statFormId)
			continue
		}

		stat := root.GetRecordByFormIdUint32(statFormId)

		if stat == nil {
			fmt.Printf("nil refr for %d\n", statFormId)
			continue
		}

		if stat.Type() != "STAT" {
			fmt.Printf("skipping non-STAT %s\n", stat.Type())
			continue
		}

		fmt.Println(refr.Dump())
		fmt.Println(stat.Dump())



		refrDATA, _ := refr.GetFieldDataForType("DATA").(esm.PosRot)
		statOBND, _ := stat.GetFieldDataForType("OBND").(esm.OBND)
		refrXSCL, _ := refr.GetFieldDataForType("XSCL").(float32)
		statMODL := esm.AsString(stat.GetFieldDataForType("MODL"))
		statMNAM, _ := stat.GetFieldDataForType("MNAM").(esm.StructLod4)



		fmt.Println(statMNAM)
		formId := refr.FormId()
		model := statMODL
		scale := refrXSCL
		posX := refrDATA.Position.X
		posY := refrDATA.Position.Y
		posZ := refrDATA.Position.Z
		rotX := refrDATA.Rotation.X
		rotY := refrDATA.Rotation.Y
		rotZ := refrDATA.Rotation.Z
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

		refrRowJson := RefrStruct{
			FormId: formId,
			StatFormId: statFormId,
			//Model: model,
			Scale: scale,
			PosX: posX,
			PosY: posY,
			PosZ: posZ,
			RotX: rotX,
			RotY: rotY,
			RotZ: rotZ,
			//BoundsX1: boundsX1,
			//BoundsY1: boundsY1,
			//BoundsZ1: boundsZ1,
			//BoundsX2: boundsX2,
			//BoundsY2: boundsY2,
			//BoundsZ2: boundsZ2,
		}

		statRowJson := StatStruct{
			FormId: statFormId,
			//StatFormId: statFormId,
			Model: model,
			//Scale: scale,
			//PosX: posX,
			//PosY: posY,
			//PosZ: posZ,
			//RotX: rotX,
			//RotY: rotY,
			//RotZ: rotZ,
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

		//fmt.Println(model, scale, posX, posY, posZ, rotX, rotY, rotZ, boundsX1, boundsY1, boundsZ1, boundsX2, boundsY2, boundsZ2)
		//jsonOut, _ := json.Marshal(refrRowJson)
		//jsonOut, _ := json.Marshal(statRowJson)
		//fmt.Println("asdf", string(jsonOut))

		//refrRows = append(refrRows, &refrRowJson)
		//statRows = append(statRows, &statRowJson)

		refrRows[formId] = refrRowJson
		statRows[statFormId] = statRowJson
	}

	refrFullOut, _ := json.Marshal(refrRows)
	statFullOut, _ := json.Marshal(statRows)

	fileRefr, err := os.Create("refr.json")
	if err != nil {
		panic(err)
	}
	defer fileRefr.Close()

	_, err2 := fileRefr.Write(refrFullOut)
	if err2 != nil {
		panic(err2)
	}
	fileRefr.Sync()




	fileStat, err3 := os.Create("stat.json")
	if err3 != nil {
		panic(err3)
	}
	defer fileStat.Close()

	_, err4 := fileStat.Write(statFullOut)
	if err4 != nil {
		panic(err4)
	}
	fileStat.Sync()


	//fmt.Println(string(jsonFullOut))
	fmt.Printf("%d records\n", len(records))

}

