package utils

func GetBitByIndex(data int64, index uint) bool {
	return (data>>index)&1 == 1
}

func GetBitByID(data int64, id string) bool {
	return GetBitByIndex(data, SettingsTable[id])
}

func UpdateBitByIndex(data *int64, index uint, value bool) {
	if value {
		*data |= (1 << index)
	} else {
		*data &= ^(1 << index)
	}
}

func UpdateBitByID(data *int64, id string, value bool) {
	UpdateBitByIndex(data, SettingsTable[id], value)
}

type Setting struct {
	Index uint
	Name  string
}

// 如果改了Pub，要去privacy.go里面改
var SettingsTable = map[string]uint{
	"PubEmail":    1,
	"PubPhone":    2,
	"PubQQ":       3,
	"PubGameName": 4,
	"PubGuild":    5,
	"PubProps":    6,
	"PubComments": 7,
	"PubAlbums":   8,
	"PubCoin":     9,
	"PubGameData": 10,
	"PubDonation": 11,
	"UseMFA":      12,
}

type SettingsSliceElem struct {
	Name string
	ID   string
}

var SettingsSlice = []SettingsSliceElem{
	{"公开邮箱", "PubEmail"},
	{"公开手机号", "PubPhone"},
	{"公开QQ号", "PubQQ"},
	{"公开游戏名", "PubGameName"},
	{"公开我的公会", "PubGuild"}, 
	{"公开我的道具", "PubProps"}, 
	{"公开我的评论", "PubComments"},
	{"公开我的相册", "PubAlbums"},
	{"公开我的金额", "PubCoin"},
	{"公开游戏数据", "PubGameData"},
	{"公开捐赠数额", "PubDonation"},
	{"启用MFA", "UseMFA"},
}
