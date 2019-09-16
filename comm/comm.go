package comm

type Bloguser struct {
	Id       int        `json:"Id"`
	Username string		`json:"Username"`
	Password string		`json:"Password"`
	Email    string		`json:"Email"`
	Roles    int		`json:"Roles"`
}

type Question struct {
	Id              int     `json:"id"`
	Question        string  `json:"Question"`
	OptionA         string	`json:"Option_A"`
	OptionB         string	`json:"Option_B"`
	OptionC         string	`json:"Option_C"`
	OptionD         string	`json:"Option_D"`
	OptionE         string	`json:"Option_E"`
	OptionF         string	`json:"Option_F"`
	OptionG         string	`json:"Option_G"`
	CorrectAnswer1 	string	`json:"Correct_Answer_1"`
	CorrectAnswer2 	string	`json:"Correct_Answer_2"`
	CorrectAnswer3 	string	`json:"Correct_Answer_3"`
	CorrectAnswer4 	string	`json:"Correct_Answer_4"`
	CorrectAnswer5 	string	`json:"Correct_Answer_5"`
	CorrectAnswer6 	string	`json:"Correct_Answer_6"`
	CorrectAnswer7 	string	`json:"Correct_Answer_7"`
}

type ChangeQuestion struct {
	key   string
	value string
}

type User struct {
	Id         int     `json:"id"`
	UserName   string  `json:"username"`
	PassWord   string  `json:"password"`
	Email      string  `json:"email"`
	Roles       string  `json:"roles"`
}

