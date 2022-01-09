package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

//go:embed fonts/* graphics/*
var EmbedData embed.FS

func NewImageFromEmbedFile(path string) (*ebiten.Image, image.Image, error) {
	file, err := EmbedData.Open(path)
	if err != nil {
		return nil, nil, err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, nil, err
	}
	img2 := ebiten.NewImageFromImage(img)
	return img2, img, err
}
