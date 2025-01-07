module github.com/McaxDev/backend/limiter

go 1.23.3

replace github.com/McaxDev/backend/dbs => ../dbs

require github.com/McaxDev/backend/dbs v0.0.0-00010101000000-000000000000

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.21.0 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	gorm.io/gorm v1.25.12 // indirect
)
