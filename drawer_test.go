package text2img

import (
	"image/jpeg"
	"os"
	"testing"
)

func TestDraw(t *testing.T) {
	path := "fonts/mplus-1c-bold.ttf"
	d, err := NewDrawer(Params{
		FontPath: path,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	img, err := d.Draw("text2img generates the image from a text")
	if err != nil {
		t.Fatal(err.Error())
	}
	file, err := os.Create("test.jpg")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		t.Fatal(err.Error())
	}
}

func TestDrawMultiline(t *testing.T) {

	drawMultiline(
		[]string{
			"text2img",
			"text2img generates the image from a text",
			"text2img generates the image from a text",
		},
		"multiline3.jpg",
		t)

	drawMultiline(
		[]string{
			"text2img",
			"text2img generates the image from a text",
		},
		"multiline2.jpg",
		t)

	drawMultiline(
		[]string{
			"text2img",
		},
		"multiline1.jpg",
		t)

}

func drawMultiline(texts []string, filename string, t *testing.T) {
	path := "fonts/mplus-1c-bold.ttf"
	d, err := NewDrawer(Params{
		FontPath: path,
		FontSize: 36,
		Width:    1280,
		Height:   600,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	img, err := d.DrawMultiline(texts)
	if err != nil {
		t.Fatal(err.Error())
	}
	file, err := os.Create(filename)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer file.Close()
	if err = jpeg.Encode(file, img, &jpeg.Options{Quality: 100}); err != nil {
		t.Fatal(err.Error())
	}
}
