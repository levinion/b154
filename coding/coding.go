package coding

import (
	"strconv"
)

// main encode
func EncodeToString[T []rune | string | []byte](str T) (r string) {
	rstr := []rune(string(str))
	for i := range rstr {
		rstr[i] = decimalToTrinary(rstr[i])
	}
	r += "114514"
	for _, v := range rstr {
		ori := strconv.Itoa(int(v))
		neck := runeReplace(rune(decimalToTrinary(len(ori))))
		head := headEncodeList[len(neck)-1]
		r += head + neck + runeReplace(v)
	}
	return r
}

func EncodeToBytes[T []rune | string | []byte](str T) (r []byte) {
	r = []byte(EncodeToString(str))
	return r
}

// main decode
func DecodeToString[T string | []byte | []rune](str T) (r string) {
	newStr := trimHead(str)
	r = string(decodeToRunes(newStr))
	return
}

func DecodeToBytes[T string | []byte | []rune](str T) (r []byte) {
	r = []byte(DecodeToString(str))
	return
}

func decodeToRunes[T string | []byte | []rune](str T) (r []rune) {
	s := string(str)
	if len(str) == 0 {
		return
	}
	neckLen := headDecodeList[s[0]]
	neck := trinaryToDecimal(s[1 : 1+neckLen])
	body := trinaryToDecimal(s[1+neckLen : 1+neckLen+neck])

	r = append(r, rune(body))
	r = append(r, decodeToRunes(s[1+neckLen+neck:])...)
	return
}

func trimHead[T string | []byte | []rune](str T) (r string) {
	r = string(str)[6:]
	return
}

// 三进制（114514表示法）到十进制
func trinaryToDecimal(s string) (t int) {
	bit := 1
	for i := len(s) - 1; i >= 0; i-- {
		t += decodeList[rune(s[i])] * bit
		bit *= 3
	}
	return
}

// 十进制到三进制的转换
func decimalToTrinary[T int | rune](s T) (t T) {
	for i := T(1); s != 0; i *= 10 {
		t += s % 3 * i
		s = s / 3
	}
	return
}

// 对三进制数进行114514替换
func runeReplace(s rune) (r string) {
	ori := strconv.Itoa(int(s))
	for _, v := range ori {
		r += encodeList[v]
	}
	return r
}
