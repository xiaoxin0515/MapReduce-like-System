package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("nonono, we can't read it")

	} else {
		var word string
		current_word := ""
		current_count := 0
		f := strings.Split(string(data), "\n")
		for _, line := range f {
			if line == " " || line == "" {
				continue
			}
			l := strings.TrimSpace(line)
			if l == "" {
				continue
			}
			seg := strings.SplitN(l, " ", 2)
			//ioutil.WriteFile("subprocess.out", []byte("we got"+seg[0]), 0644)
			//ioutil.WriteFile("subprocess.out", []byte("we got"+seg[1]), 0644)
			word = seg[0]
			count, err := strconv.Atoi(seg[1])
			if err != nil {
				log.Fatal(err)
			}
			if current_word == word {
				current_count += count
			} else {
				if current_word != "" {
					fmt.Println(current_word, current_count)
				}
				current_count = count
				current_word = word
			}
		}
		if current_word == word {
			fmt.Println(current_word, current_count)
		}
	}
}
