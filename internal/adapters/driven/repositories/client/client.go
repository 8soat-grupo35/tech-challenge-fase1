package client

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"gorm.io/gorm"
)

type clientRepository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) repository.ClientRepository {
	return &clientRepository{orm: orm}
}

func (c *clientRepository) GetAll() (clients []domain.Client, err error) {
	result := c.orm.Find(&clients)

	if result.Error != nil {
		log.Println(result.Error)
		return clients, result.Error
	}

	return clients, err
}

func (c *clientRepository) GetOne(clientFilter domain.Client) (client *domain.Client, err error) {
	result := c.orm.Where(clientFilter).First(&client)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return client, nil
}

func (c *clientRepository) Create(client domain.Client) (*domain.Client, error) {
	result := c.orm.Create(&client)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &client, nil
}

func (c *clientRepository) Update(clientId uint32, client domain.Client) (*domain.Client, error) {
	clientModel := domain.Client{Id: clientId}
	result := c.orm.Model(&clientModel).Updates(&client)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &clientModel, nil
}

func (c *clientRepository) Delete(clientId uint32) error {
	result := c.orm.Delete(&domain.Client{}, clientId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
