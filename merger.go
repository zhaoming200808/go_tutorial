package main

import (
	"encoding/xml"
	"fmt"
	//	"os"
)

type Merger struct {
	XMLName       xml.Name `xml:"Hoolai"`
	Mysqldest     string   `xml:"Merger>mysqldest"`
	Mysql1        Mysql    `xml:"Merger>mysql1"`
	Mysql2        Mysql    `xml:"Merger>mysql2"`
	Mysql3        Mysql    `xml:"Merger>mysql3"`
	Mergetype     int      `xml:"Merger>mergetype"`
	Maxcountrynum int      `xml:"Merger>maxcountrynum"`
}

type Mysql struct {
	Mysqlsrc string `xml:",chardata"`
	Country1 int    `xml:"country1,attr"`
	Country2 int    `xml:"country2,attr"`
	Country3 int    `xml:"country3,attr"`
}

func main() {
	conf := getMergerXml("Union")
	if conf == "" {
		fmt.Printf("error\n")
	}
	fmt.Printf("%s\n", conf)
}

func getMergerXml(t string) string {
	var mergetype int
	var maxcountrynum int
	if t == "Merger" {
		mergetype = 1
	} else if t == "Union" {
		mergetype = 2
	} else {
		return ""
	}

	maxcountrynum = 3
	m1 := Mysql{Mysqlsrc: "mysqlsrc1", Country1: 1, Country2: 1, Country3: 1}
	m2 := Mysql{Mysqlsrc: "mysqlsrc2", Country1: 2, Country2: 2, Country3: 2}
	m3 := Mysql{Mysqlsrc: "mysqlsrc3", Country1: 3, Country2: 3, Country3: 3}
	v := Merger{Mysqldest: "mysqldest", Mysql1: m1, Mysql2: m2, Mysql3: m3, Mergetype: mergetype, Maxcountrynum: maxcountrynum}
	output, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	//	os.Stdout.Write([]byte(xml.Header))
	//	os.Stdout.Write(output)
	//	fmt.Printf("\n\n%#v\n\n", v)
	return xml.Header + string(output)
}
