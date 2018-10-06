package main

import "fmt"

func minInt(nums ...int) int{
	minimumValue := nums[0]
	for _,value := range nums{
		if value<minimumValue{
			minimumValue = value
		}
	}
	return minimumValue
}
func main()  {
	fmt.Print("hello")

}