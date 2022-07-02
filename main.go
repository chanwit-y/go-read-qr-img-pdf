package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"

	// _ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/liyue201/goqr"

	"github.com/karmdip-mi/go-fitz"
)

func convertPdfToImage() string {
	var fileName string
	files := []string{"out.pdf"}

	for _, file := range files {
		doc, err := fitz.New(file)
		if err != nil {
			panic(err)
		}
		folder := strings.TrimSuffix(path.Base(file), filepath.Ext(path.Base(file)))

		// Extract pages as images
		for n := 0; n < doc.NumPage(); n++ {
			img, err := doc.Image(n)
			if err != nil {
				panic(err)
			}
			err = os.MkdirAll("img/"+folder, 0755)
			if err != nil {
				panic(err)
			}

			fileName = fmt.Sprintf("image-%05d.jpg", n)
			f, err := os.Create(filepath.Join("img/"+folder+"/", fileName))
			if err != nil {
				panic(err)
			}

			err = jpeg.Encode(f, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
			if err != nil {
				panic(err)
			}

			f.Close()

		}
	}

	return fileName
}

func readImageQrCode(path string) {
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}
	for _, qrCode := range qrCodes {
		fmt.Printf("qrCode text: %s\n", qrCode.Payload)
	}
	return
}

func main() {
	// readImageQrCode("qr.png")
	// convertPdf2Image()
	readImageQrCode("./img/out/image-00002.jpg")

	// api.SplitFile("test.pdf", "./outDir", 1, nil)

	// inFiles := []string{"qr.pdf", "test.pdf"}
	// api.MergeCreateFile(inFiles, "out.pdf", nil)

}
