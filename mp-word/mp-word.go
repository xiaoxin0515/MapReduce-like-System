package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("nonono, it can't be read")

	} else {
		f := strings.Split(string(data), "\n")
		for _, line := range f {
			for _, word := range strings.Split(line[:], " ") {
				if word == "" {
					continue
				} else {
					fmt.Println(word + " " + "1")
				}
			}
		}
	}
}
