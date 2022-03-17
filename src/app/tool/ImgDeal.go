package tool

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

//图片处理
func WriteImg(file *multipart.FileHeader, name string) (string, error) {
	filepoint, err := file.Open()
	if err != nil {
		return "", err
	}
	defer filepoint.Close()
	//创建目录
	if err := os.MkdirAll("./src/static/serverimg/", os.ModePerm); err != nil {
		return "", errors.New("创建目录失败")
	}
	//创建新文件进行存储
	time := strconv.FormatInt(time.Now().Unix(), 15)
	name = "/src/static/serverimg/" + time + name + ".jpg"
	newfile, err := os.Create(name)
	if err != nil {
		return "", err
	}
	defer newfile.Close()
	//把旧文件的内容放入新文件
	var context []byte = make([]byte, 1024)
	for {
		n, err := filepoint.Read(context)
		newfile.Write(context[:n])
		if err != nil {
			if err == io.EOF {
				return name, nil
			} else {
				return "", err
			}
		}
	}
}
