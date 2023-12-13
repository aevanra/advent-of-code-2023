package trebuchet

import (
    "strconv"
)

func Trebuchet(document []string) int {
    response := 0
    results := []int{}
    for _, v := range(document) {
        digits := make([]int, 2)   

        for _, c := range(v) {
            if digit, err := strconv.Atoi(string(c)); err != nil {
               digits = append(digits, digit) 
            }
        }

        if len(digits) == 1 {
            digits[1] = digits[0]
        }

        results = append(results, sliceToInt(digits))
    }

    for _, i := range(results) {
        response += i 
    }

    return response
}

func sliceToInt(s []int) int {
    res := 0
    op := 1

    for i := len(s) - 1; i >= 0; i-- {
        res += s[i] * op
        op *= 10
    }

    return res
}
