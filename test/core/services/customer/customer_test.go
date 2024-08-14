package customer

import (
	"errors"
	"testing"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/adapters/driver/dto"
	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/domain"

	"github.com/8soat-grupo35/tech-challenge-fase1/internal/core/services/customer"
	mock_repository "github.com/8soat-grupo35/tech-challenge-fase1/test/core/ports/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAll_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetAll().Return([]domain.Customer{{ID: 1, Name: "John Doe"}}, nil)

	service := customer.NewCustomerService(mockRepo)
	customers, err := service.GetAll()

	assert.NoError(t, err)
	assert.Len(t, customers, 1)
	assert.Equal(t, "John Doe", customers[0].Name)
}

func TestGetAll_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetAll().Return(nil, errors.New("repository error"))

	service := customer.NewCustomerService(mockRepo)
	customers, err := service.GetAll()

	assert.Error(t, err)
	assert.Empty(t, customers)
}

func TestCreate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}
	newCustomer := &domain.Customer{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().Create(*newCustomer).Return(newCustomer, nil)

	service := customer.NewCustomerService(mockRepo)
	createdCustomer, err := service.Create(customerDto)

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", createdCustomer.Name)
}

func TestCreate_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "JD", Email: "john@example.com", CPF: "12345678901"}

	service := customer.NewCustomerService(nil)
	createdCustomer, err := service.Create(customerDto)

	assert.Error(t, err)
	assert.Nil(t, createdCustomer)
}

func TestCreate_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}
	newCustomer := &domain.Customer{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().Create(*newCustomer).Return(nil, errors.New("repository error"))

	service := customer.NewCustomerService(mockRepo)
	createdCustomer, err := service.Create(customerDto)

	assert.Error(t, err)
	assert.Nil(t, createdCustomer)
}

func TestGetByCpf_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{CPF: "12345678901"}).Return(&domain.Customer{Name: "John Doe"}, nil)

	service := customer.NewCustomerService(mockRepo)
	customer, err := service.GetByCpf("12345678901")

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", customer.Name)
}

func TestGetByCpf_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{CPF: "12345678901"}).Return(nil, errors.New("repository error"))

	service := customer.NewCustomerService(mockRepo)
	customer, err := service.GetByCpf("12345678901")

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestUpdate_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}
	updatedCustomer := &domain.Customer{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{ID: 1}).Return(updatedCustomer, nil)
	mockRepo.EXPECT().Update(uint32(1), *updatedCustomer).Return(updatedCustomer, nil)

	service := customer.NewCustomerService(mockRepo)
	customer, err := service.Update(1, customerDto)

	assert.NoError(t, err)
	assert.Equal(t, "John Doe", customer.Name)
}

func TestUpdate_ValidationError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "JD", Email: "john@example.com", CPF: "12345678901"}

	service := customer.NewCustomerService(nil)
	customer, err := service.Update(1, customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestUpdate_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	customerDto := dto.CustomerDto{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}
	updatedCustomer := &domain.Customer{Name: "John Doe", Email: "john@example.com", CPF: "12345678901"}

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{ID: 1}).Return(updatedCustomer, nil)
	mockRepo.EXPECT().Update(uint32(1), *updatedCustomer).Return(nil, errors.New("repository error"))

	service := customer.NewCustomerService(mockRepo)
	customer, err := service.Update(1, customerDto)

	assert.Error(t, err)
	assert.Nil(t, customer)
}

func TestDelete_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{ID: 1}).Return(&domain.Customer{ID: 1}, nil)
	mockRepo.EXPECT().Delete(uint32(1)).Return(nil)

	service := customer.NewCustomerService(mockRepo)
	err := service.Delete(1)

	assert.NoError(t, err)
}

func TestDelete_NotFoundError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{ID: 1}).Return(nil, nil)

	service := customer.NewCustomerService(mockRepo)
	err := service.Delete(1)

	assert.Error(t, err)
	assert.Equal(t, "customer not found to delete", err.Error())
}

func TestDelete_RepositoryError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockCustomerRepository(ctrl)
	mockRepo.EXPECT().GetOne(domain.Customer{ID: 1}).Return(&domain.Customer{ID: 1}, nil)
	mockRepo.EXPECT().Delete(uint32(1)).Return(errors.New("repository error"))

	service := customer.NewCustomerService(mockRepo)
	err := service.Delete(1)

	assert.Error(t, err)
	assert.Equal(t, "error on delete in repository", err.Error())
}
