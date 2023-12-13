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
    re := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)

    for _, v := range(document) {
        digits := re.FindAllString(v, -1)
        digits_int := make([]int, 0)

        if digits == nil {
            continue
        } else if len(digits) == 1 {
            digits = append(digits, digits[0])
            fmt.Println(digits)
        } else {
            digits = []string{digits[0], digits[len(digits)-1]}
        }

        if len(digits) > 2 {
            fmt.Println(digits)
        }

        for i, match := range(digits) {

            if i > 1 {
                fmt.Println(i)
            }

            if len(match) > 2 {
                
                if stringToIntConversions[match] > 9 {
                    fmt.Println(stringToIntConversions[match])
                }

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
