package cube

import (
    "fmt"
    "strconv"
    "regexp"
    "strings"
)

func CubeSolution1(games []string) int {
    maxColors := map[string]int{
        "red" : 12,
        "green" : 13,
        "blue" : 14,
    }

    // keeping track of impossible games by the game number
    impossibleGames := make([]int, 0)

    numRe := regexp.MustCompile(`\d+`)
   
    for _, game := range(games) {
        colors := []string{"red", "green", "blue"}
        isValid := true
        gameNum, _ := strconv.Atoi(numRe.FindString(game))

        for i := 0; i < len(game) && isValid; i++{
            for _, k := range colors {
                re := regexp.MustCompile(`\d+ `+ k)
                matches := re.FindAllString(game, -1)

                for _, match := range(matches) {
                    spaceIndex := strings.Index(match, " ")

                    val, _ := strconv.Atoi(match[:spaceIndex])
                    
                    if val > maxColors[k] {
                        isValid = false
                    }
                }
            }
        }

        if isValid {
            impossibleGames = append(impossibleGames, gameNum)
        }
        

    }

    fmt.Println(impossibleGames)

    sumGamenums := 0

    for i := 0; i < len(impossibleGames); i++ {
        sumGamenums += impossibleGames[i]
    }

    return sumGamenums
}

func CubeSolution2(games []string) int {
    sumGamePowers := 0

    for _, game := range(games) {
        maxEachColor := map[string]int{
            "red" : 0,
            "green" : 0,
            "blue" : 0,
        }

        for i := 0; i < len(game); i++{
            for k := range maxEachColor {
                re := regexp.MustCompile(`\d+ `+ k)
                matches := re.FindAllString(game, -1)

                for _, match := range(matches) {
                    spaceIndex := strings.Index(match, " ")

                    val, _ := strconv.Atoi(match[:spaceIndex])
                    
                    if val > maxEachColor[k] {
                        maxEachColor[k] = val
                    }
                }
            }
        }

        sumGamePowers += maxEachColor["red"] * maxEachColor["green"] * maxEachColor["blue"] 
    }

    return sumGamePowers
}
