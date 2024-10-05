package main

import (
    "bufio"
    "fmt"
    "os"
    "dict/getword"
    "os/exec"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    dir := "/home/silver/Study/etyma/mindmap/"
    for {
        fmt.Println("")
        randomWord := getWord.GetWord(dir);
        fmt.Println(randomWord, "\n\nremember it? y or n\n")
        fmt.Printf("> ")
        scanner.Scan()
        input := scanner.Text()

        if input == ":exit" {
            fmt.Println("exit succeed")
            break
        }
        if input == "y" {
            fmt.Println("")
            cmd := exec.Command("bash", "-c", "wd "+ randomWord)
            output, err := cmd.Output()
            if err != nil {
                fmt.Print("Error: ", err)
                return
            }
            fmt.Println(string(output))
            fmt.Printf("Is your memory correct? y or n\n> ")
            scanner.Scan()
            ifcorrect := scanner.Text()

            if ifcorrect == "n" {
                fmt.Printf("Add it to vocabulary book? y or n\n> ")
                scanner.Scan()
                add := scanner.Text()
                if add == "n" {
                }else {
                    // TODO vocabulary bood
                    fmt.Println("Add to vocabulary book succeed")
                }
            }
        }

    }
}
