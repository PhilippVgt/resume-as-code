package model

type CoverLetter struct {
	Name      string    `yaml:"name" json:"name" required:"true"`
	Date      *Date     `yaml:"date" json:"date" required:"true"`
	Title     string    `yaml:"title" json:"title"`
	Recipient Recipient `yaml:"recipient" json:"recipient" required:"true"`
	Address   Address   `yaml:"address" json:"address" required:"true"`
	Letter    Letter    `yaml:"letter" json:"letter" required:"true"`
}

type Recipient struct {
	Name string `yaml:"name" json:"name" required:"true"`
	Role string `yaml:"role" json:"role"`
}

type Letter struct {
	Salutation string `yaml:"salutation" json:"salutation" required:"true"`
	Body       string `yaml:"body" json:"body" required:"true"`
	Closing    string `yaml:"closing" json:"closing" required:"true"`
}
