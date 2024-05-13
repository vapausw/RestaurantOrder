package util

import (
	"RestaurantOrder/model"
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	mr "math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// GenValidateCode 生成一个安全的随机令牌，长度为 16 位，由大小写字母和数字组成
func GenValidateCode() string {
	// 定义令牌的字节长度，base32 编码每5个比特表示一个字符，因此对于8个字符，需要5*16/8=10个字节
	tokenLength := 10
	b := make([]byte, tokenLength)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalf("Failed to generate random token: %v", err)
	}
	// 使用 base32 编码生成令牌，然后取前16位作为结果
	// 只取前16个字符
	token := base32.StdEncoding.EncodeToString(b)[:16]
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

// ValidateEmail 使用正则表达式检验邮箱格式
func ValidateEmail(email string) bool {
	/*
		^[a-z0-9._%+-]+：邮箱用户名部分，允许小写字母、数字、点、下划线、百分号、加号和减号。
		@：必须包含一个@符号。
		[a-z0-9.-]+：域名部分，允许小写字母、数字、点和减号。
		\.[a-z]{2,4}$：点后面跟2到4个字母的顶级域名。
	*/
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}

// GetRandomExpirationInSeconds 生成一个随机的过期时间
// minSeconds 和 maxSeconds 定义生成随机秒数的范围
func GetRandomExpirationInSeconds(minSeconds, maxSeconds int) time.Duration {
	randGen := mr.New(mr.NewSource(time.Now().UnixNano())) //
	// 保证最小值不大于最大值
	if minSeconds > maxSeconds {
		minSeconds, maxSeconds = maxSeconds, minSeconds
	}
	// 生成[minSeconds, maxSeconds]范围内的随机时间
	return time.Duration(randGen.Intn(maxSeconds-minSeconds+1)+minSeconds) * time.Second
}

// GenerateRandomNickname 生成随机昵称
func GenerateRandomNickname() string {
	randGen := mr.New(mr.NewSource(time.Now().UnixNano())) //
	// 定义可能的字母组成
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	nicknameLength := 8 // 设定昵称的长度

	// 生成随机字母部分
	nickname := make([]rune, nicknameLength)
	for i := range nickname {
		nickname[i] = letters[randGen.Intn(len(letters))]
	}

	// 添加一些随机数字
	numbers := fmt.Sprintf("%04d", randGen.Intn(10000)) // 生成一个0到9999之间的数字

	return string(nickname) + numbers
}

// 序列化与反序列化

// Serialize 将任何 变量序列化为 JSON 字符串
func Serialize(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Deserialize 将 JSON 字符串反序列化为指定的结构体
func Deserialize(jsonStr string, target interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		return err
	}
	return nil
}

// SerializeShops 将 model.ShopList 数组序列化为字符串数组
func SerializeShops(data []model.ShopList) ([]string, error) {
	var serializedStrings []string
	for _, item := range data {
		bytes, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		serializedStrings = append(serializedStrings, string(bytes))
	}
	return serializedStrings, nil
}

// DeserializeShops 将字符串数组反序列化为 model.ShopList 数组
func DeserializeShops(jsonStrings []string, targetType *[]model.ShopList) error {
	for _, jsonString := range jsonStrings {
		var item model.ShopList
		err := json.Unmarshal([]byte(jsonString), &item)
		if err != nil {
			return err
		}
		*targetType = append(*targetType, item)
	}
	return nil
}

func SerializeMenus(menus []model.MenuList) ([]string, error) {
	var serializedStrings []string
	for _, item := range menus {
		bytes, err := json.Marshal(item)
		if err != nil {
			return nil, err
		}
		serializedStrings = append(serializedStrings, string(bytes))
	}
	return serializedStrings, nil
}

func DeserializeShopMenus(menus []string, i *[]model.MenuList) error {
	for _, menu := range menus {
		var item model.MenuList
		err := json.Unmarshal([]byte(menu), &item)
		if err != nil {
			return err
		}
		*i = append(*i, item)
	}
	return nil
}

// StructToMap 将结构体转换为 map[string]interface{}
func StructToMap(item interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	v := reflect.ValueOf(item)

	// 如果是指针类型，需要取出指针指向的元素
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 如果不是结构体类型，返回错误
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input must be a struct or a pointer to a struct")
	}

	// 遍历结构体的字段
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		// 获取字段的 db tag，如果没有则使用字段名
		tag := field.Tag.Get("db")
		if tag == "" {
			tag = field.Name
		}

		// 直接使用原始类型
		switch fieldValue.Kind() {
		case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result[tag] = fieldValue.Interface()
		default:
			return nil, fmt.Errorf("unsupported field type %s for field %s", fieldValue.Type(), field.Name)
		}
	}
	return result, nil
}

// MapToStruct 将 map 转换为结构体
func MapToStruct(data map[string]string, result interface{}) error {
	v := reflect.ValueOf(result)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return fmt.Errorf("result must be a non-nil pointer to a struct")
	}
	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("result must point to a struct")
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			tag = field.Name
		}

		valueStr, ok := data[tag]
		if !ok {
			continue // Field not found in map, skip
		}

		fieldValue := v.Field(i)
		if !fieldValue.CanSet() {
			continue // Ignore unexported fields
		}

		switch fieldValue.Kind() {
		case reflect.String:
			fieldValue.SetString(valueStr)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
				fieldValue.SetInt(value)
			} else {
				return fmt.Errorf("could not parse integer for field %s: %v", field.Name, err)
			}
		default:
			return fmt.Errorf("unsupported field type %s for field %s", fieldValue.Type(), field.Name)
		}
	}
	return nil
}
