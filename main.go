package main

import (
	"encoding/json"
	"fmt"
	db "gomysql/library/database"
	"time"
)

type User struct {
	Id       int    `json:"id,omitempty,string"`
	Email    string `json:"email"`
	Passwd   string `json:"passwd,omitempty"`
	Name     string `json:"name"`
	Hashcode string `json:"hashcode"`
	Tpuser   string `json:"tpuser"`
	Status   string `json:"status"`
	Master   string `json:"master"`
	Dtupdate string `json:"dtupdate"`
	Dtcreate string `json:"dtcreate"`
}

type User_count struct {
	Status string `json:"status"`
	Count  int    `json:"count,string"`
}

func (u User_count) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprint(string(b))
}

func (u User) String() string {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return fmt.Sprint(string(b))
}

func (u User) Save() (*User, error) {
	if u.Passwd != "" && u.Email != "" {
		u.Hashcode = db.GetHash(u.Email, u.Passwd)
		u.Passwd = ""
	}
	t := time.Now().UTC()
	u.Dtupdate = t.Format("2006-01-02 15:04:05")

	if u.Id == 0 {
		response, err := db.Set[User]("user", u)
		if err != nil {
			fmt.Println("Error on insert!", err.Error())
			return nil, err
		}
		u.Id = int(response)
		return &u, nil
	}

	response, err := db.Update[User]("user", u, fmt.Sprintf("id = '%d'", u.Id))
	if err != nil {
		fmt.Println("Error on Update!", err.Error())
		return nil, err
	}
	u.Id = int(response)
	return &u, nil

}

func (u *User) GetAll() (*[]User, error) {
	response, err := db.GetAll[User]("user")

	if err != nil {
		fmt.Println("Error on GetAll!")
		return nil, err
	}

	return &response, nil
}

func (u *User) GetByID() (*User, error) {
	response, err := db.GetFirst[User]("user", &db.GetProps{Where: fmt.Sprintf("id='%d'", u.Id)})

	if err != nil {
		fmt.Println("Error on GetAll!")
		return nil, err
	}
	return response, nil
}

func (u *User) Delete() (bool, error) {
	response, err := db.Delete[User]("user", fmt.Sprintf("id='%d'", u.Id))
	if err != nil {
		fmt.Println("Error on GetAll!")
		return false, err
	}

	return response > 0, nil
}

func main() {
	u := &User{Status: "R", Id: 10015}
	u.Save()
}
