package utils

import "unicode/utf16"

func StringToUTF16(s string) (v []uint16) {
	return utf16.Encode([]rune(s))
}

func UTF16ToString(v []uint16) (s string) {
	return string(utf16.Decode(v))
}