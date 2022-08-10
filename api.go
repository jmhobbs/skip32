package skip32


func Encrypt(key [10]uint8, input uint32) uint32 {
	return skip32(key, input, 0, true)
}

func Decrypt(key [10]uint8, input uint32) uint32 {
	return skip32(key, input, 23, false)
}

func KeyFromSlice(inputKey []byte) [10]byte {
	if len(inputKey) < 10 {
		inputKey = append(make([]byte, 10 - len(inputKey)), inputKey...)
	}
	key := [10]byte{}
	copy(key[:], inputKey)
	return key
}