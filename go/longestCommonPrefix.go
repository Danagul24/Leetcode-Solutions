package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	var prefix string
	if len(strs) == 0 {
		return prefix
	}

	for i := 0; i < len(strs[0]); i++ {
		char := string(strs[0][i])
		for j := 0; j < len(strs); j++ {
			if string(strs[j][i]) != char {
				return prefix
			}
		}
		prefix += char
	}
	return prefix
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	words, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(longestCommonPrefix(strings.Fields(words)))

}
