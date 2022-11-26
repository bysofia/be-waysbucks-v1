package authdto

type LoginResponse struct {
	//ID       int    `gorm:"type: varchar(255)" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"fullName"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	//Password string `gorm:"type: varchar(255)" json:"password"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	//Role     string `json:"role" gorm:"type: varchar(255)"`
}

type RegisterResponse struct {
	Name  string `gorm:"type: varchar(255)" json:"fullName"`
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"fullName"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role  string `gorm:"type: varchar(50)"  json:"role"`
}
