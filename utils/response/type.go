package response


//Type 用于ResponseType 标识返回类型
type Type int


const TypeJSON Type=1
const TypeFile Type =2
const TypeFailed Type =3 //这个fail并不是status中的error，而是其他状态码的意思，比如40x，50x
const TypeRedirect Type=4
const TypeImage Type=5

//String 基本用于调试
func (t Type) String() string {
	switch t {
	case TypeJSON:
		return "JSON"
	case TypeFile:
		return "File"
	case TypeFailed:
		return "Failed"
	case TypeRedirect:
		return "Redirect"

	case TypeImage:
		return "Images"
	default:
		return "<unknown>"
	}
}

