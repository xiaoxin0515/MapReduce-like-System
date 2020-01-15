package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	flag := true
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("nonono, we can't read it")

	} else {
		f := strings.Split(string(data), "\n")
		for _, line := range f {
			if line == "" {
				continue
			} else {
				l := strings.TrimSpace(line)
				if l == "" {
					continue
				}
				link := strings.SplitN(l, " ", 2)
				if flag {
					fmt.Println(link[0])
					flag = false
				}
				fmt.Println(link[1])
			}
		}
	}
}
