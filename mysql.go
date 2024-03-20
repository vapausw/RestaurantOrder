package main

// 对于mysql数据库的一些操作
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// DSN（Data Source Name）格式：用户名:密码@协议(地址:端口)/数据库名
// dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8"
const (
	dsn string = "wby:YJjt2245!@tcp(192.168.30.128:3306)/RestaurantOrder?charset=utf8"
)

type MysqlData struct {
	db *sql.DB
}

func (m *MysqlData) Connect() {
	m.db, _ = sql.Open("mysql", dsn)
	// 确保连接成功
	err := m.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the database successfully!")
}

func (m *MysqlData) Close() {
	m.db.Close()
}

func (m *MysqlData) insertUser(user Customer) {
	query := `INSERT INTO Users (CustomerID, Name, Email, Password, IsVIP, ResetToken, TokenExpiry) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.db.Exec(query, user.CustomerID, user.Name, user.Email, user.Password, user.IsVIP, user.ResetToken, user.TokenExpiry)
	if err != nil {
		log.Fatalf("Failed to insert user: %v", err)
	}
	fmt.Println("Inserted user successfully")
}

func (m *MysqlData) queryUsers() (users []Customer) {
	query := `SELECT CustomerID, Name, Email, Password, IsVIP, ResetToken, TokenExpiry FROM Users`
	rows, err := m.db.Query(query)
	if err != nil {
		log.Fatalf("Failed to query users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user Customer
		var tokenExpiry []byte // 使用[]byte类型接收TokenExpiry
		err := rows.Scan(&user.CustomerID, &user.Name, &user.Email, &user.Password, &user.IsVIP, &user.ResetToken, &tokenExpiry)
		if err != nil {
			log.Fatal(err)
		}

		// 手动解析TokenExpiry
		if tokenExpiry != nil {
			user.TokenExpiry, err = time.Parse("2006-01-02 15:04:05", string(tokenExpiry))
			if err != nil {
				log.Fatal("Failed to parse TokenExpiry: ", err)
			}
		}

		users = append(users, user)
	}
	return
}

func (m *MysqlData) deleteUser(userID string) {
	query := `DELETE FROM Users WHERE CustomerID = ?`
	_, err := m.db.Exec(query, userID)
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}
	fmt.Println("Deleted user successfully")
}

func (m *MysqlData) updateUserStatus(user1, user2 Customer) { // 采用删除加插入的方式
	m.deleteUser(user1.CustomerID)
	m.insertUser(user2)
	fmt.Println("Updated user status successfully")
}
