package handlers

import (
	"fmt"
	"regexp"
)

const (
	patternUserName = `^[a-zA-Z0-9_-]{4,16}$`
)

type UserRegisterReq struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
	Email    string `json:"email" form:"email"`
}

func (c *UserRegisterReq) Validate() error {
	if ok, _ := regexp.Match(patternUserName, []byte(c.Username)); !ok {
		return fmt.Errorf("user_name仅支持4到16位，字母、数字、下划线、减号组成的字符")
	}
	return nil
}

type UserLoginReq struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (c *UserLoginReq) Validate() error {
	if ok, _ := regexp.Match(patternUserName, []byte(c.UserName)); !ok {
		return fmt.Errorf("user_name仅支持4到16位，字母、数字、下划线、减号组成的字符")
	}
	return nil
}
