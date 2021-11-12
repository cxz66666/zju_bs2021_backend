package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
)

// Package file provide some extra methods to help us handle with file

// GetSize return the multipart.File's size and error
func GetSize(f multipart.File) (int,error) {
	content,err:=ioutil.ReadAll(f)
	return len(content),err
}

// CheckExist accept a src path, and return true if file exist, return false if it doesn't
func CheckExist(src string) bool {
	_,err:=os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// CheckPermission accept a src path, and return true if we don't have the permission to this file
func CheckPermission(src string)bool  {
	_,err:=os.Stat(src)
	return os.IsPermission(err)
}

// MkDir will create all folder according to the file path, and perm is used 0777 which called os.ModePerm
func MkDir(src string) error {
	err:=os.MkdirAll(src,os.ModePerm)
	if err!=nil	{
		return err
	}
	return nil
}

// IsNotExistMkDir will create the whole folder to path if there doesn't exist the src path
func IsNotExistMkDir(src string) error {
	if exist:=CheckExist(src);exist==false{
		return MkDir(src)
	}
	return nil
}
// Open is simply use os.OpenFile to open the given file path
func Open(name string,flag int, perm os.FileMode) (*os.File,error) {
	f,err:=os.OpenFile(name,flag,perm)
	if err!=nil	{
		return nil,err
	}
	return f,nil
}

