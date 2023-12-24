package models



type LogInOut struct {
	Password	string	`json:"init_password"`
	Steps		uint	`json:"num_of_steps"`
}