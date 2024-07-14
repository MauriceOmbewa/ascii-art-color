package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var colour, input, banner, substr string
	banner = "standard"

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . --color=<color> [substring] <string> [banner]")
		fmt.Println("Example: go run . --color=red \"Hello World!\" standard")
		fmt.Println("Example: go run . --color=blue Hello")
		fmt.Println("Example: go run . --color=green \"Hello\" \"Hello World!\" standard")
		os.Exit(0)
	}

	args := os.Args[1:]
	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			colour = strings.TrimPrefix(arg, "--color=")
		} else {
			if input == "" {
				input = arg
			} else if substr == "" {
				substr = input
				input = arg
			} else {
				banner = arg
			}
		}
	}

	// INPUT____________________
	switch input {
	case "":
		return
	case "\\a", "\\0", "\\f", "\\v", "\\r":
		fmt.Println("Error: Non printable character", input)
		return
	}

	input = strings.ReplaceAll(input, "\\t", "    ")
	input = strings.ReplaceAll(input, "\\b", "\b")
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Logic process for handling the backspace.
	for i := 0; i < len(input); i++ {
		indexB := strings.Index(input, "\b")
		if indexB > 0 {
			input = input[:indexB-1] + input[indexB+1:]
		}
	}

	// Split our input text to a string slice and separate with a newline.
	words := strings.Split(input, "\n")

	// BANNER_____________________
	// Check if the banner has an extension
	if strings.Contains(banner, ".") {
		// Check if the extension is not .txt
		if !strings.HasSuffix(banner, ".txt") {
			fmt.Println("Error: Required format: banner.txt")
			return
		}
	} else {
		// If no extension, add .txt
		banner = banner + ".txt"
	}
	banner = strings.ToLower(banner)

	bannerFile := banner

	// Read the contents of the banner file.
	bannerText, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	// Confirm file information.
	fileInfo, err := os.Stat(bannerFile)
	if err != nil {
		fmt.Println("Error reading file information", err)
		return
	}
	fileSize := fileInfo.Size()

	if fileSize == 6623 || fileSize == 4702 || fileSize == 7462 || fileSize == 4496 {
		// Split the content to a string slice and separate with newline.
		contents := strings.Split(string(bannerText), "\n")

		asciiArt := ""
		if substr == ""{
			asciiArt = AsciiArts(words, contents)
			colorizeTexts(asciiArt, colour)
		} else {
			//asciiArt = AsciiArt(words, contents, substr)
			AsciiArt(words, contents, substr)
			//coloredOutput := colorizeText(asciiArt, colour, substr)
			//fmt.Print(coloredOutput)
		}
		//asciiArt := AsciiArt(words, contents, substr)
		// coloredOutput := colorizeText(asciiArt, colour, substr)
		// fmt.Print(coloredOutput)
	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}
}

// AsciiArt generates ASCII art based on input and banner contents
func AsciiArt(input []string, contents []string, substr string) {
    final := ""
    //countSpace := 0
	for _, word := range input {
		cont := strings.Contains(word, substr)
		if cont {
			startIndex := strings.Index(word, substr)
			endIndex := startIndex + len(substr)
	
			// ANSI escape codes for coloring
			red := "\033[31m"
			reset := "\033[0m"
	
			// Print the string with the substring in red
			for i := 0; i < 8; i++ {
				for m, char := range word {
					if char == '\n' {
						continue
					}
	
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						os.Exit(0)
					}
	
					if m >= startIndex && m < endIndex {
						fmt.Print(red + contents[int(char-' ')*9+1+i] + reset)
					} else {
						fmt.Print(contents[int(char-' ')*9+1+i])
					}
				}
				final += "\n"
				fmt.Println()
			}
		} else {
			reset := "\033[0m"
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						os.Exit(0)
					}
					fmt.Print(reset + contents[int(char-' ')*9+1+i])
				}
				fmt.Println()
			}
		}
	}
	
   // return final
}

// AsciiArt generates ASCII art based on input and banner contents
func AsciiArts(input []string, contents []string) string {
	final := ""
	countSpace := 0
	for _, word := range input {
		if word != "" {
			for i := 0; i < 8; i++ {
				for _, char := range word {
					if char == '\n' {
						continue
					}
					if !(char >= 32 && char <= 126) {
						fmt.Println("Error: Input contains non-ASCII characters")
						os.Exit(0)
					}
					// Print the calculated index of 'char' Ascii Art in content2.
					final += contents[int(char-' ')*9+1+i]
				}
				final += "\n"
			}
		} else {
			countSpace++
			if countSpace < len(input) {
				final += "\n"
			}
		}
	}
	return final
}

// colorizeText applies color to the text based on the given color name and substring
func colorizeText(text, colour, substr string) string {
	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	reset := "\033[0m"

	if color, exists := colorMap[colour]; exists {
		return color + text + reset
	}
	return text
}

func colorizeTexts(text, colour string) {
	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	reset := "\033[0m"

	if color, exists := colorMap[colour]; exists {
		fmt.Println(color + text + reset)
	}
}