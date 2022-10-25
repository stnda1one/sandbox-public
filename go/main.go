package main

import (
	_ "embed"
	"bytes"
	"fmt"
	"text/template"
)

//go:embed embed/sample.template
var sampleTemplate string

func main() {  
  roles := make([]string, 0, 5)
  roles = append(roles, "test")
  buf := bytes.Buffer{}
  template.Must(template.New("").Parse(sampleTemplate)).Execute(&buf, roles)
  fmt.Println("------------------- start -------------------")
  fmt.Println(buf.String())
  fmt.Println("-------------------- end --------------------")

}
