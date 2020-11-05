package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

const (
	imageWidth        = 25
	imageHeight       = 6
	imageSizeInPixels = imageWidth * imageHeight
)

func main() {
	img := readInImage()

	doPart1(img)
	doPart2(img)
}

func doPart1(img image) {
	layerWithMost0s := findLayerWithFewest0s(img.layers)

	n1Digits := bytes.Count(layerWithMost0s, []byte{'1'})
	n2Digits := bytes.Count(layerWithMost0s, []byte{'2'})

	log.Printf("num 1 digits * num 2 digits = %d * %d = %d", n1Digits, n2Digits, n1Digits*n2Digits)
}

func findLayerWithFewest0s(layers [][]byte) []byte {
	var fewestZerosLayer []byte
	fewestZerosInLayer := -1

	for _, layer := range layers {
		nZeros := bytes.Count(layer, []byte{'0'})
		if fewestZerosInLayer == -1 || nZeros < fewestZerosInLayer {
			fewestZerosInLayer = nZeros
			fewestZerosLayer = layer
		}
	}

	return fewestZerosLayer
}

type image struct {
	bytes  []byte
	layers [][]byte
}

func readInImage() image {
	imageBytes, _ := ioutil.ReadAll(os.Stdin)

	log.Printf("image consists of %d bytes, %d layers of length %d", len(imageBytes), len(imageBytes)/imageSizeInPixels, imageSizeInPixels)

	var img image
	img.bytes = imageBytes

	nLayers := len(imageBytes) / imageSizeInPixels
	img.layers = make([][]byte, nLayers)
	for i := 0; i < nLayers; i++ {
		low := i * imageSizeInPixels
		high := low + imageSizeInPixels
		img.layers[i] = imageBytes[low:high]
	}

	return img
}

func doPart2(img image) {
	var finalImage [imageSizeInPixels]byte

	for _, layer := range img.layers {
		for i, v := range layer {
			imgV := finalImage[i]

			if imgV == 0 && (v == '0' || v == '1') {
				finalImage[i] = v
			}
		}
	}

	bufStdout := bufio.NewWriter(os.Stdout)
	defer bufStdout.Flush()
	for i := 0; i < imageHeight; i++ {
		for j := 0; j < imageWidth; j++ {
			pos := i*imageWidth + j
			if finalImage[pos] == '1' {
				_, _ = bufStdout.WriteRune('O')
			} else {
				_, _ = bufStdout.WriteRune(' ')
			}
		}
		_, _ = bufStdout.WriteRune('\n')
	}
}
