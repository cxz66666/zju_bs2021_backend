package numberu

import (
	"reflect"
	"strconv"
)

func BoolToInt(i bool) int {
	if i {
		return 1
	}
	return 0
}

//ToInt64 will try its best to convert others kind to int64
func ToInt64(i interface{}) int64 {
	if i==nil{
		return 0
	}
	var r int64
	var v=reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		r,_=strconv.ParseInt(i.(string),10,64)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		r = v.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		r = int64(v.Uint())
	case reflect.Float32,reflect.Float64:
		r=int64(v.Float())
	}
	return r
}

//ToInt will try its best to convert others kind to int, and maybe when int64 convert to int32, will cause truncation
func ToInt(i interface{}) int {
	return int(ToInt64(i))
}

//ToFloat64 will try its best to convert others kind to float64
func ToFloat64(i interface{}) float64 {
	if i==nil{
		return 0
	}
	var r float64
	var v=reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		r,_=strconv.ParseFloat(i.(string),64)
	case reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		r=float64(v.Int())
	case reflect.Uint8,reflect.Uint16,reflect.Uint32,reflect.Uint64:
		r=float64(v.Uint())
	case reflect.Float32,reflect.Float64:
		r=v.Float()
	}
	return r
}


func AddrInt(i int) *int {
	return &i
}

func AddrInt64(i int64)*int64  {
	return &i
}

func AddrUint(u uint) *uint {
	return &u
}
func AddrFloat64(f float64) *float64 {
	return &f
}

func ElemInt(p *int) int {
	if p != nil {
		return *p
	}
	return 0
}

func ElemInt64(p *int64) int64 {
	if p != nil {
		return *p
	}
	return 0
}

func ElemFloat64(p *float64) float64 {
	if p != nil {
		return *p
	}
	return 0
}
