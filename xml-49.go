package main

import (
	"encoding/xml"
	"fmt"
)

type Plant struct {
	XMLName xml.Name `xml:"plant"`   //name of xml
	Id      int      `xml:"id,attr"` //attr states id is not a field but its a attribute
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"` //kinda array i.e. multiple values
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	//Prints <plant id="27">  <name>Coffee</name> <origin>Ethiopia</origin> <origin>Brazil</origin> </plant>
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))
	// Adds header <?xml version="1.0" encoding="UTF-8"?>
	fmt.Println(xml.Header + string(out))

	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	//Adding a new plant
	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	//Parent child tree form for existing xml
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
