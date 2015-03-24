package GeneratorImageColorPalette

import (
	"fmt"
	"math"
	"sort"
	"github.com/disintegration/imaging"
)

type Color struct{
	K string
	v int
}

type ByVal []Color

func (a ByVal) Len() int { return len(a) }
func (a ByVal) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByVal) Less(i, j int) bool {
	if a[i].v == a[j].v {
		return a[i].K < a[j].K
	}
	return a[i].v > a[j].v
}


func GetImageColor(url string, params ...int) ([]Color) {

	var image_granularity int
	var error uint32 = 2
	var colors = make(map[string]int)

	if len(params) >= 2 {
		image_granularity = params[1]
	} else {
		image_granularity = 5
	}

	img, err := imaging.Open(url)

	if err != nil {
		panic(err)
	}

	bounds := img.Bounds()

	width := bounds.Max.X
	height := bounds.Max.Y

	// ---

	for x := 0; x < width; x += image_granularity {
		for y := 0; y < height; y += image_granularity {
			this_color := img.At(x, y)
			var r, g, b, _ = this_color.RGBA()

			var red    = uint8(Round(Round(float64(r/error),0,0) * float64(error), 0,0))
			var green  = uint8(Round(Round(float64(g/error),0,0) * float64(error), 0,0))
			var blue   = uint8(Round(Round(float64(b/error),0,0) * float64(error), 0,0))

			if red > 255 { red = 255 }
			if green > 255 { green = 255 }
			if blue > 255 { blue = 255 }

			red_s   := fmt.Sprintf("%02x", red)
			green_s := fmt.Sprintf("%02x", green)
			blue_s  := fmt.Sprintf("%02x", blue)

         	thisRGB := red_s + green_s + blue_s;

         	colors[thisRGB] = colors[thisRGB] + 1

 		}
	}

	// ---


	var cs []Color



	for key, value := range colors {
    	cs = append(cs, Color{K: key, v: value})
  	}

  	sort.Sort(ByVal(cs))

	return cs[:params[0]]

}

func Round(val float64, roundOn float64, places int ) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}