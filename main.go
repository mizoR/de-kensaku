package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/mattn/go-runewidth"
)

func main() {
	flag.Parse()
	word := strings.Join(flag.Args(), " ")
	length := runewidth.StringWidth(word)

	fmt.Print("┏━")
	for i := 0; i < length; i++ {
		fmt.Print("━")
	}
	fmt.Println("━┳━━━━━┓")
	fmt.Print("┃ ")
	fmt.Print(word)
	fmt.Println(" ┃ 検索┃")
	fmt.Print("┗━")
	for i := 0; i < length; i++ {
		fmt.Print("━")
	}
	fmt.Println("━┻━━━━━┛")
}
