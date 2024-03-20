package main

import (
	"log"
	"testing"
	"time"
)

func setupTestDB() *MysqlData {
	m := MysqlData{}
	m.Connect()
	// 清理测试数据
	_, err := m.db.Exec("DELETE FROM Users")
	if err != nil {
		log.Fatalf("Failed to clean up test data: %v", err)
	}
	return &m
}

func TestInsertAndQueryUser(t *testing.T) {
	m := setupTestDB()
	defer m.Close()

	// 插入测试用户
	testUser := Customer{
		CustomerID:  "123",
		Name:        "Test Customer",
		Email:       "test@example.com",
		Password:    "hashedpassword",
		IsVIP:       false,
		ResetToken:  "re1Token",
		TokenExpiry: time.Now(),
	}
	m.insertUser(testUser)

	// 查询测试用户
	users := m.queryUsers()
	if len(users) != 1 || users[0].CustomerID != testUser.CustomerID {
		t.Errorf("Expected 1 user with CustomerID %v, got %v", testUser.CustomerID, len(users))
	}

	// 删除测试数据
	m.deleteUser(testUser.CustomerID)
	users = m.queryUsers()
	if len(users) != 0 {
		t.Errorf("Expected 0 user, got %v", len(users))
	}
	// 添加并更新测试数据
	m.insertUser(testUser)
	testUser1 := Customer{
		CustomerID:  "123456",
		Name:        "Test Customer",
		Email:       "test1@example.com",
		Password:    "hashedpassword",
		IsVIP:       true,
		ResetToken:  "re2Token",
		TokenExpiry: time.Now(),
	}
	m.updateUserStatus(testUser, testUser1)
	users = m.queryUsers()
	if len(users) != 1 && users[0].CustomerID != testUser1.CustomerID {
		t.Errorf("Expected 1 user with CustomerID %v, got %v", testUser1.CustomerID, len(users))
	}
	// 清理测试数据
	_, err := m.db.Exec("DELETE FROM Users WHERE CustomerID = ?", testUser.CustomerID)
	if err != nil {
		t.Fatalf("Failed to clean up test data: %v", err)
	}
}
