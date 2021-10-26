package reflectu

import (
	"os"
	"reflect"
)

// SetStructByReflect set the field value of struct to os.env!
func SetStructByReflect(valuePtr reflect.Value, envKey string,fieldName string) bool  {
	str:=os.Getenv(envKey)
	if len(str)==0{
		return false
	}
	//find the field of name
	field:=valuePtr.Elem().FieldByName(fieldName)
	// check whether zero
	if field.IsZero(){
		return false
	}

	field.Set(reflect.ValueOf(str))
	return true
}
