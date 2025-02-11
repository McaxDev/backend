package utils

func GetBitByIndex(data int64, index uint) bool {
	return (data>>index)&1 == 1
}

func UpdateBitByIndex(data *int64, index uint, value bool) {
	if value {
		*data |= (1 << index)
	} else {
		*data &= ^(1 << index)
	}
}
