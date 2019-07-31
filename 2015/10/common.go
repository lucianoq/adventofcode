package main

import "strconv"

// Fast, preallocating buffer
func EncodeIterBuf(s string) string {
	buf := make([]byte, 0, len(s)*2)

	prevChar := s[0]
	tempCount := 1

	for i := 1; i < len(s); i++ {
		if s[i] == prevChar {
			tempCount++
			continue
		}

		buf = append(buf, []byte(strconv.Itoa(tempCount))...)
		buf = append(buf, prevChar)

		tempCount = 1
		prevChar = s[i]
	}
	buf = append(buf, []byte(strconv.Itoa(tempCount))...)
	buf = append(buf, prevChar)
	return string(buf)
}

// Slow due to string allocations
func EncodeIter(s string) string {
	var resultString string

	prevChar := s[0]
	tempCount := 1

	for i := 1; i < len(s); i++ {
		if s[i] == prevChar {
			tempCount++
			continue
		}
		resultString += strconv.Itoa(tempCount) + string(prevChar)
		tempCount = 1
		prevChar = s[i]
	}
	resultString += strconv.Itoa(tempCount) + string(prevChar)
	return resultString
}

// Very very slow
func EncodeRecursive(s string) string {
	return encode(s, s[0], 0)
}

func encode(s string, prevChar uint8, countChar int) string {
	if len(s) == 0 {
		if countChar != 0 {
			return strconv.Itoa(countChar) + string(prevChar)
		} else {
			return ""
		}
	}

	if s[0] != prevChar {
		return strconv.Itoa(countChar) + string(prevChar) + encode(s[1:], s[0], 1)
	}

	return encode(s[1:], prevChar, countChar+1)
}
