package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "/Users/zhuchuanjun"
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
