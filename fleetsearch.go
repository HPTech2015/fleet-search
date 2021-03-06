package main

import (
	"fleetsearch/dirman"
	"fmt"
)

func main() {
	ls := dirman.ListDir{}

	// List all.
	fmt.Println()
	fmt.Println("LIST ALL: ")
	objs, err := ls.LsAll("/")
	if err != nil {
		fmt.Println(err)
	}
	for _, obj := range objs {
		fmt.Println(obj.Mode(), obj.Size(), obj.ModTime(), obj.Name())
	}

	// List all that contain substring. 
	fmt.Println()
	fmt.Println("LIST CONTAINS: ")
	fileObjs, err := ls.LsContains("./", "git")
	if err != nil {
		fmt.Println(err)
	}
	for _, obj := range fileObjs {
		fmt.Println(obj.Mode(), obj.Size(), obj.ModTime(), obj.Name())
	}

	// List all that start with substring. 
	fmt.Println()
	fmt.Println("LIST STARTS WITH: ")
	fileObjs, err = ls.LsStartsWith("./", "fleet")
	if err != nil {
		fmt.Println(err)
	}
	for _, obj := range fileObjs {
		fmt.Println(obj.Mode(), obj.Size(), obj.ModTime(), obj.Name())
	}

	// List all that end with substring. 
	fmt.Println()
	fmt.Println("LIST ENDS WITH: ")
	fileObjs, err = ls.LsEndsWith("./", ".go")
	if err != nil {
		fmt.Println(err)
	}
	for _, obj := range fileObjs {
		fmt.Println(obj.Mode(), obj.Size(), obj.ModTime(), obj.Name())
	}
}
