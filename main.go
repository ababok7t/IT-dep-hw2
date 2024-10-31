package main

import (
	"fmt"
	"hw2/calculator"
	"hw2/file"
	"hw2/shapes"
	"strconv"
	"strings"
)

func main() {

	var filePath string
	fmt.Println("Enter file path")
	_, scanErr := fmt.Scan(&filePath)
	if scanErr != nil {
		fmt.Println("scanning error: ", scanErr)
		return
	}

	fileText := file.ReadFile(filePath)

	arrayOfShapes := strings.Split(fileText, "\n")
	var arrayOfCircles []shapes.Circle
	var arrayOfRectangles []shapes.Rectangle

	for i := range arrayOfShapes {
		n := strings.Split(arrayOfShapes[i], " ")
		if n[0] == "circle" {
			par, err := strconv.ParseFloat(n[1], 64)
			if err != nil {
				fmt.Println("parse error: ", par)
				return
			}
			shape := shapes.Circle{Radius: par}
			arrayOfCircles = append(arrayOfCircles, shape)
		}
		if n[0] == "rectangle" {
			par1, err1 := strconv.ParseFloat(n[1], 64)
			if err1 != nil {
				fmt.Println("parse error: ", par1)
				return
			}
			par2, err2 := strconv.ParseFloat(n[2], 64)
			if err2 != nil {
				fmt.Println("parse error: ", par2)
				return
			}
			shape := shapes.Rectangle{Width: par1, Height: par2}
			arrayOfRectangles = append(arrayOfRectangles, shape)
		}
	}
	fmt.Println(calculator.TotalArea(arrayOfCircles, arrayOfRectangles))
}
