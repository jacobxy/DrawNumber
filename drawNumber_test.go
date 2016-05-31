package DrawNumber

import (
	"testing"
)

func Test_DrawNumber(t *testing.T) {
	img := NewImage("base.jpg", "new2.jpg")
	x, y := img.ChangePos(164, 124)
	img.DrawNumber("17.11", x, y, 6, 255, 255, 255)

	img2 := NewImageBySlice(img.Content, "new22.jpg")
	img2.DrawNumber("100.01", 500, 500, 10, 255, 255, 0)
	defer img.Close()
	defer img2.Close()
	t.Log("success")
}
