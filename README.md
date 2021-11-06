# Golang Embed

Sumber Tutorial:
[Udemy](https://www.udemy.com/course/pemrograman-go-lang-pemula-sampai-mahir/learn/lecture/25663716#overview) |
[Slide](https://docs.google.com/presentation/d/1d7NZMrQwDvwRYZzqVgQKOAPuw_f7HEcIQdI6Rfvd2G4/edit#slide=id.p)


## Pengenalan Golang Embed
---


### Embed Package

- Sejak Golang versi 1.16, terdapat package baru dengan nama embed
- Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan isi file nya dalam variable
- https://golang.org/pkg/embed/ 


### Cek Versi Golang

```bash
go version
```


### Cara Embed File

- Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
- Selajutnya kita bisa tambahkan komentar `//go:embed` diikuti dengan nama file nya, diatas variable yang kita tuju
- Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara otomatis ketika kode golang di compile
- Variable yang dituju tidak bisa disimpan di dalam function


## Embed File ke String
---

- Embed file bisa kita lakukan ke variable dengan tipe data string
- Secara otomatis isi file akan dibaca sebagai text dan masukkan ke variable tersebut


### Kode: Embed File ke String

```go
import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
}
```


## Embed File ke []Byte
---

- Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
- Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain


### Kode: Embed File ke []Byte

```go
//go:embed img.jpg
var logo []byte

func TestEmbedSliceByte(t *testing.T) {
	fmt.Println(len(logo))

	err := ioutil.WriteFile("logo_baru.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	fmt.Println("logo_baru.jpg saved")
}
```


## Embed Multiple File
---

- Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
- Hal ini juga bisa dilakukan menggunakan embed package
- Kita bisa menambahkan komentar `//go:embed` lebih dari satu baris
- Selain itu variable nya bisa kita gunakan tipe data `embed.FS`


### Kode: Embed Multiple File

```go
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
```


## Path Matcher
---

- Selain manual satu per satu, kita bisa mengguakan path matcher untuk membaca multiple file yang kita inginkan
- Ini sangat cocok ketika misal kita punya pola jenis file yang kita inginkan untuk kita baca
- Caranya, kita perlu menggunakan path matcher seperti pada package function path.Match


### Dokumentasi `func Match()`

![Documentation](https://user-images.githubusercontent.com/69947442/140601662-f62c1ca1-2f1a-4bf4-93ab-08435ae93fd3.png)


### Kode: Path Matcher

```go
//go:embed files/*.txt
var path embed.FS

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
```


## Hasil Embed di Compile
---

- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file golang nya
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya, dan bahkan jika diubah file external nya, isi variable nya tidak akan berubah lagi


### Kode: Main Function

```go
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
```


### Build

```bash
go build
```