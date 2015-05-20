package main

import (
	"fmt"
	"gopkg.in/h2non/bimg.v0"
	"os"
)

func main() {
	buffer, err := bimg.Read("image.jpg")
	if err != nil {
		fmt.Println(os.Stderr, err)
	}

	newImage, err := bimg.NewImage(buffer).Resize(800, 600)
	if err != nil {
		fmt.Println(os.Stderr, err)
	}

	size, err := bimg.NewImage(newImage).Size()
	if size.Width == 400 && size.Height == 300 {
		fmt.Println("The image size is valid")
	}

	bimg.Write("new.jpg", newImage)
}
