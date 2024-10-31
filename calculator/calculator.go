package calculator

import (
	"errors"
	"fmt"
	"hw2/shapes"
)

func TotalArea(circles []shapes.Circle, rectangles []shapes.Rectangle) float64 {
	var sum float64
	sum = 0
	if len(circles) == 0 && len(rectangles) == 0 {
		err := errors.New("no shapes")
		fmt.Println(err)
		return 0
	}
	for i := range circles {
		area, err := circles[i].GetArea()
		if err != nil {
			fmt.Println("error: ", err)
		}
		sum += area
	}
	for i := range rectangles {
		area, err := rectangles[i].GetArea()
		if err != nil {
			fmt.Println("error: ", err)
		}
		sum += area
	}
	return sum
}
