package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var total_lc, total_wc, total_cc, total_files int // counters total number of lines, words, characters and files
	for _, fname := range os.Args[1:] {
		var lc, wc, cc int // counters to line, words and characters of each file

		file, err := os.Open(fname)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		scan := bufio.NewScanner(file)

		for scan.Scan() {
			s := scan.Text()

			cc += len(s)
			wc += len(strings.Fields(s))
			lc++
		}

		fmt.Printf("|Lines: %7d |Words: %7d |Characters: %7d |%10s|\n", lc, wc, cc, fname)
		total_wc += wc
		total_lc += lc
		total_cc += cc

		total_files++
	}

	fmt.Printf("|Lines: %7d |Words: %7d |Characters: %7d |%4d files|\n", total_lc, total_wc, total_cc, total_files)
}
