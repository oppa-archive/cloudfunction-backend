package application

import (
	"github.com/oremi-co/api-template/models"
)

// Hello get the greeters details.
func Hello(Name string, repository models.GreetingRepository) (*models.Greeting, error) {
	g, err := repository.GetGreeter(Name)
	if err != nil {
		return nil, err
	}
	return g, err
}
