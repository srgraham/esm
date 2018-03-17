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
	"encoding/json"
	"os"
)


func TestReadStringsXxx(t *testing.T) {
	fmt.Printf("strings yo")
	//r, root, err := esm.OpenReader("C:/Program Files (x86)/Steam/steamapps/common/Fallout 4/Data/Fallout4.esm", allowedGroupTypes)
	stringsRoot, err := esm.ReadStrings("./Fallout4_en.STRINGS")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stringsRoot.GetStringForId(77289 ))


	allStringsJson, _ := json.Marshal(stringsRoot.GetAllStrings())
	fileStrings, err1 := os.Create("strings.json")
	if err1 != nil {
		panic(err1)
	}
	defer fileStrings.Close()

	_, err = fileStrings.Write(allStringsJson)
	if err != nil {
		panic(err)
	}
	fileStrings.Sync()


}

