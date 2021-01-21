package dyanmic_params

import "time"

// Use any function started with Q as a faster way of calling
// methods. It suppress any error and returns the zero-value
// of the associated type. So, if errors are very important,
// do not use Q* methods and go with original methods without
// Q at the beginning of their name.
func QNewDynamicParams(vars ...interface{}) *DynamicParams {
	source := SrcNameInternal
	var ok bool
	if len(vars) != 0 {
		source, ok = vars[0].(string)
		if !ok {
			panic("cannot accepts a parameter which is not a string, as the source name")
		}
	}
	if len(vars) > 1 {
		NewDynamicParams(source, vars[1:]...)
	} else {
		NewDynamicParams(source)
	}
}

func (d *DynamicParams) QGetString(key string) string {
	v, err := d.GetAsString(key)
	if err != nil {
		return ""
	}
	return v
}

func (d *DynamicParams) QGetInt(key string) int {
	v, err := d.GetAsInt(key)
	if err != nil {
		return 0
	}
	return v
}

func (d *DynamicParams) QGetInt64(key string) int64 {
	v, err := d.GetAsInt64(key)
	if err != nil {
		return 0
	}
	return v
}
func (d *DynamicParams) QGetInt32(key string) int32 {
	v, err := d.GetAsInt32(key)
	if err != nil {
		return 0
	}
	return v
}
func (d *DynamicParams) QGetInt8(key string) int8 {
	v, err := d.GetAsInt8(key)
	if err != nil {
		return 0
	}
	return v
}
func (d *DynamicParams) QGetBool(key string) bool {
	v, err := d.GetAsBool(key)
	if err != nil {
		return false
	}
	return v
}
func (d *DynamicParams) QGetQuotedString(key string) string {
	v, err := d.GetAsQuotedString(key)
	if err != nil {
		return ""
	}
	return v
}
func (d *DynamicParams) QGetBytes(key string) []byte {
	v, err := d.GetAsBytes(key)
	if err != nil {
		return nil
	}
	return v
}

func (d *DynamicParams) QGetStringAsBool(key string) bool {
	v, err := d.GetStringAsBool(key)
	if err != nil {
		return false
	}
	return v
}
func (d *DynamicParams) QGetStringAsInt(key string) int {
	v, err := d.GetStringAsInt(key)
	if err != nil {
		return 0
	}
	return v
}
func (d *DynamicParams) QGetTimeDuration(key string) *time.Duration {
	v, err := d.GetAsTimeDuration(key)
	if err != nil {
		return nil
	}
	return v
}
func (d *DynamicParams) QGetStringAsTimeDuration(key string) *time.Duration {
	v, err := d.GetStringAsTimeDuration(key)
	if err != nil {
		return nil
	}
	return v
}
