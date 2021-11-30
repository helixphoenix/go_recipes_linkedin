package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"os"
)

var (
	Red   = color.RGBA{0xFF, 0, 0, 0xFF}
	Green = color.RGBA{0, 0xFF, 0, 0xFF}
	Blue  = color.RGBA{0, 0, 0xFF, 0xFF}
)
type Shape struct {
	X int
	Y int 
	Color color.Color
}

type Device interface {
	Set(int, int, color.Color)
}

type Drawer interface {
	Draw(d Device)
}

type Circle struct {
	Shape
	Radius int
}


func NewCircle(x,y,r int, c color.Color) *Circle {
	circ := Circle{
		Shape: Shape{x, y,c},
		Radius: r,
	}
	return &circ

}

func (circ *Circle) Draw(d Device){
	minX, minY :=circ.X-circ.Radius, circ.Y-circ.Radius
	maxX, maxY :=circ.X+circ.Radius, circ.Y+circ.Radius
	for x := minX; x <= maxX; x++ {
		for y :=minY; y<= maxY; y++ {
			dx,dy := x-circ.X, y-circ.Y
			if int(math.Sqrt(float64(dx*dx+dy*dy))) <= circ.Radius{
				d.Set(x,y,circ.Color)
			}
		}
	}
}

// func NewImageCanvas(width, height int) (*image.RGBA, error) {
// 	canvas := image.NewRGBA(image.Rect(0, 0, width, height))
// 	canvas.Set(width, height, image.Transparent)
// 	return canvas, nil
// }

type Rectangle struct {
	Shape
	Height int 
	Width int 
}


func NewRectangle(x, y, h, w int, c color.Color) *Rectangle{
    rect := Rectangle{
		Shape: Shape{x,y,c},
		Height: h,
		Width: w,
	}
	return &rect
}

func (rect *Rectangle) Draw(d Device){
	minX, minY :=rect.X-rect.Width/2, rect.Y-rect.Height/2
	maxX, maxY :=rect.X+rect.Width/2, rect.Y+rect.Height/2
	for x := minX; x <= maxX; x++ {
		for y :=minY; y<= maxY; y++ {
			d.Set(x,y,rect.Color)

			}
		}
	}

type ImageCanvas struct {
		width int 
		height int 
		shapes []Drawer 
}

func NewImageCanvas(width, height int) (*ImageCanvas, error) {
        if width <= 0 || height <= 0{
				return nil, fmt.Errorf("negative size: width=%d, height=%d", width, height)
		}
		canvas:= ImageCanvas{
			width: width, 
			height: height,
		}
		return &canvas, nil
}


func (ic *ImageCanvas) Add(d Drawer)  {
	ic.shapes = append(ic.shapes, d)

}


func (ic *ImageCanvas) Draw(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0,0, ic.width, ic.height))
	for _, shape := range ic.shapes {
			shape.Draw(img)
	}
	return png.Encode(w, img)
}


// func Add(d Drawer) {
// 	switch d.(type){
// 	case *Circle:
// 			fmt.Println("Adding circle")
// 			draw.Draw(ic, ic.Bounds(), Circle, image.Point{0, 0}, draw.Src)
// 	case *Rectangle:
// 			fmt.Println("Adding rectangle")
// 			draw.Draw(ic, ic.Bounds(), Rectangle, image.Point{0, 0}, draw.Src)
// }
// }



func main() {
	ic, err := NewImageCanvas(200, 200)
	if err != nil {
		log.Fatal(err)
	}
  
	ic.Add(NewCircle(100, 100, 80, Green))
	ic.Add(NewCircle(60, 60, 10, Blue))
	ic.Add(NewCircle(140, 60, 10, Blue))
	ic.Add(NewRectangle(100,130, 10, 80, Red))
	ic.Add(NewCircle(100,150, 20, Red))

	f, err := os.Create("face.png")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	if err := ic.Draw(f); err!= nil {
			log.Fatal(err)
	}
}


