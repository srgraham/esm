package tests

import (
	_ "fmt"
	_ "io/ioutil"
	"testing"
	"log"
	"fmt"
	esm ".."
)

func TestXxx(t *testing.T) {
	fmt.Printf("yoooo")
	//r, err := esm.OpenReader("./SkjAlert_All_DLC.esp")
	//r, err := esm.OpenReader("./ShellRain.esp")

	allowedGroupTypes := []string{
		//"DOOR",
		//"ECZN",
		//"EQUP",
		//"FLST",
		//"FURN",
		//"GLOB",
		//"GMST",
		//"HAZD",
		//"IDLM",
		//"INGR",
		//"KEYM",
		//"KYWD",
		//"LAYZ",
		//"LCRT",
		//"LCTN",
		//"LSCR",
		//"LVLI",
		//"LVLN",
		//"MESG",
		//"MGEF",
		//"MISC",
		//"MOVT",
		//"MSTT",
		//"NAVI",
		//"NOCM",
		//"NOTE",
		//"NPC_",
		//"OTFT",
		//"PACK",
		//"PERK",
		//"PGRE",
		//"PHZD",
		//"PROJ",
		//"QUST",
		//"RACE",
		//"REFR",
		//"REGN",
		//"RELA",
		//"RFGP",
		//"SCCO",
		//"SCOL",
		//"SMBN",
		//"SMEN",
		//"SMQN",
		//"SPEL",
		//"STAT",
		//"TERM",
		//"TREE",
		//"TRNS",
		//"TXST",
		//"WATR",
		//"WEAP",
		"WRLD",
		//"WTHR",
	}

	r, err := esm.OpenReader("C:/Program Files (x86)/Steam/steamapps/common/Fallout 4/Data/Fallout4.esm", allowedGroupTypes)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()




}

