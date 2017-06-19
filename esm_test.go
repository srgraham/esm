package esm

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestXxx(t *testing.T) {
	type Item struct {
		Foo string
	}

	buf, err := ioutil.ReadFile("./armorsmith.esp")

	//b, err := esm.Marshal(&Item{Foo: "bar"})
	//if err != nil {
	//	panic(err)
	//}

	var item Item
	err = Unmarshal(buf, &item)
	if err != nil {
		panic(err)
	}
	fmt.Println(item.Foo)
	// Output: bar
}

