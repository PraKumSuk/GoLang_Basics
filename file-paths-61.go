package main

import (
	"fmt"
	"path/filepath" //package
	"strings"
)

//linux dir/file
//win dir\file
func main() {

	//join construts path provided dir names
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	//check if path is absolute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	//Seperate file extension from name using ext
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	//Print file name without extn
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	//Rel finds a relative path between a base and a target.
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
