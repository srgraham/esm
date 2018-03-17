package tests

import (
	_ "fmt"
	_ "io/ioutil"
	"testing"
	"log"
	"fmt"
	esm ".."
	_ "encoding/json"
	_ "os"
)


func TestReadStringsXxx(t *testing.T) {
	fmt.Printf("strings yo")
	//r, root, err := esm.OpenReader("C:/Program Files (x86)/Steam/steamapps/common/Fallout 4/Data/Fallout4.esm", allowedGroupTypes)
	r, stringsRoot, err := esm.ReadStrings("./Fallout4_en.STRINGS")
	defer r.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stringsRoot.GetStringForId(77289 ))

}

