package main

import (
    "fmt"
    "github.com/aevanra/advent-of-code-2023/day1"
    "github.com/aevanra/advent-of-code-2023/day2"
    "github.com/aevanra/advent-of-code-2023/day3"
    "github.com/aevanra/advent-of-code-2023/day4"
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
    
    cubeGameInput := get_input("day2/input.txt")
    fmt.Println("Cube Game Problem Part 1")
    fmt.Println(cube.CubeSolution1(cubeGameInput))
    fmt.Println("Cube Game Problem Part 2")
    fmt.Println(cube.CubeSolution2(cubeGameInput))

    gondola_input := get_input("day3/input.txt")
    fmt.Println("Gondola Problem Part 1")
    fmt.Println(gondola.GondolaPart1(gondola_input))
    fmt.Println("Gondola Problem Part 2")
    fmt.Println(gondola.GondolaPart2(gondola_input))

    scratchCardsInput := get_input("day4/input.txt")
    fmt.Println("Scratch Cards Problem Part 1")
    fmt.Println(scratchCards.ReadScratchCardsPart1(scratchCardsInput))
    fmt.Println("Scratch Cards Problem Part 2")
    fmt.Println(scratchCards.ReadScratchCardsPart2(scratchCardsInput))

}
