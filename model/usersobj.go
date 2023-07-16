package model

import "time"

type UserObj struct {
	Id               int       `json:"id" validate:"omitempty"`
	Firstname        string    `json:"firstname" validate:"omitempty"`
	Username         string    `json:"username" validate:"omitempty"`
	Lastname         string    `json:"lastname" validate:"omitempty"`
	Middlename       string    `json:"middlename" validate:"omitempty"`
	Phonenumber      string    `json:"phonenumber" validate:"omitempty"`
	DateCreated      time.Time `json:"datecreated" validate:"omitempty"`
	Roleid           int       `json:"roleid" validate:"omitempty"`
	Status           int       `json:"statusid" validate:"omitempty"`
	Email            string    `json:"email" validate:"omitempty"`
	Password         string    `json:"password" validate:"omitempty"`
	Authcode         string    `json:"authcode" validate:"omitempty"`
	GenderId         int       `json:"genderid" validate:"omitempty"`
	IdDetailID       int       `json:"iddetailid" validate:"omitempty"`
	Address          string    `json:"address" validate:"omitempty"`
	MarritalStatusID int       `json:"marritalstatusid" validate:"omitempty"`
	AccountId        int       `json:"accountid" validate:"omitempty"`
}

type TblUser struct {
	Id               int       `json:"Id" validate:"omitempty"`
	Firstname        string    `json:"Firstname" validate:"omitempty"`
	Username         string    `json:"Username" validate:"omitempty"`
	Lastname         string    `json:"Lastname" validate:"omitempty"`
	Middlename       string    `json:"Middlename" validate:"omitempty"`
	Phonenumber      string    `json:"Phonenumber" validate:"omitempty"`
	DateCreated      time.Time `json:"dateCreated" validate:"omitempty"`
	Roleid           int       `json:"Roleid" validate:"omitempty"`
	Status           int       `json:"Status" validate:"omitempty"`
	Email            string    `json:"Email" validate:"omitempty"`
	Password         string    `json:"Password" validate:"omitempty"`
	Authcode         string    `json:"Authcode" validate:"omitempty"`
	GenderId         int       `json:"GenderId" validate:"omitempty"`
	IdDetailID       int       `json:"IdDetailID" validate:"omitempty"`
	Address          string    `json:"Address" validate:"omitempty"`
	MarritalStatusID int       `json:"MarritalStatusID" validate:"omitempty"`
	AccountId        int       `json:"AccountId" validate:"omitempty"`
}
