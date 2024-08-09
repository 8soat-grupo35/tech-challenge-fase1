package customer

import (
	"errors"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/repository"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/ports/service"
)

type customerService struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerService(customerRepository repository.CustomerRepository) service.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (service *customerService) GetAll() ([]domain.Customer, error) {
	customers, err := service.customerRepository.GetAll()

	if err != nil {
		return []domain.Customer{}, errors.New("get customer from repository has failed")
	}

	return customers, nil
}

// Create implements ports.CustomerService.
func (service *customerService) Create(customer domain.Customer) (*domain.Customer, error) {
	customerSaved, err := service.customerRepository.Create(customer)

	if err != nil {
		return nil, errors.New("create customer on repository has failed")
	}

	return customerSaved, err
}

func (service *customerService) GetByCpf(cpf string) (*domain.Customer, error) {

	customer, err := service.customerRepository.GetOne(domain.Customer{
		CPF: cpf,
	})

	if err != nil {
		return nil, errors.New("error on obtain customer by CPF in repository")
	}

	return customer, err
}

// Update implements ports.CustomerService.
func (service *customerService) Update(customerId uint32, customer domain.Customer) (*domain.Customer, error) {

	customerAlreadySaved, err := service.customerRepository.GetOne(domain.Customer{
		Id: customerId,
	})

	if err != nil {
		return nil, errors.New("error on obtain customer to update in repository")
	}

	if customerAlreadySaved == nil {
		return nil, errors.New("customer not found to update")
	}

	customerUpdated, err := service.customerRepository.Update(customerId, customer)

	if err != nil {
		return nil, errors.New("updated customer on repository has failed")
	}

	return customerUpdated, err
}

// Delete implements ports.CustomerService.
func (service *customerService) Delete(customerId uint32) error {
	customerAlreadySaved, err := service.customerRepository.GetOne(domain.Customer{
		Id: customerId,
	})

	if err != nil {
		return errors.New("error on obtain customer to delete in repository")
	}

	if customerAlreadySaved == nil {
		return errors.New("customer not found to delete")
	}

	err = service.customerRepository.Delete(customerId)

	if err != nil {
		return errors.New("error on delete in repository")
	}

	return err
}
