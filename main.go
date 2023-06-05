package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      string
}

type UserResponse struct {
	Role string `json:"role"`
}

func TypeConverter[R any](data any) (*R, error) {
	var result R
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func main() {

	data, e := TypeConverter[UserResponse](User{})

	if e != nil {

		panic(e)
	}

	data.Role = "http://localhost:8080/users/:id/role"

	r, e := json.Marshal(data)
	if e != nil {

		panic(e)
	}

	fmt.Println(string(r))
}
