func ReverseBits(oct byte) byte {
    var result uint8 = 0
    for i := 0; i < 8; i++ {
        result = (result << 1) | (oct & 1)
        x >>= 1
    }
    return result
}
