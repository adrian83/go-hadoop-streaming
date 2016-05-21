package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "io"
)

func main() {

	bio := bufio.NewReader(os.Stdin)
    
    for true {
        
        line, err := bio.ReadString('\n')
    
        if err != nil  {
            if err != io.EOF {
                fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err)
            }
            // io.EOF
            break

            
        }

        line = strings.TrimSpace(line)
        
        
        if len(line) > 0 {
            words := strings.Split(line, " ")

            for _, word := range words {
                word = strings.TrimSpace(word)
                if len(word) > 0 {
                    fmt.Println(fmt.Sprintf("%s\t1", word))
                }
            }
        }
    }


}
