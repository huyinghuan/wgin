package wgin


func getDefaultInt64(value string, defaultValue ...int64) int64 {
	d := int64(0)
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return d
	}
	return v
}
func getDefaultUint64(value string, defaultValue ...uint64) uint64 {
	d := uint64(0)
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return d
	}
	return v
}

func getDefaultInt(value string, defaultValue ...int) int {
	d := 0
	if len(defaultValue) > 0 {
		d = defaultValue[0]
	}
	if value == "" {
		return d
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return d
	}
	return v
}