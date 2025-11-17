package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	RLE = "\u202B" // force RTL
	PDF = "\u202C" // end block
)

func Get_Input(Text_To_Display string) string {
	Reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v: ", Text_To_Display)
	Output, _ := Reader.ReadString('\n')
	Output = strings.TrimSpace(Output)
	return Output
}

func Convert_To_English(Input_String string) string {
	Output := ""

	for _, Character := range Input_String {
		switch string(Character) {
		case "/":
			Output += "q"
		case "'":
			Output += "w"
		case "ק":
			Output += "e"
		case "ר":
			Output += "r"
		case "א":
			Output += "t"
		case "ט":
			Output += "y"
		case "ו":
			Output += "u"
		case "ן":
			Output += "i"
		case "ם":
			Output += "o"
		case "פ":
			Output += "p"
		case "ש":
			Output += "a"
		case "ד":
			Output += "s"
		case "ג":
			Output += "d"
		case "כ":
			Output += "f"
		case "ע":
			Output += "g"
		case "י":
			Output += "h"
		case "ח":
			Output += "j"
		case "ל":
			Output += "k"
		case "ך":
			Output += "l"
		case "ף":
			Output += ":"
		case "ז":
			Output += "z"
		case "ס":
			Output += "x"
		case "ב":
			Output += "c"
		case "ה":
			Output += "v"
		case "נ":
			Output += "b"
		case "מ":
			Output += "n"
		case "צ":
			Output += "m"
		case "ת":
			Output += ","
		case "ץ":
			Output += "."
		default:
			Output += fmt.Sprintf("%c", Character)
		}

	}
	return Output
}
func Convert_To_Hebrew(Input_String string) (string, string) {
	Output := ""
	Input_String = strings.ToLower(Input_String)
	for _, Character := range Input_String {
		switch string(Character) {
		case "q":
			Output += "/"
		case "w":
			Output += "'"
		case "e":
			Output += "ק"
		case "r":
			Output += "ר"
		case "t":
			Output += "א"
		case "y":
			Output += "ט"
		case "u":
			Output += "ו"
		case "i":
			Output += "ן"
		case "o":
			Output += "ם"
		case "p":
			Output += "פ"
		case "a":
			Output += "ש"
		case "s":
			Output += "ד"
		case "d":
			Output += "ג"
		case "f":
			Output += "כ"
		case "g":
			Output += "ע"
		case "h":
			Output += "י"
		case "j":
			Output += "ח"
		case "k":
			Output += "ל"
		case "l":
			Output += "ך"
		case ":":
			Output += "ף"
		case "z":
			Output += "ז"
		case "x":
			Output += "ס"
		case "c":
			Output += "ב"
		case "v":
			Output += "ה"
		case "b":
			Output += "נ"
		case "n":
			Output += "מ"
		case "m":
			Output += "צ"
		case ",":
			Output += "ת"
		case ".":
			Output += "ץ"
		default:
			Output += fmt.Sprintf("%c", Character)
		}

	}
	return Output, Reverse_String(Output)
}
func Reverse_String(Input_String string) string {
	Runes := []rune(Input_String)
	for i, j := 0, len(Runes)-1; i < j; i, j = i+1, j-1 {
		Runes[i], Runes[j] = Runes[j], Runes[i]
	}
	return string(Runes)
}

func Show_Menu() {
	fmt.Println("0. Show this menu")
	fmt.Println("1. Convert to english")
	fmt.Println("2. Convert to hebrew")
}

func main() {
	Running := true
	Show_Menu()
	for Running {
		var User_Choice string
		fmt.Print("Choose an operation: ")
		fmt.Scanln(&User_Choice)
		switch User_Choice {
		case "0":
			Show_Menu()
		case "1":
			Text_To_Convert := Get_Input("Enter text to convert")
			Converted_Text := Convert_To_English(Text_To_Convert)
			fmt.Printf("Original: %s\nConverted: %s\n", Text_To_Convert, Converted_Text)
		case "2":
			Text_To_Convert := Get_Input("Enter text to convert")
			Converted_Text_For_Real, Converted_Text_For_Display := Convert_To_Hebrew(Text_To_Convert)
			fmt.Printf("Original: %s\nConverted: %s\nOrganized-Converted: %s\n", Text_To_Convert, Converted_Text_For_Real, Converted_Text_For_Display)
		case "3":

			str := "מה קורה "
			fmt.Println(RLE + str + PDF)
		case "10":
			Running = false
			fmt.Println("Ok, goodbye!")
			time.Sleep(2 * time.Second)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
