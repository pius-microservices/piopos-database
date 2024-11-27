package user

import "time"

type User struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	Role   Role   `gorm:"foreignKey:RoleID" json:"role"`
	RoleID string `gorm:"not null" json:"role_id" valid:"uuidv4"`

	Name      string    `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Name is required"`
	Username  string    `json:"username,omitempty" valid:"type(string)"`
	Email     string    `gorm:"not null;unique" json:"email" valid:"email, required~Email is required"`
	Password  string    `gorm:"not null" json:"password,omitempty" valid:"type(string), required~Password is required"`
	CreatedAt time.Time `json:"created_at"  valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}
