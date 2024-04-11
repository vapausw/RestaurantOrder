package utils

import (
	"RestaurantOrder/setting"
	"crypto/rand"
	"crypto/tls"
	"encoding/base32"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/smtp"
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

// SendEmail 使用SMTP发送邮件
func SendEmail(recipient, token string) error {
	from := setting.Conf.MyEmailConfig.Email
	password := setting.Conf.MyEmailConfig.Password

	// SMTP服务器地址和端口
	smtpHost := "smtp.163.com"
	smtpPort := "465"

	// 邮件消息体
	header := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Verify Your Email\r\n\r\n", from, recipient)
	body := fmt.Sprintf("Please use the following token to complete your registration: %s", token)
	message := []byte(header + body)

	// 创建TLS配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true, // 或者设置为false，并提供正确的证书链
		ServerName:         smtpHost,
	}

	// 连接到SMTP服务器
	conn, err := tls.Dial("tcp", smtpHost+":"+smtpPort, tlsconfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	// 创建smtp客户端
	client, err := smtp.NewClient(conn, smtpHost)
	if err != nil {
		return err
	}
	defer client.Close()

	// 认证
	auth := smtp.PlainAuth("", from, password, smtpHost)
	if err = client.Auth(auth); err != nil {
		return err
	}

	// 设置发送者和接收者
	if err = client.Mail(from); err != nil {
		return err
	}
	if err = client.Rcpt(recipient); err != nil {
		return err
	}

	// 发送邮件正文
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(message)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}
