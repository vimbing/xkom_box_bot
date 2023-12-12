package xkom

func (s *Scraper) GenerateAccount() (*AccountData, error) {
	s.GenerateFakeData()

	err := s.Register()

	if err != nil {
		return &AccountData{}, err
	}

	// err = s.Login()

	// if err != nil {
	// 	return &AccountData{}, err
	// }

	// err = s.RegisterToNewsletter()

	// if err != nil {
	// 	return &AccountData{}, err
	// }

	return s.internal.AccountData, nil
}
