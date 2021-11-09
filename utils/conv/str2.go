package conv
import(
	"strconv"
)
func str2int(){

}
func Int32_2str(a int32)string{
	string := strconv.Itoa(int(a))
	return string
}
func Int2Str(a int) string{
	return strconv.Itoa(a)
}

func Int64_2str(a int64)string{
	string := strconv.Itoa(int(a))
	return string
}

func Str2Int32(s string) int32{
	a, _ := strconv.ParseInt(s, 10, 32)
	return int32(a)
}

func Str2Int(s string) int{
	a,_ := strconv.Atoi(s)
	return a
}
