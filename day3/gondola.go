package gondola

import (
    "strconv"
    "regexp"
    "unicode"
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

func getNumbers(row string, index int, final bool) []int {
    if row == "" {
        return nil
    }

    nums := make([]int, 0)

    if unicode.IsDigit(rune(row[index])) {
        minIndex := index
        maxIndex := index
        for minIndex > 0 && unicode.IsDigit(rune(row[minIndex-1])) {
            minIndex--
        }
        for maxIndex < len(row)-1 && unicode.IsDigit(rune(row[maxIndex+1])) {
            maxIndex++
        }

        num, _ := strconv.Atoi(row[minIndex:maxIndex+1])

        return []int{num}
    }

    if index > 0 && unicode.IsDigit(rune(row[index - 1])) {
        if !final {
            nums = append(nums, getNumbers(row, index, true)...)
        }
    }

    if index < len(row) - 1 && unicode.IsDigit(rune(row[index+1])) {
        if !final {
            nums = append(nums, getNumbers(row, index, true)...)
        }
    }

    return nums
}

func GondolaPart2(schematics []string) int {
    ratioSum := 0

    for i := 0; i < len(schematics); i++ {
        rows := make([]string, 0)
        rows = append(rows, schematics[i])
        if i > 0 {
            rows = append(rows, schematics[i-1])
        }
        if i < len(schematics)-1 {
            rows = append(rows, schematics[i+1])
        }

        for ind, char := range(schematics[i]) {
            adjNums := make([]int, 0)
            if string(char) == "*" {
                for _, row := range(rows) {
                    adjNums = append(adjNums, getNumbers(row, ind, false)...)
                }

                if len(adjNums) == 2 {
                    ratioSum += adjNums[0] * adjNums[1]
                }
            }
        }
    }
    
    return ratioSum
}
