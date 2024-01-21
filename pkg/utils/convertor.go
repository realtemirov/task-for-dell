package utils

import (
	"bytes"
	"encoding/json"
	"strconv"
)

func StringToInt64(num string) (int64, error) {
	return strconv.ParseInt(num, 10, 64)
}

func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}
