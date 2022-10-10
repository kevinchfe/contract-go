package hash

import (
	"contract/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用bcrypt对密码加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

// BcryptCheck 对比密码和加密后的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// BcryptIsHashed 判断字符串是否加密
func BcryptIsHashed(str string) bool {
	return len(str) == 60
}
