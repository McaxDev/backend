package main

import (
	"github.com/McaxDev/backend/dbs"
	"github.com/McaxDev/backend/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Anchor struct {
	ID         string   `json:"id"`
	SubAnchors []Anchor `json:"subAnchors"`
}

func Get(c *gin.Context, req struct {
	Path string `form:"path"`
}) {

	var data dbs.Wiki
	if err := DB.Where("path = ?", req.Path).First(
		&data,
	).Error; err == gorm.ErrRecordNotFound {
		c.JSON(404, utils.Resp("当前路径不存在", nil, nil))
		return
	} else if err != nil {
		c.JSON(500, utils.Resp("wiki获取失败", err, nil))
		return
	}

	c.JSON(200, utils.Resp("", nil, gin.H{
		"anchors": []Anchor{
			{"测试锚点1", []Anchor{}},
			{"测试锚点2", []Anchor{
				{"测试锚点3", []Anchor{}},
				{"测试锚点4", []Anchor{}},
			}},
			{"测试锚点5", []Anchor{
				{"测试锚点6", []Anchor{}},
			}},
		},
		"wiki": &data,
	}))
}
