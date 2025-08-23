package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// HTime 自定义时间类型，嵌入time.Time以继承其方法
type HTime struct {
	time.Time
}

// 时间格式化模板
var (
	formatTime = "2006-01-02 15:04:05"
)

// MarshalJSON 实现JSON序列化接口
func (t HTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// UnmarshalJSON 实现JSON反序列化接口
func (t *HTime) UnmarshalJSON(data []byte) error {
	// 去除JSON字符串的引号
	str := string(data)
	if len(str) >= 2 && str[0] == '"' && str[len(str)-1] == '"' {
		str = str[1 : len(str)-1]
	}

	now, err := time.ParseInLocation(formatTime, str, time.Local)
	if err != nil {
		return err
	}
	*t = HTime{Time: now}
	return nil
}

// Value 实现数据库写入接口
func (t HTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 实现数据库读取接口
func (t *HTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = HTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to HTime", v)
}
