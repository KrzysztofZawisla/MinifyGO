package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var outputToTheSameFile = flag.Bool("outputToTheSameFile", false, "When you set this flag to the true program will save minified code to the same file. When you set this flag to the false program will create a new file with extension .min.js next to the origin file.")

func checkErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	flag.Parse()
	minifyJSON, err := ioutil.ReadFile("minify.json")
	checkErr(err)
	var minifyJSONStruct map[string]interface{}
	err = json.Unmarshal(minifyJSON, &minifyJSONStruct)
	checkErr(err)
	for key := range minifyJSONStruct {
		if key == "files" {
			for _, value := range minifyJSONStruct[key].([]interface{}) {
				cmd := exec.Command("minify", value.(string))
				fileName := strings.Split(value.(string), ".")
				outFile, err := os.Create(fileName[0] + ".min.js")
				checkErr(err)
				defer outFile.Close()
				cmd.Stdout = outFile
				err = cmd.Start()
				checkErr(err)
				cmd.Wait()
				if *outputToTheSameFile {
					minifyText, err := ioutil.ReadFile(fileName[0] + ".min.js")
					checkErr(err)
					minifyTextAsString := strings.TrimSpace(string(minifyText))
					minifyText = []byte(minifyTextAsString)
					err = ioutil.WriteFile(value.(string), minifyText, 0644)
					checkErr(err)
					outFile.Close()
					err = os.Remove(fileName[0] + ".min.js")
					checkErr(err)
				}
			}
		}
	}
}
