package models

type Workout struct{
	ID string `json:"id"`
	Type string `json:"type"`
	Distance string `json:"distance"`
	Duration string	`json:"duration"`
	Date string	`json:"date"`
	Note string	`json:"note,omitempty"`
}