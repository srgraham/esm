package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	_ "fmt"
	"github.com/srgraham/esm"
	_ "io/ioutil"
	"log"
	"os"
)

type LODType struct {
	LOD_1 string
	LOD_2 string
	LOD_3 string
	LOD_4 string
}

type RefrStruct struct {
	FormId     uint32 `json:"fid"`
	StatFormId uint32 `json:"statFid"`
	TNAM       int    `json:"TNAM"` // lctn marker icon id
	//Model string `json:"model"`
	Scale float32 `json:"scale"`
	PosX  float32 `json:"posX"`
	PosY  float32 `json:"posY"`
	PosZ  float32 `json:"posZ"`
	RotX  float32 `json:"rotX"`
	RotY  float32 `json:"rotY"`
	RotZ  float32 `json:"rotZ"`
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
	BoundsX1 int16  `json:"boundsX1"`
	BoundsY1 int16  `json:"boundsY1"`
	BoundsZ1 int16  `json:"boundsZ1"`
	BoundsX2 int16  `json:"boundsX2"`
	BoundsY2 int16  `json:"boundsY2"`
	BoundsZ2 int16  `json:"boundsZ2"`
	Lod1     string `json:"lod1"`
	Lod2     string `json:"lod2"`
	Lod3     string `json:"lod3"`
	Lod4     string `json:"lod4"`
}
type CellStruct struct {
	FormId uint32 `json:"fid"`
	GridX  int32  `json:"grid_x"`
	GridY  int32  `json:"grid_y"`
}
type WrldStruct struct {
	FormId uint32 `json:"fid"`
}
type IdAndNameStruct struct {
	FormId         uint32      `json:"fid"`
	NameLString    esm.LString `json:"nameLString"`
	NameLStringInt int         `json:"nameLSStringInt"`
}
type KywdStruct struct {
	FormId uint32 `json:"fid"`
	EDID   string `json:"EDID"`
}
type FormIdKywdAssocStruct struct {
	FormId uint32 `json:"fid"`
	KywdId uint32 `json:"kywdFid"`
}

func main() {
	fmt.Printf("yoooo")
	//r, err := esm.OpenReader("./SkjAlert_All_DLC.esp")
	//r, err := esm.OpenReader("./ShellRain.esp")

	allowedGroupTypes := []string{
		"REFR",
		"CELL",
		"STAT",
		"WRLD",
		"WEAP",
		"KYWD",
		"LCTN",
		"ALCH",
		"ARMO",
		"AMMO",
		"BOOK",
		"CONT",
		"INGR",
		"KEYM",
		"MISC",
		"SPEL",
	}

	r, root, err := esm.OpenReader("/Users/rmgraham/Downloads/Fallout4.esm", allowedGroupTypes)
	defer r.Close()

	if err != nil {
		log.Fatal(err)
	}

	esm.DumpUnimplementedFields()

	records := root.GetRecords()
	for _, row := range records {
		if row.FormId() == 121591 { //lctn
			fmt.Println("11111 %v", row.Dump())
		}
		if row.FormId() == 120551 { //refr
			fmt.Println("11112 %v", row.Dump())
		}
		if row.FormId() == 1542995 { //misc
			fmt.Println("11113 %v", row.Dump())
		}
	}

	//buildJsonFuncs["STAT"](root, "STAT")
	//buildJsonFuncs["KYWD"](root, "KYWD")
	//buildJsonFuncs["REFR"](root, "REFR")
	//buildJsonFuncs["CELL"](root, "CELL")
	//buildJsonFuncs["idAndName"](root, "WEAP")
	//buildJsonFuncs["idAndName"](root, "ALCH")
	//buildJsonFuncs["idAndName"](root, "ARMO")
	//buildJsonFuncs["idAndName"](root, "AMMO")
	//buildJsonFuncs["idAndName"](root, "BOOK")
	//buildJsonFuncs["idAndName"](root, "CONT")
	//buildJsonFuncs["idAndName"](root, "INGR")
	//buildJsonFuncs["idAndName"](root, "KEYM")
	buildJsonFuncs["idAndName"](root, "MISC")
	//buildJsonFuncs["idAndName"](root, "SPEL")
	//buildJsonFuncs["itemKywdAssoc"](root, "MISC")
}

func lstringToInt(lstr *esm.LString) uint32 {
	if lstr == nil {
		return 0
	}
	bs := []byte(*lstr)

	for charsLeft := 4 - len(bs); charsLeft > 0; charsLeft -= 1 {
		bs = append(bs, byte(0))
	}
	return binary.LittleEndian.Uint32(bs)
}

func saveJsonStrToFile(filename string, contents []byte) {
	filename = "out/" + filename
	fileStat, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fileStat.Close()

	_, err4 := fileStat.Write(contents)
	if err4 != nil {
		panic(err4)
	}
	fileStat.Sync()
}

var buildJsonFuncs = map[string]func(root *esm.Root, name string){
	// cell
	"CELL": func(root *esm.Root, name string) {
		cells := root.GetRecordsOfType(name)
		cellRows := make(map[uint32]CellStruct)
		for _, cell := range cells {
			cellXCLC, _ := cell.GetFieldDataForType("XCLC").(esm.GridXY)

			formId := cell.FormId()
			gridX := cellXCLC.X
			gridY := cellXCLC.Y

			cellRowJson := CellStruct{
				FormId: formId,
				//IsInterior: isInterior,
				GridX: gridX,
				GridY: gridY,
			}

			cellRows[formId] = cellRowJson
		}
		statFullOut, _ := json.Marshal(cellRows)
		saveJsonStrToFile(name+".json", statFullOut)

		fmt.Printf("%s: %d records\n", name, len(cells))
	},

	// weap
	"idAndName": func(root *esm.Root, name string) {
		items := root.GetRecordsOfType(name)
		rows := make(map[uint32]IdAndNameStruct)
		for _, item := range items {
			formId := item.FormId()
			nameLString, _ := item.GetFieldDataForType("FULL").(esm.LString)
			nameLStringInt := lstringToInt(&nameLString)
			fmt.Println(item.Dump())
			if formId == 105294 {
				fmt.Println("%v %v %v %v %v %v", formId, nameLString, nameLStringInt, []byte(nameLString), item, item.Dump())
				//os.Exit(0)
			}

			rowJson := IdAndNameStruct{
				FormId:         formId,
				NameLString:    nameLString,
				NameLStringInt: int(nameLStringInt),
			}

			rows[formId] = rowJson
		}
		str, _ := json.Marshal(rows)
		saveJsonStrToFile(name+".json", str)

		fmt.Printf("%s: %d records\n", name, len(rows))
	},
	// kywd
	"KYWD": func(root *esm.Root, name string) {
		items := root.GetRecordsOfType(name)
		rows := make(map[uint32]KywdStruct)
		for _, item := range items {
			formId := item.FormId()
			EDID := esm.AsString(item.GetFieldDataForType("EDID"))

			rowJson := KywdStruct{
				FormId: formId,
				EDID:   EDID,
			}

			rows[formId] = rowJson
		}
		str, _ := json.Marshal(rows)
		saveJsonStrToFile(name+".json", str)

		fmt.Printf("%s: %d records\n", name, len(rows))
	},
	"itemKywdAssoc": func(root *esm.Root, name string) {
		items := root.GetRecordsOfType(name)
		rows := make([]FormIdKywdAssocStruct, 0)
		for _, item := range items {
			fmt.Println("%v", item.Dump())
			formId := item.FormId()
			kywdIds := esm.AsUint32Arr(item.GetFieldDataForType("KWDA"))

			for _, kywdId := range kywdIds {
				rowJson := FormIdKywdAssocStruct{
					FormId: formId,
					KywdId: kywdId,
				}
				rows = append(rows, rowJson)
			}
		}
		str, _ := json.Marshal(rows)
		saveJsonStrToFile(name+"_kywd_assoc.json", str)
		fmt.Printf("%s: %d records\n", name+"_kywd_assoc", len(rows))
	},

	// refr
	"REFR": func(root *esm.Root, name string) {
		items := root.GetRecordsOfType(name)
		rows := make(map[uint32]RefrStruct)

		for _, item := range items {
			statFormId := esm.AsUint32(item.GetOneFieldForType("NAME").Data())
			if statFormId == 0 {
				fmt.Println("Skip failed .Data() %s", statFormId)
				continue
			}
			refrDATA, _ := item.GetFieldDataForType("DATA").(esm.PosRot)
			refrXSCL, _ := item.GetFieldDataForType("XSCL").(float32)
			refrTNAM, _ := item.GetFieldDataForType("TNAM").(uint16)

			formId := item.FormId()
			scale := refrXSCL
			posX := refrDATA.Position.X
			posY := refrDATA.Position.Y
			posZ := refrDATA.Position.Z
			rotX := refrDATA.Rotation.X
			rotY := refrDATA.Rotation.Y
			rotZ := refrDATA.Rotation.Z
			tnam := int(refrTNAM)

			rowJson := RefrStruct{
				FormId:     formId,
				StatFormId: statFormId,
				TNAM:       tnam,
				Scale:      scale,
				PosX:       posX,
				PosY:       posY,
				PosZ:       posZ,
				RotX:       rotX,
				RotY:       rotY,
				RotZ:       rotZ,
			}

			rows[formId] = rowJson
		}
		str, _ := json.Marshal(rows)
		saveJsonStrToFile(name+".json", str)

		fmt.Printf("%s: %d records\n", name, len(rows))
	},

	// stat
	"STAT": func(root *esm.Root, name string) {
		stats := root.GetRecordsOfType(name)
		statRows := make(map[uint32]StatStruct)

		for _, stat := range stats {
			statOBND, _ := stat.GetFieldDataForType("OBND").(esm.OBND)
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
				FormId:   formId,
				Model:    model,
				BoundsX1: boundsX1,
				BoundsY1: boundsY1,
				BoundsZ1: boundsZ1,
				BoundsX2: boundsX2,
				BoundsY2: boundsY2,
				BoundsZ2: boundsZ2,
				Lod1:     lod1,
				Lod2:     lod2,
				Lod3:     lod3,
				Lod4:     lod4,
			}

			statRows[formId] = statRowJson
		}
		statFullOut, _ := json.Marshal(statRows)
		saveJsonStrToFile(name+".json", statFullOut)

		fmt.Printf("%s: %d records\n", name, len(stats))
	},
}
