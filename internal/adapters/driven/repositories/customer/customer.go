package customer

import (
	"log"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"gorm.io/gorm"
)

type customerRepository struct {
	orm *gorm.DB
}

func NewRepository(orm *gorm.DB) repository.CustomerRepository {
	return &customerRepository{orm: orm}
}

func (c *customerRepository) GetAll() (customers []domain.Customer, err error) {
	result := c.orm.Find(&customers)

	if result.Error != nil {
		log.Println(result.Error)
		return customers, result.Error
	}

	return customers, err
}

func (c *customerRepository) GetOne(customerFilter domain.Customer) (customer *domain.Customer, err error) {
	result := c.orm.Where(customerFilter).First(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return customer, nil
}

func (c *customerRepository) Create(customer domain.Customer) (*domain.Customer, error) {
	result := c.orm.Create(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &customer, nil
}

func (c *customerRepository) Update(customerId uint32, customer domain.Customer) (*domain.Customer, error) {
	customerModel := domain.Customer{Id: customerId}
	result := c.orm.Model(&customerModel).Updates(&customer)

	if result.Error != nil {
		log.Println(result.Error)
		return nil, result.Error
	}

	return &customerModel, nil
}

func (c *customerRepository) Delete(customerId uint32) error {
	result := c.orm.Delete(&domain.Customer{}, customerId)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}
