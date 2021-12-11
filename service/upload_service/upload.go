package upload_service

import (
	"annotation/model/project"
	"annotation/model/upload"
	"annotation/utils/conv"
	"annotation/utils/crypto"
	"annotation/utils/db"
	file2 "annotation/utils/file"
	"annotation/utils/setting"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var ImageTypes = []string{"image/png", "image/jpeg"}
var VideoTypes = []string{"video/mp4"}

func Contains(slice []string, s string) int {
	for index, value := range slice {
		if value == s {
			return index
		}
	}
	return -1
}
func SaveUploadedImage(pid int, userId int, file *multipart.FileHeader) error {
	//fmt.Println(file.Header)
	var content = file.Header.Get("Content-Type")
	if Contains(ImageTypes, content) == -1 {
		return errors.New("不是合法的文件类型，目前只支持png,jpg")
	}
	src, err := file.Open()
	if err != nil {
		return errors.New("文件打开错误")
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	crc32 := crypto.EncodeCrc32(string(data))
	//filepath是文件夹地址，dst是整个路径

	if setting.UploadSetting.Type == upload.Backend {
		filePath := filepath.FromSlash(setting.UploadSetting.BackendPath) + string(filepath.Separator) + conv.Int2Str(pid) + string(filepath.Separator)
		fileSuffix := path.Ext(file.Filename)                         //获取文件后缀
		filenameOnly := strings.TrimSuffix(file.Filename, fileSuffix) //获取文件名
		dst := filePath + filenameOnly + "-" + conv.Int2Str(int(crc32)) + fileSuffix

		err = file2.IsNotExistMkDir(filePath)
		if err != nil {
			return fmt.Errorf("file.IsNotExistMkDir src:%s, err: %v", src, err)
		}
		//如果存在则不做任何操作
		if file2.CheckExist(dst) {
			return nil
		}

		err = ioutil.WriteFile(dst, data, 0666)
		if err != nil {
			return err
		}

		image := upload.Image{
			Name:       file.Filename,
			Type:       setting.UploadSetting.Type,
			ProjectId:  pid,
			StorePath:  dst,
			Crc32Hash:  crc32,
			CreatorId:  userId,
			UploadTime: time.Now(),
		}

		return SaveImage(pid, &image)
	} else if setting.UploadSetting.Type == upload.OSS {
		filePath := setting.UploadSetting.OSSPath + "/" + conv.Int2Str(pid) + "/"
		fileSuffix := path.Ext(file.Filename)                         //获取文件后缀
		filenameOnly := strings.TrimSuffix(file.Filename, fileSuffix) //获取文件名
		dst := filePath + filenameOnly + "-" + conv.Int2Str(int(crc32)) + fileSuffix
		err, url := setting.UploadImage(dst, data)
		if err != nil {
			return fmt.Errorf("文件上传oss失败，请检查配置")
		}
		image := upload.Image{
			Name:       file.Filename,
			Type:       setting.UploadSetting.Type,
			ProjectId:  pid,
			StorePath:  url,
			Crc32Hash:  crc32,
			CreatorId:  userId,
			UploadTime: time.Now(),
		}
		return SaveImage(pid, &image)
	} else {
		return errors.New("未定义的upload type")
	}

}

func SaveImage(pid int, image *upload.Image) error {
	db.MysqlDB.Where("project_id = ? and crc32_hash = ? and name = ?", image.ProjectId, image.Crc32Hash, image.Name).Delete(&upload.Image{})

	if pid > 0 {
		if err := db.MysqlDB.Model(&project.Project{Id: pid}).Association("Images").Append(image); err != nil {
			return err
		}
		return nil
	} else {
		if err := db.MysqlDB.Create(image).Error; err != nil {
			return err
		}
		return nil
	}

}

// SaveUploadedVideo 用于转化视频文件为指定的图片，同时保存到相应的目录和db中
func SaveUploadedVideo(id int, userId int, file *multipart.FileHeader) (int, error) {
	var content = file.Header.Get("Content-Type")
	if Contains(VideoTypes, content) == -1 {
		return 0, errors.New("不是合法的类型，目前只支持mp4类型文件")
	}
	src, err := file.Open()
	defer src.Close()

	if err != nil {
		return 0, errors.New("文件打开错误")
	}
	data, err := ioutil.ReadAll(src)
	if err != nil {
		return 0, err
	}
	crc32 := crypto.EncodeCrc32(string(data))

	fileSuffix := path.Ext(file.Filename)                         //获取文件后缀
	filenameOnly := strings.TrimSuffix(file.Filename, fileSuffix) //获取文件名

	if err != nil {
		return 0, fmt.Errorf("file.IsNotExistMkDir src:%s, err: %v", src, err)
	}

	tempDirPath, err := ioutil.TempDir("", filenameOnly)
	defer os.Remove(tempDirPath)
	if err != nil {
		return 0, err
	}
	//fmt.Println("Temp dir created:", tempDirPath)
	tempFile, err := ioutil.TempFile(tempDirPath, file.Filename)
	defer os.Remove(tempFile.Name())

	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(tempFile.Name(), data, 0666)
	if err != nil {
		return 0, err
	}
	tempFile.Close()
	c := exec.Command(
		"ffmpeg", "-i", tempFile.Name(), "-q:v", "2", "-vf", "fps=1", fmt.Sprintf("%s%s-%d-%%d.jpg", tempDirPath+string(filepath.Separator), filenameOnly, crc32),
	)
	c.Stderr = os.Stderr
	if err = c.Run(); err != nil {
		return 0, err
	}

	var filePath string
	if setting.UploadSetting.Type == upload.Backend {
		filePath = filepath.FromSlash(setting.UploadSetting.BackendPath) + string(filepath.Separator) + conv.Int2Str(id) + string(filepath.Separator)
		err = file2.IsNotExistMkDir(filePath)
		if err != nil {
			return 0, fmt.Errorf("file.IsNotExistMkDir src:%s, err: %v", src, err)
		}
	} else if setting.UploadSetting.Type == upload.OSS {
		filePath = setting.UploadSetting.OSSPath + "/" + conv.Int2Str(id) + "/"
	}

	i := 1
	var images []upload.Image
	for ; ; i++ {
		picOnlyName := fmt.Sprintf("%s-%d-%d.jpg", filenameOnly, crc32, i)
		picName := fmt.Sprintf("%s%s", tempDirPath+string(filepath.Separator), picOnlyName)
		if !file2.CheckExist(picName) {
			break
		}
		originalFile, _ := os.Open(picName)

		if setting.UploadSetting.Type == upload.Backend {
			newFile, _ := os.Create(filePath + picOnlyName)
			_, _ = io.Copy(newFile, originalFile)
			_ = newFile.Sync()
			newFile.Close()

			images = append(images, upload.Image{
				Name:       picOnlyName,
				Type:       setting.UploadSetting.Type,
				ProjectId:  id,
				StorePath:  filePath + picOnlyName,
				Crc32Hash:  crc32,
				CreatorId:  userId,
				UploadTime: time.Now(),
			})
		} else if setting.UploadSetting.Type == upload.OSS {
			picData, _ := ioutil.ReadAll(originalFile)

			err, url := setting.UploadImage(filePath+picOnlyName, picData)
			if err != nil {
				fmt.Printf("文件%s上传oss失败，请检查配置\n", filePath+picOnlyName)
			}
			images = append(images, upload.Image{
				Name:       picOnlyName,
				Type:       setting.UploadSetting.Type,
				ProjectId:  id,
				StorePath:  url,
				Crc32Hash:  crc32,
				CreatorId:  userId,
				UploadTime: time.Now(),
			})
		}
		originalFile.Close()
	}

	err = db.MysqlDB.Transaction(func(tx *gorm.DB) error {

		for _, image := range images {
			tx.Where("project_id = ? and crc32_hash = ? and name = ?", image.ProjectId, image.Crc32Hash, image.Name).Delete(&upload.Image{})
			if id > 0 {
				if err = tx.Model(&project.Project{Id: id}).Association("Images").Append(&image); err != nil {
					return err
				}
			} else {
				if err = tx.Create(&image).Error; err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
		return nil
	})
	return i - 1, err
}
