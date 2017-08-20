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

		//"AACT",
		//"ACTI",
		//"ADDN",
		//"AECH",
		"ALCH",
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
		//"CELL",
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
		//"REFR",
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
		//"STAT",
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

	r, err := esm.OpenReader("C:/Program Files (x86)/Steam/steamapps/common/Fallout 4/Data/Fallout4.esm", allowedGroupTypes)

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()




}

