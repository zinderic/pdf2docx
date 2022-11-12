package main

import (
	"flag"
	"log"
	"os"
	"pdf2docx/pkg/converter"

	"github.com/gen2brain/go-fitz"
)

func main() {
	flag.Parse()
	if *converter.PdfFile == "." && *converter.DirName == "" {
		err := converter.EasyMode()
		if err != nil {
			log.Fatalln(err)
		}
		os.Exit(0)
	}
	converter.LocalDir = *converter.DirName
	if *converter.PdfFile == "" {
		log.Fatalln("missing pdf file, use -file <pdf file>")
	}
	if *converter.DirName == "" {
		log.Fatalln("missing output directory, use -output <directory name>")
	}
	doc, err := fitz.New(*converter.PdfFile)
	if err != nil {
		log.Fatal(err)
	}

	defer doc.Close()

	if err := os.Mkdir(converter.LocalDir, os.ModePerm); err != nil {
		if os.IsExist(err) {
			log.Println(converter.LocalDir + " already exist, cleanup before re-run")
			os.Exit(0)
		}
		log.Fatal(err)
	}

	converter.Pdf2Docx(doc, *converter.PdfFile)

}
