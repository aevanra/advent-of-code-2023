package main

import (
    "fmt"
    "github.com/aevanra/advent-of-code-2023/day1"
    "bufio"
    "os"
)

func get_input(file_name string) []string {
    input, err := os.Open(file_name)

    if err != nil {
        panic("We cannot read!!!")
    }

    scanner := bufio.NewScanner(input)
    scanner.Split(bufio.ScanLines)

    array_input := make([]string, 0)

    for scanner.Scan() {
        array_input = append(array_input, scanner.Text())
    }

    return array_input
}

func main() {
    trebuchet_input := get_input("day1/input.txt")
    fmt.Println("Trebuchet Problem Part 1")
    fmt.Println(trebuchet.TrebuchetPart1(trebuchet_input))
    
    fmt.Println("\nTrebuchet Problem Part 2")
    fmt.Println(trebuchet.TrebuchetPart2(trebuchet_input))
    
}
