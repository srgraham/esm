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
	r, err := esm.OpenReader("./ShellRain.esp")

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()




}

