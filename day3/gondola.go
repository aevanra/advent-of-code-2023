package gondola

import (
    "fmt"
    "strconv"
    "regexp"
    "strings"
)

func GondolaPart1(schematics []string) int {
    schemaSum := 0
    numRe, _ := regexp.Compile(`\d+`)
    symbolRe, _ := regexp.Compile(`[^\.\d]`)

    for i := 0; i < len(schematics); i++ {
        lastRow := ""
        thisRow := schematics[i]
        nextRow := ""
        if i > 0 {
            lastRow = schematics[i-1]
        }
        if i < len(schematics)-1 {
            nextRow = schematics[i+1]
        }
        foundNums := numRe.FindAllString(thisRow, -1)

        for _, num := range(foundNums) {
            baseIndex := strings.Index(thisRow, num)
            indicesToCheck := make([]int, 0)
            isValid := false

            for j := -1; j <= len(num); j++ {
                indexToCheck := baseIndex + j 

                if indexToCheck >= 0 && indexToCheck <= len(thisRow)-1 {
                    indicesToCheck = append(indicesToCheck, indexToCheck)
                }
            }

            fmt.Println(indicesToCheck)
            for _, index := range(indicesToCheck) {
                fmt.Println(index)
                if !isValid {
                    thisString := symbolRe.FindString(string(thisRow[index]))
                    if thisString != "" {
                        isValid = true
                        break
                    }

                    if lastRow != "" {
                        thisString = symbolRe.FindString(string(lastRow[index]))
                    }
                    if thisString != "" {
                        fmt.Println(string(lastRow[index]))
                        isValid = true
                        break
                    }
                    
                    if nextRow != "" {
                        thisString = symbolRe.FindString(string(nextRow[index]))
                    }
                    if thisString != "" {
                        fmt.Println(string(nextRow[index]))
                        isValid = true
                        break
                    }
                }
            }

            if isValid {
                number, _ := strconv.Atoi(num)
                schemaSum += number
            }
        }
        
    }

    return schemaSum
}
