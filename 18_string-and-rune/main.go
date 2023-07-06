package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "こんにちは"

	fmt.Println("Len:", len(s)) // 15

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // 5

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'こ' {
		fmt.Println("found こ")
	}
}
