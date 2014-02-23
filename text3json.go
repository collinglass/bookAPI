package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	myBigThing := make(map[string]map[string]string)
	f, _ := os.Open("strangecountess.txt")
	r := bufio.NewReader(f)
	var currentPage map[string]string
	var nextChar []byte
	var line string
	var err error
	pageNum := 0
	for {
		line, err = r.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("Error in parsing :", err)
			}
			break
		}
		nextChar, err = r.Peek(1)
		if err != nil {
			if err != io.EOF {
				log.Println("Error in peek :", err)
			}
			break
		}
		r, _ := utf8.DecodeRune(nextChar)
		if r == '\n' {
			if currentPage == nil {
				currentPage = make(map[string]string)
				myBigThing[fmt.Sprintf("page%d", pageNum)] = currentPage
				pageNum++
			} else if line == "" {
				currentPage = nil
			} else {
				tokens := strings.Split(line, ":")
				if len(tokens) == 2 {
					currentPage[tokens[0]] = tokens[1]
				}
			}
		}
	}
	f, err = os.Create("tester.json")
	if err != nil {
		log.Println("Error :", err)
		return
	}
	defer f.Close()
	bout, _ := json.Marshal(myBigThing)
	f.Write(bout)
}
