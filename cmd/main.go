package main

import (
	"flag"
	"log"
	"os"
	"pdf2docx/pkg/converter"

	"github.com/gen2brain/go-fitz"
)

func main() {
	pdfFile := flag.String("file", "", "pdf file")
	dirName := flag.String("output", "", "directory name where docx will be created")
	flag.Parse()
	converter.LocalDir = *dirName
	if *pdfFile == "" {
		log.Fatalln("missing pdf file, use -file <pdf file>")
	}
	if *dirName == "" {
		log.Fatalln("missing output directory, use -output <directory name>")
	}
	doc, err := fitz.New(*pdfFile)
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	if err := os.Mkdir(converter.LocalDir, os.ModePerm); err != nil {
		if os.IsExist(err) {
			log.Println(converter.LocalDir + " already exist, cleanup before re-run")
			os.Exit(0)
		}
		log.Fatal(err)
	}

	converter.Pdf2Docx(doc)

}
