package esm

import (
	_ "fmt"
	_ "io/ioutil"
	"testing"
	"log"
	"fmt"
)

func TestXxx(t *testing.T) {
	fmt.Printf("yoooo")
	r, err := OpenReader("./ShellRain.esp")

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()




}

