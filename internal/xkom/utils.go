package xkom

import (
	"errors"
	"fmt"
	"strings"
	"xkomopener/internal/utils/helpers"

	"github.com/brianvoe/gofakeit"
)

func (s *Scraper) HandleEmail() {
	switch {
	case s.internal.BotConfig.Email.Catchall.Used:
		s.internal.AccountData.Email = fmt.Sprintf(
			"%s.%s%d@%s",
			strings.ToLower(s.internal.AccountData.FirstName),
			strings.ToLower(s.internal.AccountData.LastName),
			helpers.RandomInt(25, 999),
			s.internal.BotConfig.Email.Catchall.Catchall,
		)
	case s.internal.BotConfig.Email.EmailList.Used:
		panic(errors.New("unconfigured_email_type"))
	case s.internal.BotConfig.Email.DotTrick.Used:
		emailSplit := strings.Split(s.internal.BotConfig.Email.DotTrick.Email, "@")

		login := emailSplit[0]
		domain := emailSplit[1]

		s.internal.AccountData.Email = fmt.Sprintf(
			"%s+%s@%s",
			login,
			helpers.RandomString(helpers.RandomInt(8, 15), true, true, false),
			domain,
		)
	default:
		panic(errors.New("not_set_email"))
	}
}

func (s *Scraper) GenerateFakeData() {
	s.internal.AccountData.FirstName = gofakeit.FirstName()
	s.internal.AccountData.LastName = gofakeit.LastName()

	s.internal.AccountData.Password = gofakeit.Password(true, true, true, true, false, 25)

	s.HandleEmail()
}

func (a *AccountData) SaveToDb() {
	// TODO handle adding to db
}

func (s *Scraper) SetLoginCredentials(email, password string) {
	s.internal.AccountData.Email = email
	s.internal.AccountData.Password = password
}
