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
	"UseMFA":      0,
	"PubCheckin":  1,
	"PubSetting":  2,
	"PubEmail":    3,
	"PubPhone":    4,
	"PubQQ":       5,
	"PubGameName": 6,
	"PubGuild":    7,
	"PubProps":    8,
	"PubComments": 9,
	"PubAlbums":   10,
	"PubCoin":     11,
	"PubGameData": 12,
}

var SettingsSlice = []string{
	"启用MFA",
	"公开签到记录",
	"公开个人设置",
	"公开邮箱",
	"公开手机号",
	"公开QQ号",
	"公开游戏名",
	"公开我的公会",
	"公开我的道具",
	"公开我的评论",
	"公开我的相册",
	"公开我的金额",
	"公开游戏数据",
}
