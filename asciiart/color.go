package asciiart

import (
	"fmt"
	"image/color"
	"os"
	"strings"
)

var Colors = map[string]color.RGBA{
	"black":           {R: 0, G: 0, B: 0, A: 255},
	"white":           {R: 255, G: 255, B: 255, A: 255},
	"red":             {R: 255, G: 0, B: 0, A: 255},
	"green":           {R: 0, G: 255, B: 0, A: 255},
	"blue":            {R: 0, G: 0, B: 255, A: 255},
	"yellow":          {R: 255, G: 255, B: 0, A: 255},
	"cyan":            {R: 0, G: 255, B: 255, A: 255},
	"magenta":         {R: 255, G: 0, B: 255, A: 255},
	"silver":          {R: 192, G: 192, B: 192, A: 255},
	"gray":            {R: 128, G: 128, B: 128, A: 255},
	"maroon":          {R: 128, G: 0, B: 0, A: 255},
	"olive":           {R: 128, G: 128, B: 0, A: 255},
	"darkgreen":       {R: 0, G: 100, B: 0, A: 255},
	"purple":          {R: 128, G: 0, B: 128, A: 255},
	"teal":            {R: 0, G: 128, B: 128, A: 255},
	"navy":            {R: 0, G: 0, B: 128, A: 255},
	"orange":          {R: 255, G: 165, B: 0, A: 255},
	"brown":           {R: 165, G: 42, B: 42, A: 255},
	"gold":            {R: 255, G: 215, B: 0, A: 255},
	"pink":            {R: 255, G: 192, B: 203, A: 255},
	"lightblue":       {R: 173, G: 216, B: 230, A: 255},
	"lightgreen":      {R: 144, G: 238, B: 144, A: 255},
	"lightyellow":     {R: 255, G: 255, B: 224, A: 255},
	"darkblue":        {R: 0, G: 0, B: 139, A: 255},
	"darkred":         {R: 139, G: 0, B: 0, A: 255},
	"darkcyan":        {R: 0, G: 139, B: 139, A: 255},
	"darkmagenta":     {R: 139, G: 0, B: 139, A: 255},
	"darkorange":      {R: 255, G: 140, B: 0, A: 255},
	"lightgray":       {R: 211, G: 211, B: 211, A: 255},
	"darkgray":        {R: 169, G: 169, B: 169, A: 255},
	"lightpink":       {R: 255, G: 182, B: 193, A: 255},
	"lightgoldenrod":  {R: 250, G: 250, B: 210, A: 255},
}

func Color(str string) string {
	color := os.Args[1]
	if !strings.HasPrefix(color, "--color=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> \"something\"")
		os.Exit(0)
	} else {

	}
	color = color[8:]
	return color
}
