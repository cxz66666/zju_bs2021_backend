package user_service

import (
	"annotation/model/user"
	"annotation/utils/db"
)

func QueryUsers() ([]user.User,int,error) {
	var ans []user.User
	var total int64
	if err:=db.MysqlDB.Model(&user.User{}).Count(&total).Order("user_id desc").Find(&ans).Error;err!=nil{
		return nil,0,err
	}
	return ans,int(total),nil
}

func DeleteUser(id int) error {
	if err:=db.MysqlDB.Delete(&user.User{},id).Error;err!=nil{
		return err
	}
	return nil
}

func ChangeRole(id int, newRole user.Role) error {
	if err:=db.MysqlDB.Model(&user.User{UserId: id}).Update("user_type",newRole).Error;err!=nil{
		return err
	}
	return nil
}

func GetNums() (int,int,error) {
	var normal int64
	var admin int64
	if err:=db.MysqlDB.Model(&user.User{}).Where("user_type = ?",user.Staff).Count(&normal).Error;err!=nil{
		return 0,0,err
	}
	if err:=db.MysqlDB.Model(&user.User{}).Where("user_type = ?",user.Admin).Count(&admin).Error;err!=nil{
		return 0,0,err
	}

	return int(normal),int(admin),nil
}