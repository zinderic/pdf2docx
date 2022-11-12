package converter

import (
	"errors"
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gen2brain/go-fitz"
	"github.com/gingfrederik/docx"
)

var LocalDir string = ""

var (
	PdfFile = flag.String("file", ".", "pdf file")
	DirName = flag.String("output", "", "directory name where docx will be created")
)

func Pdf2Docx(doc *fitz.Document, filename string) {
	var allPages []string
	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			log.Fatalln(err)
		}
		allPages = append(allPages, text)
	}
	Text2Docx(allPages, filepath.Join(LocalDir, filename))
}

func Text2Docx(text interface{}, file string) {
	switch v := text.(type) {
	case string:
		f := docx.NewFile()
		// add new paragraph
		para := f.AddParagraph()
		// add text
		para.AddText(v).Size(16).Color("121212")
		f.Save(file)
	case []string:
		f := docx.NewFile()
		for _, t := range v {
			// add new paragraph
			para := f.AddParagraph()
			// add text
			para.AddText(t).Size(16).Color("121212")
		}
		docxFile := strings.Replace(file, "pdf", "docx", 1)
		f.Save(docxFile)
	default:
		log.Fatalln("can't parse text")
	}

}

func EasyMode() error {
	pdfs := getPdfFiles(".")
	if len(pdfs) == 0 {
		return errors.New("no pdf files found")
	}
	for _, v := range pdfs {
		doc, err := fitz.New(v)
		if err != nil {
			return err
		}
		defer doc.Close()
		Pdf2Docx(doc, v)
	}
	return nil
}

func getPdfFiles(dir string) []string {
	var files []string
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString("pdf", f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}
