package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
)

type outer struct {
	Id          string  `xml:"id"`
	Transaction string  `xml:"transaction"`
	InnerList   []inner `xml:"inner"`
}

type inner struct {
	Data string `xml:",chardata"`
}

func main() {

	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var output *os.File
	output, err = os.Create("output")
	if err != nil {
		panic(err)
	}
	defer output.Close()
	output.WriteString("id, transaction, value \n")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parseLine(scanner.Bytes(), output)
	}

}

func parseLine(line []byte, output *os.File) {
	var outer outer
	xml.Unmarshal(line, &outer)
	for _, inn := range outer.InnerList {
		csvLine := fmt.Sprintf("%v, %v, %v \n", outer.Transaction, outer.Id, inn.Data)
		output.WriteString(csvLine)
	}
}
