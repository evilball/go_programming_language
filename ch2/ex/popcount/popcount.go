package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

func ExpressionPopCount(x uint64) int {
	return int(pc[byte(x >> (0 * 8))] +
		pc[byte(x >> (1 * 8))] +
		pc[byte(x >> (2 * 8))] +
		pc[byte(x >> (3 * 8))] +
		pc[byte(x >> (4 * 8))] +
		pc[byte(x >> (5 * 8))] +
		pc[byte(x >> (6 * 8))] +
		pc[byte(x >> (7 * 8))])
}

func LoopPopCount(x uint64) int {
	var sum byte
	for i := byte(0); i < 8; i++ {
		sum += pc[byte(x >> (i * 8))]
	}
	return int(sum)
}

func ShiftPopCount(x uint64) int {
	var n int
	for i := byte(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func ClearRightPopCount(x uint64) int {
	var n int
	for x != 0 {
		x = x & (x - 1)
		n++
	}
	return n
}
