package models

import (
	"gorm.io/gorm"
)

type NewCompany struct {
	CompanyName string `json:"companyname" validate:"required" `
	Location    string `json:"location" validate:"required"`
}
type Company struct {
	gorm.Model
	CompanyName string `validate:"required,unique" gorm:"unique;not null"`
	Location    string `json:"location"`
}

type NewJob struct {
	JobRole       string `json:"role" validate:"required"`
	Description   string `json:"description" validate:"required"`
	Min_Np        uint   `json:"minimum_notice_period" validate:"required"`
	Max_Np        uint   `json:"maximum_notice_period" validate:"required"`
	Budget        uint   `json:"budget" validate:"required"`
	JobLocation   []uint `json:"job_locations" validate:"required"`
	Technology    []uint `json:"job_technology" validate:"required"`
	WorKMode      []uint `json:"job_workmode" validate:"required"`
	MinExp        uint   `json:"minimum_experience" validate:"required"`
	MaxExp        uint   `json:"maximum_experience" validate:"required"`
	Qualification []uint `json:"job_qualification" validate:"required"`
	Shift         []uint `json:"job_shift" validate:"required"`
	JobType       []uint `json:"job_type" validate:"required"`
}
type Job struct {
	gorm.Model    `json:"-"`
	Company       Company         `json:"-" gorm:"ForeignKey:cid"`
	Cid           uint            `json:"cid"`
	JobRole       string          `json:"Role"`
	Description   string          `json:"description"`
	Min_Np        uint            `json:"minimum_notice_period"`
	Max_Np        uint            `json:"maximum_notice_period"`
	Budget        uint            `json:"budget"`
	JobLocation   []Location      `gorm:"many2many:job_location"`
	Technology    []Technology    `gorm:"many2many:job_technology"`
	WorKMode      []WorKMode      `gorm:"many2many:job_workmode"`
	MinExp        uint            `json:"minimum_experience"`
	MaxExp        uint            `json:"maximum_experience"`
	Qualification []Qualification `gorm:"many2many:job_qualification"`
	Shift         []Shift         `gorm:"many2many:job_shift"`
	JobType       []JobType       `gorm:"many2many:job_type"`
}
type Location struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type Technology struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type WorKMode struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type Qualification struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type Shift struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type JobType struct {
	gorm.Model
	Name string `gorm:"unique; column:name"`
}
type RequestJob struct {
	Name         string `json:"name" validate:"required"`
	JobId        uint   `json:"jobid" validate:"required"`
	JobRole      string `json:"role" validate:"required"`
	Description  string `json:"description" validate:"required"`
	NoticePeriod uint   `json:"noticPeriod" validate:"required"`
	Budget       uint   `json:"budget" validate:"required"`
	JobLocation  []uint `json:"job_locations" validate:"required"`
	Technology   []uint `json:"job_technology" validate:"required"`
	WorKMode     uint   `json:"job_workmode" validate:"required"`
	Exp          uint   `json:"experience" validate:"required"`
	// 	Qualification []uint `json:"job_qualification" validate:"required"`
	// 	Shift         []uint `json:"job_shift" validate:"required"`
	// 	JobType       []uint `json:"job_type" validate:"required"`
}
type NewJobResponse struct {
	ID uint
}

type NewRequestJob struct {
	Name string
}
