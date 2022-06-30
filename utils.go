package petpet

import (
	"image"

	"github.com/FloatTech/zbputils/img"
)

func face(face1 string, w int, h int) (*image.NRGBA, error) {
	frame, err := img.LoadFirstFrame(face1, w, h)
	if err != nil {
		return nil, err
	}
	return frame.Circle(0).Im, nil
}

func twoface(face1 string, w1, h1 int, face2 string, w2, h2 int) (list []*image.NRGBA, err error) {
	list = make([]*image.NRGBA, 0)
	image1, err := face(face1, w1, h1)
	if err != nil {
		return
	}
	image2, err := face(face2, w2, h2)
	if err != nil {
		return
	}
	list = append(list, image1, image2)
	return
}
