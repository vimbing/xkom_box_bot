package config

type BotConfig struct {
	Email struct {
		Catchall struct {
			Used     bool   `json:"used"`
			Catchall string `json:"catchall"`
		} `json:"catchall"`
		DotTrick struct {
			Used  bool   `json:"used"`
			Email string `json:"email"`
		} `json:"dotTrick"`
		EmailList struct {
			Used     bool   `json:"used"`
			ListFile string `json:"listFile"`
		} `json:"emailList"`
	} `json:"email"`
}
