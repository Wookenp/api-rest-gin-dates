package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FullName  string             `json:"full_name"`
	Email     string             `json:"email"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type Date struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	Title     string             `json:"title"`
	Details   string             `json:"details"`
	StartDate time.Time          `json:"start_date"`
	EndDate   time.Time          `json:"end_date"`
	User      User               `json:"user"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func NewUser(fullName string, email string) *User {
	return &User{
		Id:        primitive.NewObjectID(),
		FullName:  fullName,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewDate(title string, details string, startDate time.Time, endDate time.Time, user User) *Date {
	return &Date{
		Id:        primitive.NewObjectID(),
		Title:     title,
		Details:   details,
		StartDate: startDate,
		EndDate:   endDate,
		User:      user,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
