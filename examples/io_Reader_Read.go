package main

import (
    "strings"
    "io"
    "fmt"
)

func main() {
    reader := strings.NewReader("123456789") //*strings.Reader
    buf := make([]byte, 4)
    for {
        n, err := reader.Read(buf)
        fmt.Printf("n=%v,err=%v,buf=%q\n", n, err, buf[:n]) 
        if err == io.EOF {
            break
        }
    }
    /**
        output:
            n=4,err=<nil>,buf="1234"
            n=4,err=<nil>,buf="5678"
            n=1,err=<nil>,buf="9"
            n=0,err=EOF,buf=""
     */
}
