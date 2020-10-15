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
    converted := Kor2Eng("까 a b 나")
    expected := "Rk a b sk"
    assaultString(t, converted, expected)
}

func TestKor2EngWithBraces(t *testing.T) {
    converted:=Kor2EngWithBraces("안 h i 녕 하")
    expected := "{dks} h i {sud} {gk}"
    assaultString(t, converted, expected)
}

func assaultString(t *testing.T, converted string, expected string) {
    t.Logf("converted: %s, expected: %s", converted, expected)
    if converted != expected {
        t.Error("converted not equals with expected.")
    }
}