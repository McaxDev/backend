package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func LoadUserBaseInfo(db *gorm.DB) *gorm.DB {
	return db.Select("name", "exp").Preload("Avatar", func(db *gorm.DB) *gorm.DB {
		return db.Select("filename")
	})
}

func CreateForeignKey(db *gorm.DB, fk ForeignKey) error {
	// 构建约束名称（保持一致性）
	constraintName := fmt.Sprintf("fk_%s_%s_%s",
		fk.Children, fk.ForeignKey, fk.Parent)

	// 1. 检查约束是否已存在
	var count int64
	query := `
        SELECT COUNT(*) 
        FROM information_schema.TABLE_CONSTRAINTS 
        WHERE 
            CONSTRAINT_SCHEMA = DATABASE() AND
            CONSTRAINT_NAME = ? AND
            CONSTRAINT_TYPE = 'FOREIGN KEY' AND
            TABLE_NAME = ?
    `
	if err := db.Raw(query, constraintName, fk.Children).Scan(&count).Error; err != nil {
		return fmt.Errorf("check constraint failed: %v", err)
	}

	// 约束已存在则直接返回
	if count > 0 {
		return nil
	}

	// 2. 构建并执行ALTER TABLE语句
	sql := fmt.Sprintf(
		"ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(id)",
		fk.Children, constraintName, fk.ForeignKey, fk.Parent)

	if fk.Action != "" {
		sql += " ON DELETE " + fk.Action
	}

	return db.Exec(sql).Error
}
