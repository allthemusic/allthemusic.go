package main

import (
	"bufio"
	"gopkg.in/alecthomas/kingpin.v1"
	"os"
)

var (
	debug = kingpin.Flag("debug", "enable debug mode").Default("false").Short('d').Bool()

	addMetadataCmd  = kingpin.Command("add-metadata", "Add a metadata JSON file to the index")
	addMetadataFile = addMetadataCmd.Arg("file", "path to JSON metadata file").Required().File()
)

func main() {
	switch kingpin.Parse() {
	case "add-metadata":
		AddMetadata(*addMetadataFile)
	}
}

func AddMetadata(file *os.File) {
	r := bufio.NewReader(file)
	r.WriteTo(os.Stdout)
}
