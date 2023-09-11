package test

import (
	"encoding/json"
)

// UserSession 用户会话数据
type UserSession struct {
	Id int64 `json:"id"`
}

func (u *UserSession) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *UserSession) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (*UserSession) New() *UserSession {
	return &UserSession{}
}
