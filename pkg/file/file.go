package file

import "os"

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 06444)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileCheck string) bool {
	if _, err := os.Stat(fileCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
