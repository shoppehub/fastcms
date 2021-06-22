package list

import (
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestList(t *testing.T) {
	var l List

	strb, _ := ioutil.ReadFile("./test.yaml")
	err := yaml.Unmarshal(strb, &l)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(l.Items)

}
