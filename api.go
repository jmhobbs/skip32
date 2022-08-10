package skip32

// Encrypt obfuscates the input with Skip32 using the provided key.
func Encrypt(key [10]uint8, input uint32) uint32 {
	return skip32(key, input, 0, true)
}

// Decrypt reverses the effect of Skip32 using the provided key.
func Decrypt(key [10]uint8, input uint32) uint32 {
	return skip32(key, input, 23, false)
}

// KeyFromSlice is a utility function to convert a byte slice to
// a properly sized key.  If your input is less than 10 bytes, it
// will left pad the key with 0x00.
func KeyFromSlice(inputKey []byte) [10]byte {
	if len(inputKey) < 10 {
		inputKey = append(make([]byte, 10 - len(inputKey)), inputKey...)
	}
	key := [10]byte{}
	copy(key[:], inputKey)
	return key
}