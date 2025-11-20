package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Get_Input(Text_To_Display string) string {
	Reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v:", Text_To_Display)
	Output, _ := Reader.ReadString('\n')
	Output = strings.TrimSpace(Output)
	return Output
}
func Get_MultiLine_Input(Text_2_Display string) string {
	Reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%v, make sure it ends with an empty line: ", Text_2_Display)
	var Input_Lines []string
	for {
		line, _ := Reader.ReadString('\n')
		line = strings.TrimRight(line, "\n\r")
		if line == "" {
			break
		}
		Input_Lines = append(Input_Lines, line)
	}
	return strings.Join(Input_Lines, "\n")
}

func Convert_To_English(Input_String string) string {
	var out strings.Builder
	for _, r := range Input_String {
		//if r >= 0x0590 && r <= 0x05FF {
		switch r {
		case 'ק':
			out.WriteString("e")
		case 'ר':
			out.WriteString("r")
		case 'א':
			out.WriteString("t")
		case 'ט':
			out.WriteString("y")
		case 'ו':
			out.WriteString("u")
		case 'ן':
			out.WriteString("i")
		case 'ם':
			out.WriteString("o")
		case 'פ':
			out.WriteString("p")
		case 'ש':
			out.WriteString("a")
		case 'ד':
			out.WriteString("s")
		case 'ג':
			out.WriteString("d")
		case 'כ':
			out.WriteString("f")
		case 'ע':
			out.WriteString("g")
		case 'י':
			out.WriteString("h")
		case 'ח':
			out.WriteString("j")
		case 'ל':
			out.WriteString("k")
		case 'ך':
			out.WriteString("l")
		case 'ף':
			out.WriteString(":")
		case 'ז':
			out.WriteString("z")
		case 'ס':
			out.WriteString("x")
		case 'ב':
			out.WriteString("c")
		case 'ה':
			out.WriteString("v")
		case 'נ':
			out.WriteString("b")
		case 'מ':
			out.WriteString("n")
		case 'צ':
			out.WriteString("m")
		case 'ת':
			out.WriteString(",")
		case 'ץ':
			out.WriteString(".")
		case '/':
			out.WriteString("q")
		case '\'':
			out.WriteString("w")
		default:
			out.WriteRune(r)
		}
		/*} else {
			out.WriteRune(r)
		}*/
	}
	return out.String()
}
func Convert_To_Hebrew(Input_String string) (string, string) {
	Output := ""
	Input_String = strings.ToLower(Input_String)
	for _, Character := range Input_String {
		if (Character < 'a' || Character > 'z') && Character != ':' {
			Output += string(Character)
			continue
		}
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
			Output += ","
		case ".":
			Output += "."
		default:
			Output += fmt.Sprintf("%c", Character)
		}

	}
	return Output, Reverse_Line_By_Line(Output)
}
func Reverse_String(Input_String string) string {
	Runes := []rune(Input_String)
	for i, j := 0, len(Runes)-1; i < j; i, j = i+1, j-1 {
		Runes[i], Runes[j] = Runes[j], Runes[i]
	}
	return string(Runes)
}
func Reverse_Line_By_Line(Input_String string) string {
	lines := strings.Split(Input_String, "\n")
	for i, line := range lines {
		lines[i] = Reverse_String(line)
	}
	return strings.Join(lines, "\n")
}

func Show_Menu() {
	fmt.Println("0. Show this menu")
	fmt.Println("1. Convert to english")
	fmt.Println("2. Convert to hebrew")
	fmt.Println("3. Convert to english (Multiline)")
	fmt.Println("4. Convert to hebrew (Multiline)")
	fmt.Println("10. Exit")
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
			Text_To_Convert := Get_MultiLine_Input("Enter text to convert")
			Converted_Text := Convert_To_English(Text_To_Convert)
			fmt.Printf("Original:\n %s\nConverted: \n%s\n", Text_To_Convert, Converted_Text)
		case "4":
			Text_To_Convert := Get_MultiLine_Input("Enter text to convert")
			Converted_Text_For_Real, Converted_Text_For_Display := Convert_To_Hebrew(Text_To_Convert)
			fmt.Printf("Original:\n %s\nConverted: \n%s\nOrganized-Converted: \n%s\n", Text_To_Convert, Converted_Text_For_Real, Converted_Text_For_Display)
		case "5":
			Text := Get_Input("fgd")
			fmt.Println(len(Text))
		case "10":
			Running = false
			fmt.Println("Ok, goodbye!")
			time.Sleep(2 * time.Second)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
