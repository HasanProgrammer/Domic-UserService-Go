package DomainUserEntity

import "time"

type StringifyUser struct {
	Id          string     `json:"Id"`
	FirstName   string     `json:"FirstName"`
	LastName    string     `json:"LastName"`
	Username    string     `json:"Username"`
	Password    string     `json:"Password"`
	Email       string     `json:"Email"`
	IsActive    bool       `json:"IsActive"`
	CreatedAt   time.Time  `json:"CreatedAt"`
	CreatedBy   string     `json:"CreatedBy"`
	CreatedRole string     `json:"CreatedRole"`
	UpdatedAt   *time.Time `json:"UpdatedAt"`
	UpdatedBy   string     `json:"UpdatedBy"`
	UpdatedRole *string    `json:"UpdatedRole"`
}
