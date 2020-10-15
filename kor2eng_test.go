package go_kor2eng

import (
    "fmt"
    "testing"
)

func TestRegex(t *testing.T) {
    fmt.Printf("%v\n", hasKorRegex.Split("안녕hi세요", -1))
    fmt.Printf("%v\n", hasKorRegex.MatchString("안녕hi세요"))
    fmt.Printf("%v\n", isKorRegex.MatchString("안녕hi세요"))
    fmt.Printf("%v\n", hasKorRegex.FindAllStringSubmatch("안녕hi세요", -1))
}

func TestKor2Eng(t *testing.T) {
    fmt.Println(Kor2Eng("까 a b 나"))
}

func TestKor2EngWithBraces(t *testing.T) {
    fmt.Println(Kor2EngWithBraces("안 h i 녕 하"))
}