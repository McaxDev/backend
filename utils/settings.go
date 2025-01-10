package utils

func GetBitByIndex(data int64, index int) bool {
	return (data>>index)&1 == 1
}

func GetBitByName(data int64, name string) bool {
	return GetBitByIndex(data, SetMapTable[name].Index)
}

func UpdateBitByIndex(data *int64, index int, value bool) {
	if value {
		*data |= (1 << index)
	} else {
		*data &= ^(1 << index)
	}
}

func UpdateBitByName(data *int64, name string, value bool) {
	UpdateBitByIndex(data, SetMapTable[name].Index, value)
}
