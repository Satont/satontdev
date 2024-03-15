package config

import (
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Application applicationConfig
	Site        siteConfig
	About       aboutConfig
	Career      careerConfig
}

type applicationConfig struct {
	Port int `yaml:"port"`
}

type aboutConfig struct {
	Title       string     `yaml:"title"`
	SubTitle    string     `yaml:"subTitle"`
	Description string     `yaml:"description"`
	BirthDate   string     `yaml:"birthDate"`
	Socials     []social   `yaml:"socials"`
	Languages   []language `yaml:"languages"`
	Skills      []string   `yaml:"skills"`
}

type siteConfig struct {
	Title string `yaml:"title"`
}

type careerConfig struct {
	PossibleRoles []string  `yaml:"possibleRoles"`
	Jobs          []job     `yaml:"jobs"`
	Projects      []project `yaml:"projects"`
}

var cfg *Config

func New(configsDir string) (*Config, error) {
	var app applicationConfig
	var site siteConfig
	var about aboutConfig
	var career careerConfig

	if err := cleanenv.ReadConfig(filepath.Join(configsDir, "application.yaml"), &app); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(filepath.Join(configsDir, "site.yaml"), &site); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(filepath.Join(configsDir, "about.yaml"), &about); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(filepath.Join(configsDir, "career.yaml"), &career); err != nil {
		return nil, err
	}

	cfg = &Config{
		Application: app,
		Site:        site,
		About:       about,
		Career:      career,
	}

	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}

type social struct {
	Name string `yaml:"name"`
	Href string `yaml:"href"`
}

type language struct {
	Name  string `yaml:"name"`
	Grade string `yaml:"grade"`
}

type job struct {
	Company   string  `yaml:"company"`
	Role      string  `yaml:"role"`
	DateRange string  `yaml:"dateRange"`
	Href      *string `yaml:"href"`
}

type project struct {
	Name         string   `yaml:"name"`
	Description  string   `yaml:"description"`
	Href         *string  `yaml:"href"`
	Technologies []string `yaml:"technologies"`
	Libraries    []string `yaml:"libraries"`
}
