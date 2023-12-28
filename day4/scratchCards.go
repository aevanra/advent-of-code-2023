package scratchCards

import (
    "fmt"
    "math"
    "strings"
    "regexp"
)

func stringInSlice(val string, strings []string) bool {
    for _, string := range(strings) {
        if string == val {
            return true
        }
    }
    return false
}

func ReadScratchCardsPart1(cards []string) int {
    cardSum := 0
    numRe, _ := regexp.Compile(`\d+`)
    for i := 0; i < len(cards); i++ {
        winCount := 0
        cardSplit := strings.Split(string(cards[i]), ":")
        numsSplit := strings.Split(string(cardSplit[1]), "|")
        winningNums := numRe.FindAllString(numsSplit[0], -1)
        hadNums := numRe.FindAllString(numsSplit[1], -1)

        for _, num := range(hadNums) {
            if stringInSlice(num, winningNums) {
                winCount++
            }
        }

        if winCount > 0 {
            cardSum += int(math.Pow(2, float64(winCount - 1)))
        }
    } 
    return cardSum
}

func insertIntoSlice(slice []string, insertValues []string, index int) []string {
    sliceHead := slice[:index]
    sliceTail := slice[index:]
    returnSlice := make([]string, 0)

    returnSlice = append(sliceHead, insertValues...)
    returnSlice = append(returnSlice, sliceTail...)

    return returnSlice
}

func ReadScratchCardsPart2(cards []string) int {
    numRe, _ := regexp.Compile(`\d+`)
    fmt.Println(cards)
    for i := 0; i < len(cards); i++ {
        winCount := 0
        cardSplit := strings.Split(string(cards[i]), ":")
        numsSplit := strings.Split(string(cardSplit[1]), "|")
        winningNums := numRe.FindAllString(numsSplit[0], -1)
        hadNums := numRe.FindAllString(numsSplit[1], -1)
        insertRows := make([]string, 0)

        for _, num := range(hadNums) {
            if stringInSlice(num, winningNums) {
                winCount++
            }
        }

        for j := 1; j <= winCount; j++ {
            if i + j < len(cards) {
                insertRows = append(insertRows, cards[i+j]) 
            }
        }
        if i == 4 {
            for _, line := range(cards) {
                fmt.Println(line)
            }
            fmt.Println(len(cards))
            return 0
        }

        cards = insertIntoSlice(cards, insertRows, i+1)
    }

    return len(cards)
}
