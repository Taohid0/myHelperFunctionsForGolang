package main

func maxInt(nums ...int) int{
	maximumValue := nums[0]
	for _,value := range nums{
		if value>maximumValue{
			maximumValue = value
		}
	}
	return maximumValue
}

