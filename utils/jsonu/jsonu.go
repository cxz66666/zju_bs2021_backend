package jsonu

import (
	"encoding/json"
	"log"
)

//	In this file I use encoding/json to Marshal or Unmarshal, because I don't ensure the specfical used places, maybe you can choose following
//  fastjson  https://github.com/valyala/fastjson (But it's useless when you try map or struct)

// 	If you don't care about `Intrusive code`, you also can try `msgp` https://github.com/tinylib/msgp, but you need generate a file for encode/decode (I like to use it)


//Marshal will panic if fail
func Marshal(v interface{} ) string  {
	marshal,err:=json.Marshal(v)
	if err!=nil{
		log.Panicln("Marshal failed: ",err)
	}
	return string(marshal)
}

//Unmarshal convert json string to map/struct object
func Unmarshal(s string, obj interface{})  {
	err:=json.Unmarshal([]byte(s),&obj)
	if err!=nil{
		log.Panicln("Unmarshal failed:",err)
	}
}


