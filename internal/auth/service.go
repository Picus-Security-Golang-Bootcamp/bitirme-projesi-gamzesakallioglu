package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/internal/api"
	"github.com/gamze.sakallioglu/learningGo/bitirme-projesi-gamzesakallioglu/internal/models"
)

type authService struct {
	repo Repository
}

type Service interface {
	GetCustomer(ctx context.Context, email *string, password *string) (*models.Customer, error)
	GetUser(ctx context.Context, email *string, password *string) (*models.User, error)
	CheckCustomerExists(ctx context.Context, email *string) (*models.Customer, error)
	CreateCustomer(ctx context.Context, customer *api.CustomerSignUp) (*models.Customer, error)
	CreateUser(ctx context.Context, customer *api.UserSignUp) (*models.User, error)
}

func NewAuthService(repo Repository) Service {
	if repo == nil {
		return nil
	}

	return &authService{repo: repo}
}

func (a authService) GetCustomer(ctx context.Context, email *string, password *string) (*models.Customer, error) {
	if len(*email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}

	customer := a.repo.GetCustomer(ctx, email, password)
	return customer, nil
}

func (a authService) GetUser(ctx context.Context, email *string, password *string) (*models.User, error) {
	if len(*email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}

	user := a.repo.GetUser(ctx, email, password)
	return user, nil
}

func (a authService) CheckCustomerExists(ctx context.Context, email *string) (*models.Customer, error) {
	if len(*email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}

	customer := a.repo.CheckCustomerExists(ctx, email)
	return customer, nil
}

func (a authService) CreateCustomer(ctx context.Context, customer *api.CustomerSignUp) (*models.Customer, error) {

	customerToGo := models.Customer{
		Name:     customer.Name,
		Email:    customer.Email,
		Address:  customer.Address,
		Password: customer.Password,
		Phone:    customer.Phone,
	}
	customerExisted, err := a.CheckCustomerExists(ctx, customer.Email)

	if err != nil {
		return nil, err
	}

	if customerExisted != nil {
		return nil, errors.New("this email already taken by another user")
	}

	err = a.repo.CreateCustomer(ctx, &customerToGo)
	if err != nil {
		return nil, err
	}

	return &customerToGo, nil

}

func (a authService) CreateUser(ctx context.Context, user *api.UserSignUp) (*models.User, error) {

	userToGo := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Phone:    user.Phone,
	}
	customerExisted, err := a.CheckCustomerExists(ctx, user.Email)

	if err != nil {
		return nil, err
	}

	if customerExisted != nil {
		return nil, errors.New("this email already taken by another user")
	}

	err = a.repo.CreateUser(ctx, &userToGo)
	if err != nil {
		return nil, err
	}

	return &userToGo, nil

}
