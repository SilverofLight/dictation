package getWord

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"

	// "strings"
	"time"
)

func GetWord(dir string) string {
    // dir := "/home/silver/Study/etyma/mindmap/" // 替换为你的目录路径

    // 读取目录内容
    entries, err := os.ReadDir(dir)
    if err != nil {
        fmt.Println("读取目录失败:", err)
        return ""
    }

    var fileNames []string
    for _, entry := range entries {
        if !entry.IsDir() {
            fileNames = append(fileNames, entry.Name())
        }
    }

    if len(fileNames) == 0 {
        fmt.Println("目录中没有文件")
        return ""
    }

    // 随机选择一个文件
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    randomFile := fileNames[rng.Intn(len(fileNames))]
    // fmt.Println("随机选择的文件:", randomFile)

    // 从随机选择的文件中随机选择一行
    file, err := os.Open(dir + "/" + randomFile)
    if err != nil {
        fmt.Println("打开文件失败:", err)
        return ""
    }
    defer file.Close()

    // 读取文件内容，跳过第一行
    scanner := bufio.NewScanner(file)
    if scanner.Scan() {
        // 跳过第一行
    }

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("读取文件失败:", err)
        return ""
    }

    if len(lines) == 0 {
        fmt.Println("文件中没有内容")
        return ""
    }

    // 随机选择一行
    randomLine := lines[rng.Intn(len(lines))]
    // fmt.Println("随机选择的行:", randomLine)
    for {
        randomLine = lines[rng.Intn(len(lines))]
        if randomLine != "" {
            break // 如果选中的行非空，退出循环
        }
        // fmt.Println("enpty, retry")
        // 如果选中的行为空，继续循环并重试
    }

    // 提取单词
    re := regexp.MustCompile(`\b\w+(-\w+)*\b`)
    matches := re.FindAllStringSubmatch(randomLine, -1)
    fmt.Println("\nmatches: ", matches)

    var words []string
    for _,match := range matches {
        if len(matches) >= 1 {
            words = append(words, match[0])
        }
    }

    fmt.Println("word: :", words)

    // for len(words)==0 {
    //   fmt.Println("error: words enpty")
    //   fmt.Println("filename: ", randomFile)
    //   words = re.FindAllString(randomLine, -1)
    // }
    randomWord := words[rng.Intn(len(words))]
    if len(words) > 0 {
        // 随机选择一个单词
        // fmt.Println("随机选择的单词:", randomWord)
    } else {
        fmt.Println("没有找到单词")
    }
    return randomWord
}
