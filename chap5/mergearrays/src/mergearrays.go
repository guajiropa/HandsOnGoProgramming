/*
** AUTHOR	: Robert James Patterson
** DATE		: 09/06/2019
** SYNOPSIS	: We are going to see how to merge two arrays easily in Go. Consider that we have two
** arrays and we are going to merge them. If you used append earlier, you will know that it can
** accept as many parameters as required.
 */
package main

import (
	"fmt"
)

func main() {
	items1 := []int{3, 4}
	items2 := []int{1, 2}

	// The 'append' function expects a slice and a list of values, we however, are using append
	// to merge two slices. This will require us to unpack the values from the second slice, we
	// do this by using the ellipse you see after 'items2'. This is refered to in GO as a
	// 'Variadic function' and you can find out more about this at:
	// [https://gobyexample.com/variadic-functions].
	results := append(items1, items2...)

	fmt.Println(results)
}
