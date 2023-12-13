package main

import (
    "fmt"
    "github.com/aevanra/advent-of-code-2023/day1"
    "bufio"
    "os"
)

func main() {
    day1Input, err := os.Open("day1/input.txt")

    if err != nil {
        panic("We cannot read!!!")
    }

    scanner := bufio.NewScanner(day1Input)
    scanner.Split(bufio.ScanLines)

    trebuchet_input := make([]string, 0)

    for scanner.Scan() {
        trebuchet_input = append(trebuchet_input, scanner.Text())
    }

    fmt.Println("Trebuchet Problem Part 1")
    fmt.Println(trebuchet.TrebuchetPart1(trebuchet_input))
    
    fmt.Println("\nTrebuchet Problem Part 2")
    fmt.Println(trebuchet.TrebuchetPart2(trebuchet_input))
    
}
