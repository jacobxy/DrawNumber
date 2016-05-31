package DrawNumber

import (
	"fmt"
	"github.com/Comdex/imgo"
	"log"
	"os"
)

var Numbers map[string]*[][][]uint8

type Image struct {
	Content *[][][]uint8
	File    string
	FileOut string
}

func NewImage(file string, fileout string) *Image {

	img := imgo.MustRead(file)
	return &Image{
		Content: &img,
		File:    file,
		FileOut: fileout,
	}
}

func NewImageBySlice(img *[][][]uint8, fileout string) *Image {
	return &Image{
		Content: img,
		File:    "",
		FileOut: fileout,
	}
}

func (i *Image) Close() {
	imgo.SaveAsJPEG(i.FileOut, *(i.Content), 100)
}

func init() {

	Numbers = make(map[string]*[][][]uint8)

	path, err := os.Getwd()

	if err != nil {
		log.Println("file Path Error")
		return
	}

	for i := 0; i < 10; i++ {
		sti := fmt.Sprintf("%d", i)
		img := imgo.MustRead(path + string(os.PathSeparator) + sti + ".jpg")
		if _, ok := Numbers[sti]; !ok {
			Numbers[sti] = &img
		}
	}

	sti := "."
	img := imgo.MustRead(path + string(os.PathSeparator) + sti + ".jpg")
	if _, ok := Numbers[sti]; !ok {
		Numbers[sti] = &img
	}
}

func Print() {
	for k, v := range Numbers {
		fmt.Println(k, *v)
		break
	}
}

func (img *Image) ChangePos(x int, y int) (x1 int, y1 int) {
	height := len(*(img.Content))
	x, y = (x-64)*height/128, (y-64)*height/128
	return height - y, x
}

func (img *Image) DrawNumber(num string, x, y, rad int, r, g, b uint8) {
	for _, v := range num {
		n := string(v)
		vs, ok := Numbers[n]
		if !ok || len(*vs) == 0 {
			log.Println(" n is Error", n)
			return
		}
	}

	x1 := 0
	for _, v := range num {
		n := string(v)
		vs, _ := Numbers[n]
		img.DrawNumberOne(n, x, y+x1/rad, rad, r, g, b)
		x1 += len((*vs)[0])
	}
}

func (img *Image) DrawNumberOne(num string, x, y, rad int, r, g, b uint8) {
	sti := (num)
	ptr, ok := Numbers[sti]
	if !ok {
		log.Println("Number", num)
		return
	}
	height := len(*(img.Content))
	v := *ptr

	if height < len(v)/rad+x || height < len((v)[0])/rad+y {
		log.Println(" Error ")
		return
	}

	for i := 0; i*rad < len(v); i++ {
		for j := 0; j*rad < len(v[0]); j++ {
			flag := false
			for idx := 0; idx < 3; idx++ {
				if v[i*rad][j*rad][idx] > 150 {
					flag = true
					break
				}
			}
			if !flag {
				//for idx := 0; idx < 3; idx++ {
				//(*(img.Content))[x+i][y+j][idx] = v[i*rad][j*rad][idx]
				//}
				(*(img.Content))[x+i][y+j][0] = r
				(*(img.Content))[x+i][y+j][1] = g
				(*(img.Content))[x+i][y+j][2] = b
			}
		}
	}
}
