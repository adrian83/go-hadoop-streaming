package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
    "io"
    "strconv"
)

func main() {
	bio := bufio.NewReader(os.Stdin)
counts := make(map[string]int)
    
    for true {
        
        line, err := bio.ReadString('\n')
    
        if err != nil  {
            if err != io.EOF {
                fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err)
            }
            // io.EOF
            break

            
        }

        parts := strings.Split(line, "\t")
        
        word := parts[0]
        countStr := parts[1]
        
        count, err := strconv.Atoi(countStr)
        if err != nil {
              fmt.Fprintf(os.Stderr, "not a number! ", err)  
        
        break
        }
        counts[word] = counts[word] + count
        
    }
    for word, count := range counts {
                fmt.Printf("%s\t%d\n", word, count)
        }
}
    
    
    
