// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package types

type CreateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResp struct {
	Id int64 `json:"id"`
}

type GetUserReq struct {
	Id int64 `path:"id"`
}

type GetUserResp struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
