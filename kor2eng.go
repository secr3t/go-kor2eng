package go_kor2eng

import (
    "regexp"
    "strings"
)

const (
    korRegex            = "[ㄱ-ㅎㅏ-ㅣ가-힣]+"
    koreanStartUnicode  = 0xAC00
    koreanEndUnicode    = 0xD79F
    numOfMiddleSound    = 21
    numOfLastSound      = 28
    openBrace           = "{"
    closeBrace          = "}"
)

var (
    hasKorRegex = regexp.MustCompile(korRegex)
    isKorRegex  = regexp.MustCompile("^" + korRegex + "$")

    firstSound = []string{"r", "R", "s", "e", "E", "f", "a", "q", "Q", "t",
        "T", "d", "w", "W", "c", "z", "x", "v", "g"}

    middleSound = []string{"k", "o", "i", "O", "j", "p", "u", "P", "h", "hk",
        "ho", "hl", "y", "n", "nj", "np", "nl", "b", "m", "ml", "l"}

    lastSound = []string{"", "r", "R", "rt", "s", "sw", "sg", "e", "f", "fr", "fa",
        "fq", "ft", "fx", "fv", "fg", "a", "q", "qt", "t", "T",
        "d", "w", "c", "z", "x", "v", "g"}
)

func Kor2Eng(kor string) string {
    if !HasKorean(kor) {
        return kor
    }

    sb := strings.Builder{}

    for _, k := range kor {
        initUc := k - koreanStartUnicode

        if initUc >= 0 {
            firstLetter := getFirstLetter(initUc)
            sb.WriteString(firstLetter)

            middleLetter := getMiddleLetter(initUc)
            sb.WriteString(middleLetter)

            lastLetter := getLastLetter(initUc)
            if lastLetter != "" {
                sb.WriteString(lastLetter)
            }
        } else {
            sb.WriteRune(k)
        }

    }
    return sb.String()
}

func Kor2EngWithBraces(kor string) string {
    if !HasKorean(kor) {
        return kor
    }

    isPrevKor := false
    sb := strings.Builder{}

    for _, k := range kor {
        initUc := k - koreanStartUnicode

        if initUc >= 0 {
            if !isPrevKor {
                sb.WriteString(openBrace)
            }
            isPrevKor = true

            firstLetter := getFirstLetter(initUc)
            sb.WriteString(firstLetter)

            middleLetter := getMiddleLetter(initUc)
            sb.WriteString(middleLetter)

            lastLetter := getLastLetter(initUc)
            if lastLetter != "" {
                sb.WriteString(lastLetter)
            }
        } else {
            if isPrevKor {
                sb.WriteString(closeBrace)
            }
            isPrevKor = false

            sb.WriteRune(k)
        }
    }

    if isPrevKor {
        sb.WriteString(closeBrace)
    }

    return sb.String()
}

func getLastLetter(initUc int32) string {
    lastLetterVal := initUc % numOfLastSound
    lastLetter := lastSound[lastLetterVal]
    return lastLetter
}

func getMiddleLetter(initUc int32) string {
    middleLetterVal := initUc / numOfLastSound % numOfMiddleSound
    middleLetter := middleSound[middleLetterVal]
    return middleLetter
}

func getFirstLetter(initUc int32) string {
    firstLetterVal := initUc / numOfMiddleSound / numOfLastSound
    firstLetter := firstSound[firstLetterVal]
    return firstLetter
}

func HasKorean(text string) bool {
    return hasKorRegex.MatchString(text)
}

func IsTextAllKorean(text string) bool {
    return isKorRegex.MatchString(text)
}

func isKoreanUnicode(c rune) bool {
    return c >= koreanStartUnicode && c <= koreanEndUnicode
}
