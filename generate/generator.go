package generate

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

type Output struct {
	DataString string
	Key string
}

func byteToString(by []byte) string {

	var dataSlice []string

	for _, b := range by {
		bString := fmt.Sprintf("%v", b)
		dataSlice = append(dataSlice, bString)
	}

	dataString := strings.Join(dataSlice, ", ")

	return dataString

}

func writeToTemplate(templatename string, outfile *os.File, outStruct *Output) {

	// Open output-[linux,windows].tmpl
	t, err := template.ParseFiles(templatename)
	if err != nil {
		log.Print(err)
		return
	}

	err = t.Execute(outfile, outStruct)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
	outfile.Close()

}


func Generate(input, output *string, win *bool) {

	opersys := "linux"

	// Logic to determine if building for Windows or Linux
	if *win {	opersys = "windows"	}

	// Create filename.go which will contain binary
	outputFileName := *output + ".go"

	outfile, err := os.Create(outputFileName)
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()

	// Process executable
	infile, err := ioutil.ReadFile(*input)
	if err != nil {
		panic(err.Error())
	}

	// Compress bytes
	compressedString := Compress(infile)

	// Encrypt byte
	encryptedBytes, keyBytes := Encrypt(compressedString)

	// Convert bytes to string for template
	dataString := byteToString(encryptedBytes)
	keyString := byteToString(keyBytes)

	// Create a structure for template processing
	outStruct := Output{DataString: dataString, Key: keyString }


	// Write to file
	if opersys == "windows" {
		templatename := "output-windows.tmpl"
		writeToTemplate(templatename, outfile, &outStruct)
	} else {
		templatename := "output-linux.tmpl"
		writeToTemplate(templatename, outfile, &outStruct)
	}

}