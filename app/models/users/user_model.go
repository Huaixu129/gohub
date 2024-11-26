package users

import (
	"gohub/app/models"
    "gohub/pkg/database"
)

// User 用户模型
// 因为不希望将敏感信息输出给用户，所以这里 Email 、Phone 、Password 后面设置了 json:"-" ，这是在指示 JSON 解析器忽略字段
type User struct {
    models.BaseModel

    Name     string `json:"name,omitempty"`
    Email    string `json:"-"`
    Phone    string `json:"-"`
    Password string `json:"-"`

    models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
    database.DB.Create(&userModel)
}