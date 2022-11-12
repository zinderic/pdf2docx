package converter

import (
	"fmt"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
	"github.com/gingfrederik/docx"
)

var LocalDir string = ""

func Pdf2Docx(doc *fitz.Document) {
	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		Text2Docx(text, filepath.Join(LocalDir, fmt.Sprintf("page%03d.docx", n)))

	}
}

func Text2Docx(text string, file string) {
	f := docx.NewFile()
	// add new paragraph
	para := f.AddParagraph()
	// add text
	para.AddText(text).Size(16).Color("121212")
	f.Save(file)
}
