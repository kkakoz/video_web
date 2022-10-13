package timex

import (
	"database/sql/driver"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

func Now() Time {
	return Time{time.Now()}
}

func Parse(layout, value string) (Time, error) {
	t, err := time.Parse(layout, value)
	if err != nil {
		return Time{}, err
	}
	return Time{t}, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	num, err := strconv.Atoi(string(data))
	if err != nil {
		return errors.New("timex.Time需要接受时间戳(秒)类型初始化")
	}
	t.Time = time.Unix(int64(num), 0)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", t.Time.Unix())), nil
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 方法来自mysql包复制,没有测试过其他的数据库
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case []byte:
		curTime, err := parseDateTime(v, time.UTC)
		if err != nil {
			t.Time = curTime
		}
		return err
	case string:
		curTime, err := parseDateTime([]byte(v), time.UTC)
		if err != nil {
			t.Time = curTime
		}
		return err
	}

	return fmt.Errorf("Can't convert %T to time.Time", value)
}
