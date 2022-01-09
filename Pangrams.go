package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
    // Write your code here
    
    char := strings.Split(strings.TrimSpace(s), "")
    
    hashMap := make(map[string]int,0)
    for _, val := range  char{
       val = strings.ToLower(val)
        if _,ok := hashMap[val];ok{
            hashMap[val]=1
        }else{
            hashMap[val]=1
        }
    }
    fmt.Println(hashMap)
    //An extra char count is added because we initialized hashMap as not nil
    if len(hashMap)==27{
        return "pangram"
    }
    return "not pangram"
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := pangrams(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
