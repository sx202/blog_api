package comm

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


