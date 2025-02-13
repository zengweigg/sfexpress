package tool

const BlockSize int = 32

func PKCS7Padding(count int) []byte {
	amountToPad := BlockSize - (count % BlockSize)
	if amountToPad == 0 {
		amountToPad = BlockSize
	}
	b := chr(amountToPad)
	tmp := make([]byte, amountToPad)
	for index, _ := range tmp {
		tmp[index] = b[0]
	}
	return GetBytes(string(tmp))
}

func chr(n int) []byte {
	b := (byte)(n & 0xFF)
	return []byte{b}
}
