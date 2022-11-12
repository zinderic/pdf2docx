# pdf2docx

Simple pdf to docx converter. It will extract only the text from pdf and put it in docx for editing.

Usage:

```
pdf2docx -file <pdf file> -output <output directory>
```

## Build on Mac

You need `mupdf` and `mupdf-tools` installed. You also need `mingw` for Windows cross-compilation.

```
brew install mupdf
brew install mupdf-tools
brew install mingw-w64
```

Example build:

```
CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" CXX="x86_64-w64-mingw32-g++" GOOS=windows GOARCH=amd64 go build -tags "" -ldflags "-extldflags -static" -o pdf2docx ./cmd/main.go
```
