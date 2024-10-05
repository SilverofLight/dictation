package addtobook

import (
    "bufio"
    "fmt"
    "os"
)

func Add(word string, book string)  {
    file, err := os.OpenFile(book, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    // fmt.Println("word:", word)
    // fmt.Println("file:", book)
    if err != nil {
        fmt.Println("打开文件错误:", err)
        return
    }
    defer file.Close()
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
    
    fmt.Println("Add to vocabulary book succeed")
}
