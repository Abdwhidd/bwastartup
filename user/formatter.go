package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

func FormatterUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.Id,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}

	return formatter
}
