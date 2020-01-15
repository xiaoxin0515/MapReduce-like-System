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
		fmt.Println("nonono, we can't read it")

	} else {
		f := strings.Split(string(data), "\n")
		for _, line := range f {
			if line == "" {
				continue
			} else {
				link := strings.SplitN(line, " ", 2)
				fmt.Println(link[0] + " " + link[1])
			}
		}
	}
}
