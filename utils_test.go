package petpet

import (
	"testing"

	"github.com/FloatTech/zbputils/img/writer"
)

func TestFace(t *testing.T) {
	dst, err := face("test/hinami.png", 500, 500)
	if err != nil {
		t.Fatal(err)
	}
	writer.SavePNG2Path("test/hinami_testFace.png", dst)
}

func TestTwoFace(t *testing.T) {
	list, err := twoface("test/hinami.png", 500, 500, "test/mimimi.png", 500, 500)
	if err != nil {
		t.Fatal(err)
	}
	writer.SavePNG2Path("test/hinami_testTwoFace.png", list[0])
	writer.SavePNG2Path("test/mimimi_testTwoFace.png", list[1])
}
