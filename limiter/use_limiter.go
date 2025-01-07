package limiter

import (
	"errors"
	"fmt"
	"time"

	"github.com/McaxDev/backend/dbs"
)

func UseLimiter(user, action string) error {

	query := DB.Where(
		"user = ? AND action = ?", user, action,
	)
	var histories []dbs.LimiterRecord
	if err := query.Find(&histories).Error; err != nil {
		return errors.New("数据库查询失败")
	}

	rules := Limiter[action]
	if rules == nil {
		return errors.New("不存在这个规则")
	}

	for index := range histories {
		histories[index].Valid = false
	}

	now := time.Now()
	var result string
	for _, rule := range rules {
		var times uint
		for index, history := range histories {
			timeLeft := history.Time.Sub(now.Add(-rule.Duration))
			if timeLeft > 0 {
				histories[index].Valid = true
				times++
			}
		}
		if times > rule.Count {
			result = fmt.Sprintf(
				"limit %d times in %s, ",
				rule.Count, rule.Duration.String(),
			)
		}
	}

	for _, history := range histories {
		if !history.Valid {
			if err := DB.Delete(&history).Error; err != nil {
				return err
			}
		}
	}

	if result != "" {
		return errors.New(result)
	}

	return nil
}
