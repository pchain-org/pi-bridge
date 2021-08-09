package merkle

// return the number of 1 bit
func countBit(num uint32) uint {
	var count uint
	for num != 0 {
		num &= (num - 1)
		count += 1
	}
	return count
}

func isPower2(num uint32) bool {
	return countBit(num) == 1
}

// return the position of the heightest 1 bit
// 1-based index
func highBit(num uint32) uint {
	var hiBit uint
	for num != 0 {
		num >>= 1
		hiBit += 1
	}
	return hiBit
}

// return the position of the lowest 1 bit
// 1-based index
func lowBit(num uint32) uint {
	return highBit(num & -num)
}
