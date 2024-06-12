func SwapBits(octet byte) byte {
	var a uint8 = octet
	var first uint8 = a << 4
	var second uint8 = a >> 4
	var final uint8 = first | second
	return final
}
