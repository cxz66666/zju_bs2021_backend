package user_service

import (
	"annotation/model/user"
	"annotation/utils/cache"
	"annotation/utils/db"
	"annotation/utils/logging"
	"regexp"
)

func QueryUserByEmail(email string) user.User {

	//email to uid
	uid:=cache.GetOrCreate(cache.GetKey(cache.EmailToId,email), func() interface{} {
		return Email2Id(email)
	}).(int)

	res:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,uid), func() interface{} {
		cacheUser,err:=GetUserById(uid)
		if err!=nil{
			logging.InfoF("a error occur in query user: %v\n",err)
			return &user.User{
				UserId: -1,
			}
		}
		return cacheUser
	}).(*user.User)

	return *res
}

func QueryUserByName(name string) user.User {

	//email to uid
	uid:=cache.GetOrCreate(cache.GetKey(cache.NameToId,name), func() interface{} {
		return Name2Id(name)
	}).(int)

	res:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,uid), func() interface{} {
		cacheUser,err:=GetUserById(uid)
		if err!=nil{
			logging.InfoF("a error occur in query user: %v\n",err)
			return &user.User{
				UserId: -1,
			}
		}
		return cacheUser
	}).(*user.User)

	return *res
}


// QueryUserById 根据Id查用户
func QueryUserById(uid int) user.User {
	res:=cache.GetOrCreate(cache.GetKey(cache.UserInfo,uid), func() interface{} {
		cacheUser,err:=GetUserById(uid)
		if err!=nil{
			return &user.User{
				UserId: -1,
			}
		}
		return cacheUser
	}).(*user.User)

	return *res
}

// CreateUser 创建用户
func CreateUser(userCreate *user.UserCreateReq) error {
	us:=user.User{
		UserName: userCreate.UserName,
		UserPhone: userCreate.UserPhone,
		UserEmail: userCreate.UserEmail,
		UserSecret: userCreate.UserSecret,
		UserType: user.Staff,
	}

	if err:=db.MysqlDB.Create(&us).Error;err!=nil{
		return err
	}
	return nil
}
// UpdateUser 更新用户数据
func UpdateUser(user user.User) error {
	if err:=db.MysqlDB.Select("*").Updates(&user).Error;err!=nil{
		return err
	}
	return nil
}

// CleanUserCache 删除缓存
func CleanUserCache(user user.User)  {
	cache.Remove(cache.GetKey(cache.UserInfo,user.UserId))
	cache.Remove(cache.GetKey(cache.EmailToId,user.UserEmail))
}




func Email2Id(email string) int {
	user:=user.User{}
	if err:=db.MysqlDB.First(&user,"user_email = ?",email).Error;err!=nil{
		return -1
	}
	return user.UserId
}

func Name2Id(name string) int  {
	user:=user.User{}
	if err:=db.MysqlDB.First(&user,"user_name = ?",name).Error;err!=nil{
		return -1
	}
	return user.UserId
}



func GetUserById(uid int) (*user.User,error)  {
	user:=user.User{}
	if err:=db.MysqlDB.First(&user,"user_id = ?",uid).Error;err!=nil{
		return nil,err
	}
	return &user,nil
}


const EmailRegex= "^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\\.[a-zA-Z0-9_-]+)+$"
const PhoneRegex= "^1(3\\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\\d|9[0-35-9])\\d{8}$"
// 长度 6-18 字母开头 只能包含数字 字母下划线
const PasswordRegex= "^[a-zA-Z]\\w{5,17}$"


//4-16位字母,数字,汉字,下划线
const NameRegex= "^([\u4e00-\u9fa5]{2,4})|([A-Za-z0-9_]{4,16})|([a-zA-Z0-9_\u4e00-\u9fa5]{3,16})$"


func ValidUser(user user.UserCreateReq) bool {
	if len(user.UserName)==0 || len(user.UserEmail)==0 || len(user.UserPhone)==0 || len(user.UserSecret)==0 {
		return false
	}

	if m, _ := regexp.MatchString(NameRegex, user.UserName); !m {
		logging.InfoF("用户名格式错误 name:%s\n",user.UserName)
		return false
	}

	if m, _ := regexp.MatchString(PhoneRegex, user.UserPhone); !m {
		logging.InfoF("手机号格式错误 phoneNum:%s\n",user.UserPhone)
		return false
	}

	if m, _ := regexp.MatchString(EmailRegex, user.UserEmail); !m {
		logging.InfoF("email格式错误 email:%s\n",user.UserEmail)
		return false
	}

	if m, _ := regexp.MatchString(PasswordRegex, user.UserSecret); !m {
		logging.InfoF("密码格式错误 secret:%s\n",user.UserSecret)
		return false
	}
	return true;

}