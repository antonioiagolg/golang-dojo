package iteration

import "strings"

func Repeat(a string, repeatCount int) string {

    var repeated strings.Builder
    for range repeatCount {
        repeated.WriteString(a)
    }

    return repeated.String()
}
