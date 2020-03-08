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


func Generate(input, output *string) {

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

	// Open output.tmpl
	t, err := template.ParseFiles("output.tmpl")
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