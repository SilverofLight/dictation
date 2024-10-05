package main

import (
    "bufio"
    "fmt"
    "os"
    "dict/getword"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    dir := "/home/silver/Study/etyma/mindmap/"
    for {
        randomWord := getWord.GetWord(dir);
        fmt.Println(randomWord, "\nremember it? y or n")
        scanner.Scan()
        input := scanner.Text()

        if input == ":exit" {
            fmt.Println("exit succeed")
            break
        }

    }
}
