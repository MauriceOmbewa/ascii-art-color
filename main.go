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
	for k, arg := range args {
		if strings.HasPrefix(arg, "--color=") && k == 0{
			colour = strings.TrimPrefix(arg, "--color=")
		} else if strings.HasPrefix(arg, "--color ") && k==0 {
			colour = strings.TrimPrefix(arg, "--color ")
		} else if (!strings.HasPrefix(arg, "--color=") || !strings.HasPrefix(arg, "--color ")) && k == 0 {
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
			os.Exit(0)
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
	words := strings.Split(input, "\r\n")

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

	if fileSize == 6623 || fileSize == 5556 || fileSize == 7463 {
		// Split the content to a string slice and separate with newline.
		var contents []string
		if banner == "thinkertoy.txt"{
			contents = strings.Split(string(bannerText), "\r\n")
		} else {
			contents = strings.Split(string(bannerText), "\n")
		}
		

		asciiArt := ""
		if substr == ""{
			asciiArt = AsciiArts(words, contents)
			colorizeTexts(asciiArt, colour)
		} else {
			AsciiArt(words, contents, substr, colour)
		}
	} else {
		fmt.Println("Error with the file size", fileSize)
		return
	}
}

// AsciiArt generates ASCII art based on input and banner contents
func AsciiArt(input []string, contents []string, substr, colour string) {
    color := getColor(colour)
	for _, word := range input {
		reset := "\033[0m"
		for i := 0; i < 8; i++ {
			for ind := 0; ind < len(word); ind++ {
				char := word[ind]
				if char == '\n' {
					continue
				}

				if !(char >= 32 && char <= 126) {
					fmt.Println("Error: Input contains non-ASCII characters")
					os.Exit(0)
				}

				// Check if the current position is the start of the substring
				if ind <= len(word)-len(substr) && word[ind:ind+len(substr)] == substr {
					// Print the substring in color
					for j := 0; j < len(substr); j++ {
						char := word[ind+j]
						fmt.Print(color + contents[int(char-' ')*9+1+i] + reset)
					}
					// Skip the length of the substring
					ind += len(substr) - 1
				} else {
					// Print the normal character
					fmt.Print(contents[int(char-' ')*9+1+i])
				}
			}
			fmt.Println()
		}
	}
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

func getColor(colour string) string {
	colorMap := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"bright_black":   "\033[90m",
		"bright_red":     "\033[91m",
		"bright_green":   "\033[92m",
		"bright_yellow":  "\033[93m",
		"bright_blue":    "\033[94m",
		"bright_magenta": "\033[95m",
		"bright_cyan":    "\033[96m",
		"bright_white":   "\033[97m",
	}
	if color, exists := colorMap[colour]; exists {
		return color
	} else {
		fmt.Println("Error: color choosed not available")
		os.Exit(0)
	}
	return "\033[0m" // default to reset if color not found
}
