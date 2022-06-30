package petpet

import (
	"fmt"
	"image"

	"github.com/FloatTech/zbputils/img"
	"github.com/FloatTech/zbputils/img/writer"
)

func Kiss(face1, face2 string, output string) error {
	list, err := twoface(face1, 40, 40, face2, 50, 50)
	if err != nil || len(list) != 2 {
		return err
	}
	userLocs := [][]int{{58, 90}, {62, 95}, {42, 100}, {50, 100}, {56, 100}, {18, 120}, {28, 110}, {54, 100}, {46, 100}, {60, 100}, {35, 115}, {20, 120}, {40, 96}}
	selfLocs := [][]int{{92, 64}, {135, 40}, {84, 105}, {80, 110}, {155, 82}, {60, 96}, {50, 80}, {98, 55}, {35, 65}, {38, 100}, {70, 80}, {84, 65}, {75, 65}}
	kiss := make([]*image.NRGBA, len(userLocs))
	for i := 0; i < len(userLocs); i++ {
		d, err := img.Load(fmt.Sprintf("resources/images/kiss/%v.png", i))
		if err != nil {
			return err
		}
		fac, err := img.LoadFirstFrame(fmt.Sprintf("resources/images/kiss/%v.png", i), d.Bounds().Dx(), d.Bounds().Dy())
		if err != nil {
			return err
		}
		kiss[i] = fac.InsertUp(list[0], list[0].Bounds().Dx(), list[0].Bounds().Dy(), userLocs[i][0], userLocs[i][1]).
			InsertUp(list[1], list[1].Bounds().Dx(), list[1].Bounds().Dy(), selfLocs[i][0], selfLocs[i][1]).Im
	}
	writer.SaveGIF2Path(output, img.MergeGif(6, kiss))
	return nil
}
