var hashTable map[uint32] uint32

func reverseBits(num uint32) uint32 {
    hashTable = map[uint32] uint32{}
    var res uint32
    res = 0
    i := 0
    for num > 0{
        res *= 2
        res += num % 2
        num = uint32(num / 2)
        i++
    }
    for i < 32{
        i++
        res *= 2
    }
    hashTable[num] = res
    return res
}