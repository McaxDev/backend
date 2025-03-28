package utils

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type MySQLConfig struct {
	User     string `env:"MYSQL_USER" def:"root"`
	Password string `env:"MYSQL_PASSWORD"`
	Host     string `env:"MYSQL_HOST" def:"127.0.0.1"`
	Port     string `env:"MYSQL_PORT" def:"3306"`
	Name     string `env:"MYSQL_NAME"`
}

type SSLConfig struct {
	Enable bool   `env:"SSL_ENABLE" def:"false"`
	Cert   string `env:"CERT"`
	Key    string `env:"KEY"`
}

type RedisConfig struct {
	Host     string `env:"HOST" def:"127.0.0.1"`
	Port     string `env:"PORT" def:"6379"`
	Password string `env:"PASSWORD"`
	DB       int    `env:"DB"`
}

func LoadConfig(cfg any) error {
	val := reflect.ValueOf(cfg).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// 处理嵌套结构体
		if field.Kind() == reflect.Struct {
			if err := LoadConfig(field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		// 获取标签信息
		envKey := fieldType.Tag.Get("env")
		defVal := fieldType.Tag.Get("def")

		// 读取环境变量
		envVal := os.Getenv(envKey)
		if envVal == "" {
			envVal = defVal
		}

		// 跳过未设置且没有默认值的字段
		if envVal == "" {
			continue
		}

		// 设置字段值
		switch field.Kind() {
		case reflect.String:
			field.SetString(envVal)
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(envVal)
			if err != nil {
				return fmt.Errorf("invalid bool value for %s: %v", envKey, err)
			}
			field.SetBool(boolVal)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			intVal, err := strconv.ParseInt(envVal, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid int value for %s: %v", envKey, err)
			}
			field.SetInt(intVal)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			uintVal, err := strconv.ParseUint(envVal, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid uint value for %s: %v", envKey, err)
			}
			field.SetUint(uintVal)
		case reflect.Float32, reflect.Float64:
			floatVal, err := strconv.ParseFloat(envVal, 64)
			if err != nil {
				return fmt.Errorf("invalid float value for %s: %v", envKey, err)
			}
			field.SetFloat(floatVal)
		default:
			return fmt.Errorf("unsupported type %s for field %s", field.Kind(), fieldType.Name)
		}
	}
	return nil
}
