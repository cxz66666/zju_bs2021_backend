package cache

import (
	"annotation/utils/logging"
	"annotation/utils/setting"
	"github.com/hashicorp/golang-lru"
	"log"
)

var cache *lru.ARCCache

//Setup is used to init a new cache with a size, you must use it after setting.Setup()
func Setup()  {
	var err error
	cache,err=lru.NewARC(setting.ServerSetting.CacheSize)
	if err!=nil{
		log.Panicf("cache init fail :%v\n",err)
	}
}



// ForceFlush wipe the entire cache table
func ForceFlush()  {
	logging.Info("force flush cache")
	cache.Purge()
}

// Remove purge a key from the cache
func Remove(key string) {
	cache.Remove(key)
}

// GetOrCreate accept a key and a function to return data (interface), if there isn't key-value in cache, it will call the
// function to get data and store it
func GetOrCreate(key string, f func()interface{}) interface{} {
	res,exist:=cache.Get(key)

	if !exist{
		newValue:=f()
		cache.Add(key,newValue)
		return newValue
	}
	return res
}


// Set put the key and value to the cache, and will force cover the old things
func Set(key string,value interface{})  {
	cache.Add(key,value);
}