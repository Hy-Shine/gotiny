package container

func SwapCase(s byte) byte {
	if !IsValidLetter(s) {
		return s
	}
	return s ^ 0x20
}

func IsValidLetter(s byte) bool {
	return s >= 'A' && s <= 'Z' || s >= 'a' && s <= 'z'
}

func IsNumber(s byte) bool {
	return s >= '0' && s <= '9'
}
