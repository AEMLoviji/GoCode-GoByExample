package main

import "fmt"

// Assignment: create a map for data line "Golang": ".go" and search for some values.
func main() {
	fileExt := map[string]string{
		"Golang": ".go",
		"C++":    ".cpp",
		"Java":   ".java",
		"C#":     ".cs",
	}

	fmt.Println(fileExt)
	fmt.Println(len(fileExt))

	if ext, ok := fileExt["C#"]; ok {
		fmt.Println(ext, ok)
	}

	if ext, ok := fileExt["Ruby"]; ok {
		fmt.Println(ext, ok)
	} else {
		fmt.Println("Not found!")
	}
}
