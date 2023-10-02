package responses

import (
	"MusicServiceBackend/internal/role"
	"time"
)

type RespCr struct {
	Ok              bool   `json:"ok"`
	CreateComponent string `json:"create_component"`
}

type RespDel struct {
	Ok           bool `json:"ok"`
	Delcomponent int  `json:"del_component"`
	User         int  `json:"user"`
}

type RespAdd struct {
	Ok           bool `json:"ok"`
	Addcomponent int  `json:"add_component"`
	User         int  `json:"user"`
}

type RegisterResp struct {
	Ok bool `json:"ok"`
}

type LoginResp struct {
	ExpiresIn   time.Duration `json:"expires_in"`
	AccessToken string        `json:"access_token"`
	TokenType   string        `json:"token_type"`
	Username    string        `json:"username"`
	Email       string        `json:"email"`
	Role        role.Role     `json:"role"`
}
