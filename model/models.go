package model

type Payload struct {
	Title       string `yaml:"title" binding:"required"`
	Version     string `yaml:"version" binding:"required"`
	Maintainers []Maintainers `yaml:"maintainers" binding:"required,dive"`
	Company     string `yaml:"company" binding:"required"`
	Website     string `yaml:"website" binding:"required"`
	Source      string `yaml:"source" binding:"required"`
	License     string `yaml:"license" binding:"required"`
	Description string `yaml:"description" binding:"required"`
}

type Maintainers struct {
	Name  string `yaml:"name" binding:"required"`
	Email string `yaml:"email" binding:"required,email"`
}

