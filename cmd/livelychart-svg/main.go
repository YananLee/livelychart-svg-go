package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	livelychart "github.com/YananLee/livelychart-svg-go"
)

func main() {
	in := flag.String("in", "", "CSV file with label,value rows (stdin if omitted)")
	out := flag.String("out", "", "output SVG file (stdout if omitted)")
	flag.Parse()
	var reader io.Reader = os.Stdin
	if *in != "" {
		file, err := os.Open(*in)
		if err != nil {
			fail(err)
		}
		defer file.Close()
		reader = file
	}
	rows, err := csv.NewReader(reader).ReadAll()
	if err != nil {
		fail(err)
	}
	bars := make([]livelychart.Bar, 0, len(rows))
	for i, row := range rows {
		if len(row) != 2 {
			fail(fmt.Errorf("row %d must contain label,value", i+1))
		}
		value, err := strconv.ParseFloat(strings.TrimSpace(row[1]), 64)
		if err != nil && i == 0 && strings.EqualFold(strings.TrimSpace(row[1]), "value") {
			continue
		}
		if err != nil {
			fail(fmt.Errorf("row %d: %w", i+1, err))
		}
		bars = append(bars, livelychart.Bar{Label: row[0], Value: value})
	}
	svg, err := livelychart.Render(bars, livelychart.Options{})
	if err != nil {
		fail(err)
	}
	if *out == "" {
		fmt.Println(svg)
		return
	}
	if err := os.WriteFile(*out, []byte(svg), 0644); err != nil {
		fail(err)
	}
}

func fail(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
