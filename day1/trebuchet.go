package trebuchet

import (
    "fmt"
    "strconv"
    "regexp"
)

func TrebuchetPart1(document []string) int {
    response := 0
    results := []int{}

    for _, v := range(document) {
        digits := make([]int, 0)   

        for _, c := range(v) {
            if digit, err := strconv.Atoi(string(c)); err == nil {
                
                digits = append(digits, digit) 
            }
        }
        
        if len(digits) == 1 {
            digits = append(digits, digits[0])
        }

        results = append(results, getNum(digits))
    }

    for _, i := range(results) {
        response += i 
    }

    return response
}

func getNum(s []int) int {
    if len(s) > 1 {
        s = []int{s[0], s[len(s)-1]} 
    } else {
        s = []int{s[0],s[0]}
    }
    res := 0
    op := 1

    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= 10
    }

    return res
}

// Can vastly improve the above by having a pointer at the beginning and a pointer at the end of every string and only 
// searching until two numbers are found instead of finding each one -- may implement that may not.


func TrebuchetPart2(document []string) int {
    res := 0
    result := []int{}
    stringToIntConversions := map[string]int{
        "one": 1,
        "two": 2,
        "three": 3, 
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }
    re, err := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
        
    if err != nil {
        panic("You don't know regex")
    }


    for _, v := range(document) {
        firstDigit := re.FindString(v)
        lastDigit := ""
        digits_int := make([]int,0)

        for lastDigit == "" {
            if len(v) < 5 {
                matches := re.FindAllString(v, -1)
                lastDigit = matches[len(matches)-1]
                break
            }
            for i := len(v); i >= 0; i-- {
                checkString := v[i-5:i]
                matches := re.FindAllString(checkString, -1)
                if len(matches) > 0 {
                    lastDigit = matches[len(matches)-1]
                }

                if lastDigit != "" {    
                    break
                }
            }
        }

        if lastDigit == "" {
            panic("No Number Found???")
        }

        digits := []string{firstDigit, lastDigit}

        for i, match := range(digits) {
            if len(match) > 1 {
                digits_int = append(digits_int, stringToIntConversions[match])
            } else {
                d, err := strconv.Atoi(digits[i])

                if err != nil {
                    panic("We can't count!!!")
                }

                digits_int = append(digits_int, d)
            }
        }

        result = append(result, getNum(digits_int))
    }

    for _, i := range(result) {
        res += i 
    }

    return res

}
