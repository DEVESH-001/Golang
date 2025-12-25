package main

func main() {
	// var gender = "male"

	// switch gender {
	// case "male":
	// 	println("male")
	// case "female":
	// 	println("female")
	// default:
	// 	println("unknown")
	// }

	//multiple conditions switch

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("It's the weekend BC!")
	// default:
	// 	println("It's a weekday")
	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("Integer h yo")
		case string:
			println("String h yo")
		case bool:
			println("Boolean h yo")
		default:
			println("Unknown type", t)
		}
	}
	whoAmI("https://devesh.work")
}
