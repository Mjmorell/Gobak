package strutils

import "strconv"

func STR(data interface{}) string {
	switch data.(type) {

	// INT
	case int:
		return strconv.FormatInt(int64(data.(int)), 10)
	case int8:
		return strconv.FormatInt(int64(data.(int8)), 10)
	case int16:
		return strconv.FormatInt(int64(data.(int16)), 10)
	case int32:
		return strconv.FormatInt(int64(data.(int32)), 10)
	case int64:
		return strconv.FormatInt(data.(int64), 10)

	// UINT
	case uint:
		return strconv.FormatUint(uint64(data.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(data.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(data.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(data.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(data.(uint64), 10)

	// FLOAT
	case float32:
		return strconv.FormatFloat(float64(data.(float32)), 'f', 2, 64)
	case float64:
		return strconv.FormatFloat(data.(float64), 'f', 2, 64)

	// OTHERS
	case bool:
		return strconv.FormatBool(data.(bool))
	case string:
		return data.(string) // blubbering idiot

	// ELSE
	default:
		return "ERROR IN CONVERSION"
	}
}
