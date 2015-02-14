package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/SaidinWoT/fractal"
)

const (
	width  = 40
	height = 40

	scale  = 2
	minX   = -1 * scale
	minY   = -1 * scale
	deltaX = 1*scale - minX
	deltaY = 1*scale - minY

	pointsPerWalk = 10000000
)

func pointToPixel(p *fractal.Point) (int, int) {
	return int((p.X - minX) / deltaX * width),
		int((p.Y - minY) / deltaY * height)
}

func main() {
	rnd := rand.New(rand.NewSource(99))
	percent := flag.Float64("p", 0, "")
	dir := flag.String("dir", ".", "")

	flag.Parse()
	cs, vs := configure(*percent)

	hist := make([][]float64, height)
	for i := range hist {
		hist[i] = make([]float64, width)
	}

	max := float64(0)

	pt := &fractal.Point{
		rnd.Float64()*2 - 1,
		rnd.Float64()*2 - 1,
	}

	for i := 0; i < pointsPerWalk; i++ {
		r := rnd.Float64()
		pt = fractal.RandomFunc(r, cs, vs)(pt)

		if i < 20 {
			continue
		}
		x, y := pointToPixel(pt)
		if x < 0 || y < 0 || x >= width || y >= height {
			continue
		}
		hist[y][x] += 1
		if hist[y][x] > max {
			max = hist[y][x]
		}
	}

	path := filepath.Join(*dir, fmt.Sprintf("img%.2f.png", *percent))
	fmt.Println(path)
	w, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer w.Close()

	img := createImage(hist, max)
	png.Encode(w, img)
}

func createImage(hist [][]float64, max float64) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	c := color.RGBA{255, 255, 255, 255}
	for y, h := range hist {
		for x, v := range h {
			c.R = uint8(math.Log(v) * 255 / math.Log(max))
			c.G = c.R/2 + c.R/4
			c.B = 0
			fmt.Print(c)
			img.Set(x, y, c)
		}
	}
	return img
}
