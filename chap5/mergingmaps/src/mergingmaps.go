/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/06/2019
** SYNOPSIS	: In this lesson we will learn how to merge maps. As you can see, there are four items,
** and the maps are basically mapping a string to an integer. Let's go ahead and merge these two
** maps together. Unfortunately, there's no built-in way of doing this, so all we have to do is
** just to iterate these two maps and then merge them together.
 */
package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{
		"Michael": 10,
		"Jessica": 20,
		"Tarik":   33,
		"Jon":     22,
	}
	fmt.Println(map1)

	map2 := map[string]int{
		"Lord":  11,
		"Of":    22,
		"The":   36,
		"Rings": 23,
	}

	for key, value := range map2 {
		map1[key] = value
	}
	fmt.Println(map1)
}
