package main

import (
	"encoding/json"
	"fmt"
	"os"
	"ryangurnick.com/mapy/app/models"
)

var app = models.App{
	Name: "Mapy Import Tool",
	Major: 0,
	Minor: 1,
	Build: 1,
	Authors: []string {
		"Ryan Gurnick",
	},
}

func main() {
	fmt.Println(app.GetName());
	fmt.Println(app.GetVersion());
	
	// ensure that we have at least 1 parameter (aka the program itself)
	if (len(os.Args) <= 1) {
		fmt.Println("No input file provided as the first parameter")
		return
	}
	
	importFile := os.Args[1]
	
	fmt.Println("Input file " + importFile + " specified for importing")
	file, err := os.Stat(importFile)
	if err != nil {
		fmt.Println("An input related file error occurred, the file might not exist or some other issue may have occurred");
		return
	}
	
	fmt.Println("Input file " + importFile + " exists and import process continuing")
	fmt.Println(fmt.Sprintf("Input file size: %d", file.Size()))
	
	data, err := os.ReadFile(importFile)
	if err != nil {
		fmt.Println("Cannot open file: " + importFile)
	}
	
	fileModel := models.File{File: importFile, Size: file.Size(), Data: data}
	if err := json.Unmarshal(fileModel.Data, &fileModel.GeoJSON); err != nil {
		panic(err)
	}
	
	fmt.Println("Loaded GeoJSON with name: " + fileModel.GeoJSON.Properties.Name)
		
	
	return
}
