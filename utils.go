package main

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// GenerateSecureToken 生成一个安全的随机令牌，长度为 8 位，由大小写字母和数字组成
func GenerateSecureToken() string {
	// 定义令牌的字节长度，base32 编码每5个比特表示一个字符，因此对于8个字符，我们需要5*8/8=5个字节
	tokenLength := 5
	b := make([]byte, tokenLength)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate random token: %v", err)
	}
	// 使用 base32 编码生成令牌，然后取前8位作为结果
	// 注意：这里使用了 base32，因为它比 base64 更容易产生人类可读的字符（大小写字母和数字）
	// 但由于输出会更长，我们只取前8个字符
	token := base32.StdEncoding.EncodeToString(b)[:8]
	return token
}

// HashPassword 生成密码的哈希值
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash 检查密码哈希值是否匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CreateID 生成一个唯一的 ID
func CreateID() string {
	return uuid.New().String()
}
