package main

import (
	"bufio"
	addtobook "dict/addToBook"
	"dict/getword"
	"fmt"
	"os"
	"os/exec"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    dir := "/home/silver/Study/etyma/mindmap/"
    voc_book := "/home/silver/Study/etyma/voc_book.md"
    lightblue := "\033[48;5;153m"
    black := "\033[30m"
    reset := "\033[0m"
    for {
        fmt.Println("")
        randomWord := getWord.GetWord(dir);
        fmt.Println(lightblue + black + "    " + randomWord, "                                                " + reset)
        fmt.Println("\n remember it? y or n\n")
        fmt.Printf("> ")
        // scanner.Scan()

        if !scanner.Scan() {
            if err := scanner.Err(); err != nil {
                fmt.Println("读取输入错误:", err)
            } else {
                fmt.Println("输入结束，退出...")
            }
            break
        }

        input := scanner.Text()

        if input == ":exit" {
            fmt.Println("exit succeed")
            break
        }

        if input == "n" {
            // don't remember
            // print wd 
            fmt.Println("")
            cmd := exec.Command("bash", "-c", "wd "+ randomWord)
            output, err := cmd.Output()
            if err != nil {
                fmt.Print("Error: ", err)
                return
            }
            fmt.Println(string(output))

            // ask if add
            fmt.Printf("Add it to vocabulary book? y or n\n> ")
            scanner.Scan()
            add := scanner.Text()
            if add == "n" {
            }else {
                addtobook.Add(randomWord, voc_book)
            }
        }else {
            // can remember
            // print wd 
            fmt.Println("")
            cmd := exec.Command("bash", "-c", "wd "+ randomWord)
            output, err := cmd.Output()
            if err != nil {
                fmt.Print("Error: ", err)
                return
            }
            fmt.Println(string(output))

            // ask if remember right
            fmt.Printf("Is your memory correct? y or n\n> ")
            scanner.Scan()
            ifcorrect := scanner.Text()

            // remember wrong
            if ifcorrect == "n" {
                fmt.Printf("Add it to vocabulary book? y or n\n> ")
                scanner.Scan()
                add := scanner.Text()
                if add == "n" {
                }else {
                    addtobook.Add(randomWord, voc_book)
                }
            }
        }

    }
}
