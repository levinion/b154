package coding

var encodeList = map[rune]string{
	'0': "1",
	'1': "4",
	'2': "5",
}

var headEncodeList = []string{"1", "4", "5"}

var headDecodeList = map[byte]int{
	'1': 1,
	'4': 2,
	'5': 3,
}

var decodeList = map[rune]int{
	'1': 0,
	'4': 1,
	'5': 2,
}
