package scratchCards

import (
    "math"
    "strings"
    "regexp"
    "fmt"
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

func addCard(key string, cards map[string]int, mult ...int) map[string]int {
    count, exists := cards[key]
    if !exists {
        count = 0
    }

    if len(mult) > 0 {
        count = count + (1*mult[0])
    } else {
        count++
    }


    cards[key] = count

    return cards
}

func ReadScratchCardsPart2(cards []string) int {
    numRe, _ := regexp.Compile(`\d+`)
    wonCards := make(map[string]int, 0)
    cardCount := 0

    for i := 0; i < len(cards); i++ {
        winCount := 0
        cardSplit := strings.Split(string(cards[i]), ":")
        cardID := cardSplit[0]
        numsSplit := strings.Split(string(cardSplit[1]), "|")
        winningNums := numRe.FindAllString(numsSplit[0], -1)
        hadNums := numRe.FindAllString(numsSplit[1], -1)
        wonCards = addCard(cardID, wonCards)
        cardMultiplier := wonCards[cardID]


        for _, num := range(hadNums) {
            if stringInSlice(num, winningNums) {
                winCount++
            }
        }

        for j:= 1; j <= winCount; j++ {
            winSplit := strings.Split(string(cards[i+j]), ":")
            cardID = winSplit[0]

            wonCards = addCard(cardID, wonCards, cardMultiplier)

        }
    }

    for _, v := range(wonCards) {
        cardCount += v
    }

    fmt.Println(wonCards)

    return cardCount
}
