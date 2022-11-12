package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/gen2brain/go-fitz"
)

var localDir string = "temp"

func main() {
	doc, err := fitz.New("test.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	// TODO create in the current directory instead of temp dir
	if err := os.Mkdir(localDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	// Extract pages as images
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(localDir, fmt.Sprintf("test%03d.jpg", n)))
		if err != nil {
			panic(err)
		}

		err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(localDir, fmt.Sprintf("test%03d.txt", n)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(text)
		if err != nil {
			panic(err)
		}

		f.Close()
	}

	// Extract pages as html
	for n := 0; n < doc.NumPage(); n++ {
		html, err := doc.HTML(n, true)
		if err != nil {
			panic(err)
		}

		f, err := os.Create(filepath.Join(localDir, fmt.Sprintf("test%03d.html", n)))
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString(html) // TODO output to word file
		if err != nil {
			panic(err)
		}

		f.Close()
	}
}
