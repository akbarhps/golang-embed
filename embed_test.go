package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"testing"
)

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}

func TestEmbedSliceByte(t *testing.T) {
	fmt.Println(len(logo))

	err := ioutil.WriteFile("logo_baru.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("logo_baru.jpg saved")
}

func TestDeleteLogoBaru(t *testing.T) {
	err := os.Remove("logo_baru.jpg")
	if err != nil {
		panic(err)
	}

	fmt.Println("logo_baru.jpg removed")
}

func TestEmbedMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

func TestEmbedPathMatcher(t *testing.T) {
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
