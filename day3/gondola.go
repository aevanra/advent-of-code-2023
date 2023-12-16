package gondola

import (
    "strconv"
    "regexp"
)

func isSymbol(str string) bool {
    symbolRe, _ := regexp.Compile(`[^\.\d]`)
    chr := symbolRe.FindString(str)
    return chr != ""
}

func dedupeSlice(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func GondolaPart1(schematics []string) int {
    schemaSum := 0
    numRe, _ := regexp.Compile(`\d+`)

    for i := 0; i < len(schematics); i++ {
        lastRow := ""
        thisRow := "." + schematics[i] + "."
        nextRow := ""
        if i > 0 {
            lastRow = "." + schematics[i-1] + "."
        }
        if i < len(schematics)-1 {
            nextRow = "." + schematics[i+1] + "."
        }
        foundNums := numRe.FindAllString(thisRow, -1)
        foundNums = dedupeSlice(foundNums)

        for _, num := range(foundNums) {
            thisNumRe, _ := regexp.Compile(`[^\d]`+ num + `[^\d]`)
            numIndices := thisNumRe.FindAllStringIndex(thisRow, -1)

            for _, indices := range(numIndices) {
                startIndex, endIndex := indices[0], indices[1]-1
                thisNum, _ := strconv.Atoi(num)

                if (startIndex >= 0 && isSymbol(string(thisRow[startIndex]))) ||
                    (endIndex < len(thisRow) && isSymbol(string(thisRow[endIndex]))) {
                    schemaSum += thisNum
                    break
                }

                for j := startIndex; j <= endIndex; j++ {
                    if j < 0 || j >= len(thisRow) { continue }

                    if (len(lastRow) > 0 && isSymbol(string(lastRow[j]))) ||
                     (len(nextRow) > 0 && isSymbol(string(nextRow[j]))) {
                        schemaSum += thisNum
                        break
                    }
                }
            }
        }
    }
    return schemaSum
}
