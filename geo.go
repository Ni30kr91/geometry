
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var area int 
	var perimeter int 
	var operation int
	fmt.Println("enter shape- operation :[1] Rectangle, [2] Circle [3] Trinagle")
	var shape int
	fmt.Scanln(&shape)

	if shape == 1 {
		fmt.Println("Enter operation type- operation : [1] area, [2] perimeter")

		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Println("Enter Rectangle length")
			var length int
			fmt.Scanln(&length)
			fmt.Println("Enter Rectangle width")
			var width int
			fmt.Scanln(&width)
			area = Rectanglearea(length,width)
			fmt.Println("Area of Rectangle is:",area)
		} else if operation == 2 {
			fmt.Println("Enter Rectangle length")
			var length int 
			fmt.Scanln(&length)
			fmt.Println("Enter Rectangle width")
			var width int
			fmt.Scanln(&width)
			perimeter = Rectangleperimeter(length,width)
			fmt.Println("Perimeter of Rectangle is:",perimeter)
		} else {
			fmt.Println("Invalid Operation")
			return
		}
	} else if  shape == 2 {
		fmt.Println("Enter operation type- operation : [1] area, [2] perimeter")

		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Println("Enter Circe radius")
			var radius float32
			fmt.Scanln(&radius)

			var Circlearea = Circlearea(radius)
			fmt.Println("Area of Triangle is:",Circlearea)
		} else if operation == 2 {
			fmt.Println("Enter Circe radius")
			var radius float32
			fmt.Scanln(&radius)

			var Circleperimeter = CirclePerimeter(radius)
			fmt.Println("Perimeter of Circle is:",Circleperimeter)
		} else {
			fmt.Println("Invalid Operation")
			return
		}
	} else if shape == 3 {
		fmt.Println("Enter operstion type- operation : [1] area, [2] perimeter")
		
		fmt.Scanln(&operation)
		if operation == 1 {
			fmt.Println("Enter height of Triangle")
			var height int
			fmt.Scanln(&height)

			fmt.Println("Enter base of Triangle")
			var base int
			fmt.Scanln(&base)

			area = TriangleArea(height, base)
			fmt.Println("Area of Triangle is:",area)
		} else if operation == 2 {
			fmt.Println("Enter Triangle side A")
			var sideA int
			fmt.Scanln(&sideA)

			fmt.Println("Enter Triangle side B")
			var sideB int
			fmt.Scanln(&sideB)

			fmt.Println("Enter Triangle side C")
			var sideC int
			fmt.Scanln(&sideC)

			var Perimeter = TrianglePerimeter(sideA, sideB ,sideC)
			fmt.Println("Perimeter of Triangle is:", Perimeter)
		} else {
			fmt.Println("Invalid operation")
		}
	} else {
		fmt.Println("Invalid shape")
	}	
	
}
func Rectanglearea(length, width int) int {
	return length * width

}
func Rectangleperimeter(len, width int) int {
	return 2 * (len + width)
}         

func Circlearea(radius float32) float32 {
	var radiusSquare float32 = radius * radius
	var area = 3.14 * radiusSquare
	return area
}

func CirclePerimeter(radius float32) float32 {

	var perimeter = 2 * 3.14 * radius
	return perimeter
}

func TriangleArea(base, height int) int {

	var area = (base * height) / 2
	return area
}

func TrianglePerimeter(a, b, c int) int {

	return a + b + c

}