package stringu

import (
	"fmt"
	"strconv"
)

//Tostring convert any kinds of params to string try its best!
func Tostring(i interface{}) string {
	if i==nil{
		return ""
	}
	var r string
	switch val := i.(type) {
	case string:
		r=val
	case int8:
		r=strconv.FormatInt(int64(val),10)
	case int:
		r=strconv.FormatInt(int64(val),10)
	case int32:
		r=strconv.FormatInt(int64(val),10)
	case int64:
		r=strconv.FormatInt(val,10)
	case uint8:
		r=strconv.FormatUint(uint64(val),10)
	case uint32:
		r=strconv.FormatUint(uint64(val),10)
	case uint:
		r=strconv.FormatUint(uint64(val),10)
	case uint64:
		r=strconv.FormatUint(val,10)
	case float32:
		r=fmt.Sprintf("%f",val)
	case float64:
		r=fmt.Sprintf("%f",val)
	}
	return r
}

func AddrString(s string)*string  {
	return &s
}

func ElemString(p *string)string  {
	if p!=nil{
		return *p
	}
	return ""
}