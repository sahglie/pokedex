package config

type AppConfig struct {
	NextPage int
	PrevPage int
	Prompt   string
}

func NewAppConfig() AppConfig {
	return AppConfig{
		NextPage: 1,
		PrevPage: 0,
		Prompt:   "Pokedex > ",
	}
}

func (c *AppConfig) AdvancePager() {
	c.NextPage++
	c.PrevPage++
}

func (c *AppConfig) RewindPager() {
	if c.PrevPage == 0 {
		return
	}

	c.NextPage--
	c.PrevPage--
}
