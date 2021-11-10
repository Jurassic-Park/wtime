package wtime

import (
	"errors"
	"fmt"
	"time"
)

// TimeFormat 时间格式
const TimeFormat = "2006-01-02 15:04:05"

// DateFormat 日期格式
const DateFormat = "2006-01-02"
const DateFormatChinese = "2006年01月02日"

// TimeFormatNoTag 没有符号的时间格式
const TimeFormatNoTag = "20060102150405"

// WTime 时间处理
type WTime struct {
	_time interface{}

	_timeObj time.Time
}

func NewWTime(_time interface{}) *WTime {
	return &WTime{_time: _time}
}

func (t *WTime) handle() (err error) {
	if !t._timeObj.IsZero() {
		return
	}

	if v, ok := t._time.(string); ok {
		// 时间字符串
		switch len(v) {
		case 19:
			t._timeObj, err = time.ParseInLocation(TimeFormat, v, time.Local)
		case 10:
			t._timeObj, err = time.ParseInLocation(DateFormat, v, time.Local)
		default:
			err = errors.New("时间字符串格式不正确[2021-11-10/2021-11-10 00:00:00]")
		}
	} else if v, ok := t._time.(int64); ok {
		switch len(fmt.Sprintf("%d", v)) {
		case 13:
			t._timeObj = time.Unix(0, v * int64(time.Millisecond))
		case 10:
			t._timeObj = time.Unix(v, 0)
		}
	} else if v, ok := t._time.(int); ok {
		switch len(fmt.Sprintf("%d", v)) {
		case 13:
			t._timeObj = time.Unix(0, int64(v) * int64(time.Millisecond))
		case 10:
			t._timeObj = time.Unix(int64(v), 0)
		}
	} else if v, ok := t._time.(time.Time); ok {
		t._timeObj = v
	} else {
		err = errors.New("时间格式不正确，请检查")
	}

	return
}

// ToTimeObj 变成time对象
func (t *WTime) ToTimeObj() (time.Time, error) {
	if err := t.handle(); err != nil {
		return time.Time{}, err
	}
	return t._timeObj, nil
}

// ToTimeStr 变成 2021-11-10 00:00:00
func (t *WTime) ToTimeStr() (string, error) {
	if err := t.handle(); err != nil {
		return "", err
	}
	return t._timeObj.Format(TimeFormat), nil
}

// ToDateStr 2021-11-10
func (t *WTime) ToDateStr() (string, error) {
	if err := t.handle(); err != nil {
		return "", err
	}
	return t._timeObj.Format(DateFormat), nil
}

// ToTimestamp 变成时间戳
func (t *WTime) ToTimestamp() (int64, error) {
	if err := t.handle(); err != nil {
		return 0, err
	}
	return t._timeObj.Unix(), nil
}

// ToMillisecond 变成毫秒
func (t *WTime) ToMillisecond() (int64, error) {
	if err := t.handle(); err != nil {
		return 0, err
	}
	return t._timeObj.UnixNano() / 1e6, nil
}
