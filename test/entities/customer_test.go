package entities_test

import (
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomer_ValidData(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "Valid Name",
		Email: "valid@example.com",
		CPF:   "12345678901",
	}

	customer, err := entities.NewCustomer(customerDto)

	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, customerDto.Name, customer.Name)
	assert.Equal(t, customerDto.Email, customer.Email)
	assert.Equal(t, customerDto.CPF, customer.CPF)
}

func TestNewCustomer_InvalidName(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "abc",
		Email: "valid@example.com",
		CPF:   "12345678901",
	}

	customer, err := entities.NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestNewCustomer_InvalidEmail(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "Valid Name",
		Email: "invalid-email",
		CPF:   "12345678901",
	}

	customer, err := entities.NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestNewCustomer_InvalidCPF(t *testing.T) {
	customerDto := dto.CustomerDto{
		Name:  "Valid Name",
		Email: "valid@example.com",
		CPF:   "123",
	}

	customer, err := entities.NewCustomer(customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestCustomer_Validate_ValidCustomer(t *testing.T) {
	customer := entities.Customer{
		Name:  "Valid Name",
		Email: "valid@example.com",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.NoError(t, err)
}

func TestCustomer_Validate_InvalidName(t *testing.T) {
	customer := entities.Customer{
		Name:  "abc",
		Email: "valid@example.com",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.Error(t, err)
}

func TestCustomer_Validate_InvalidEmail(t *testing.T) {
	customer := entities.Customer{
		Name:  "Valid Name",
		Email: "invalid-email",
		CPF:   "12345678901",
	}

	err := customer.Validate()

	assert.Error(t, err)
}

func TestCustomer_Validate_InvalidCPF(t *testing.T) {
	customer := entities.Customer{
		Name:  "Valid Name",
		Email: "valid@example.com",
		CPF:   "123",
	}

	err := customer.Validate()

	assert.Error(t, err)
}
