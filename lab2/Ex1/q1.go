package main

import (
	"fmt"
	"sync"
)

func countCharacter(char rune, str string, ch chan map[rune]int) {
	frequency := make(map[rune]int)
	for _, c := range str {
		if c == char {
			frequency[char]++
		}
	}
	ch <- frequency
}

func main() {
	var input string
	fmt.Scanln(&input)
	ch := make(chan map[rune]int)
	var wg sync.WaitGroup

	for _, char := range input {
		wg.Add(1)
		go func(char rune) {
			defer wg.Done()
			countCharacter(char, input, ch)
		}(char)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	frequency := make(map[rune]int)
	for freq := range ch {
		for char, count := range freq {
			frequency[char] += count
		}
	}

	fmt.Println("Character frequency:")
	for char, count := range frequency {
		fmt.Printf("%c: %d\n", char, count)
	}
}
