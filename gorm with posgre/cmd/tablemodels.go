package main

import "log"

const (
	tableNameForUser  = "information"
	tableNameForAdmin = "admin_information"
)
const (
	AdminRole = "admin"
)

type Information struct {
	Role      string
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func NewTableInformation(role string) *Information {
	return &Information{
		Role: role,
	}
}

func (Information) TableName() string {
	return tableNameForUser
}

func (i Information) TableNameByRole() string {
	if i.Role == AdminRole {
		return tableNameForAdmin
	}
	return tableNameForUser
}

func (i *Information) BeforeSave() error {
	log.Println("TRIGGER: Before save")
	return nil
}
