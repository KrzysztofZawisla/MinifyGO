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

func main() {
	flag.Parse()
	_, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}
	minifyJSON, err := ioutil.ReadFile("minify.json")
	if err != nil {
		log.Fatal(err.Error())
	}
	var minifyJSONStruct map[string]interface{}
	err = json.Unmarshal(minifyJSON, &minifyJSONStruct)
	if err != nil {
		log.Fatal(err.Error())
	}
	for key := range minifyJSONStruct {
		if key == "files" {
			for _, value := range minifyJSONStruct[key].([]interface{}) {
				cmd := exec.Command("minify", value.(string))
				fileName := strings.Split(value.(string), ".")
				outFile, err := os.Create(fileName[0] + ".min.js")
				if err != nil {
					log.Fatal(err.Error())
				}
				defer outFile.Close()
				cmd.Stdout = outFile
				err = cmd.Start()
				if err != nil {
					log.Fatal(err.Error())
				}
				cmd.Wait()
				if *outputToTheSameFile {
					if err != nil {
						log.Fatal(err.Error())
					}
					minifyText, err := ioutil.ReadFile(fileName[0] + ".min.js")
					if err != nil {
						log.Fatal(err.Error())
					}
					minifyTextAsString := strings.TrimSpace(string(minifyText))
					minifyText = []byte(minifyTextAsString)
					err = ioutil.WriteFile(value.(string), minifyText, 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
					outFile.Close()
					err = os.Remove(fileName[0] + ".min.js")
					if err != nil {
						log.Fatal(err.Error())
					}
				}
			}
		}
	}
}
