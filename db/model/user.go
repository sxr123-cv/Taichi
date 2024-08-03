package model

import "Taichi/db"

type User struct {
	Id       int64  `json:"id" db:"id"`
	Name     string `json:"u_name" db:"u_name"`
	Password string `json:"u_password" db:"u_password"`
	Sex      string `json:"sex" db:"sex"`
}

func (User) FindOneUserById(id int64) (error, *User) {
	var u User
	err := db.Db.Get(&u, "SELECT u_name FROM t_user where id= ?", id)
	if err != nil {
		return err, nil
	}
	return nil, &u
}

func (User) FindAllUserById() (error, *[]User) {
	var u []User
	err := db.Db.Select(&u, "SELECT * FROM t_user ")
	if err != nil {
		return err, nil
	}
	return nil, &u
}

func (User) InsertUser(u *User) error {
	result, err := db.Db.NamedExec("insert into t_user (u_name,u_password,sex) values(:u_name,:u_password,:sex)", u)
	if err != nil {
		return err
	}
	u.Id, _ = result.LastInsertId()
	return nil
}
