# LivelyChart SVG for Go

A dependency-free Go library and CLI that renders a basic animated horizontal bar chart SVG from nonnegative values. It runs locally and does not use the LivelyChart service. For spreadsheet import, chart templates, and video export, use [LivelyChart](https://www.livelychart.com/).

## Install

```sh
go get github.com/YananLee/livelychart-svg-go@v0.1.0
```

## Go API

```go
bars := []livelychart.Bar{{Label: "Nova", Value: 94}, {Label: "Orbit", Value: 88}}
svg, err := livelychart.Render(bars, livelychart.Options{})
if err != nil { panic(err) }
_ = os.WriteFile("chart.svg", []byte(svg), 0644)
```

Import `livelychart "github.com/YananLee/livelychart-svg-go"` and `os`.

## CLI

```sh
printf 'label,value\nNova,94\nOrbit,88\n' | go run github.com/YananLee/livelychart-svg-go/cmd/livelychart-svg@v0.1.0 -out chart.svg
```

Use `-in data.csv` for a file or omit `-out` to print SVG. The package renders a single bar chart; it does not reproduce the website's chart race or media export features.

## License

MIT.
