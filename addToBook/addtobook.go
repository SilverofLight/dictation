package addtobook

import (
	"bufio"
	"fmt"
	"os"
)

func Add(word string, book string)  {
    file, err := os.OpenFile(book, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
    // fmt.Println("word:", word)
    // fmt.Println("file:", book)
    if err != nil {
        fmt.Println("打开文件错误:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        if scanner.Text() == word {
            fmt.Println("/nAlready exist!!")
        }
        return
    }

    writer := bufio.NewWriter(file)
    newLine := word + "\n"
    
    _,err = writer.WriteString(newLine)
    if err != nil {
      fmt.Println("write err:", err)
      return
    }
    // 刷新缓冲区，确保写入文件
    if err := writer.Flush(); err != nil {
        fmt.Println("刷新错误:", err)
        return
    }
    
    fmt.Println("\nAdd to vocabulary book succeed")
}
