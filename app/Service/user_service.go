package Service

import (
	"UserLogin/app/Model"
	"UserLogin/app/config"
	"fmt"
)

// user_service.go

func GetAllUsers() ([]Model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email, password FROM user")
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	defer rows.Close()

	var users []Model.User
	for rows.Next() {
		var u Model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		if err != nil {
			return nil, fmt.Errorf("扫描行失败: %v", err)
		}
		users = append(users, u)
	}
	return users, nil
}

// (Model.User, error)是返回值，同时返回结果和可能发生的错误。
func GetUserByID(id int) (Model.User, error) {
	var u Model.User
	err := config.DB.QueryRow("SELECT id, name, email FROM user WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return Model.User{}, fmt.Errorf("查询用户失败: %v", err)
	}
	return u, nil
}

func CreateUser(u Model.User) error {
	_, err := config.DB.Exec("INSERT INTO user (name, email, password) VALUES (?, ?, ?)",
		u.Name, u.Email, u.Password)
	if err != nil {
		return fmt.Errorf("创建用户失败: %v", err)
	}
	return nil
}

func UpdateUser(u Model.User) error {
	_, err := config.DB.Exec("UPDATE user SET name = ?, email = ? WHERE id = ?",
		u.Name, u.Email, u.ID)
	if err != nil {
		return fmt.Errorf("更新用户失败: %v", err)
	}
	return nil
}

func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("删除用户失败: %v", err)
	}
	return nil
}

func GetUserByEmail(email string) (Model.User, error) {
	var u Model.User
	err := config.DB.QueryRow("SELECT id, name, email, password FROM user WHERE email = ?", email).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return Model.User{}, fmt.Errorf("查询用户失败: %v", err)
	}
	fmt.Println("查询用户成功:", u)
	return u, nil
}
