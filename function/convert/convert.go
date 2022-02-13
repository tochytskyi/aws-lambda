package convert

import (
	"io"
	"log"

	"github.com/sunshineplan/imgconv"
)

func Convert() {
	src, err := imgconv.Open("testdata/video-001.png")

	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, dst, imgconv.FormatOption{Format: imgconv.TIFF})

	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
}
