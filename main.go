package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed files/version.txt
var version string

//go:embed files/img.jpg
var logo []byte

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_baru.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}

		fmt.Println(entry.Name())
		content, _ := path.ReadFile("files/" + entry.Name())
		fmt.Println(string(content))
	}
}
