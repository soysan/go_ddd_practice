package customer

import (
	"errors"
	"github.com/google/uuid"
	"github.com/soysan/go_ddd_practice/aggregate"
)

var (
	ErrCustomerNotFound   = errors.New("the customer was not found in the repository")
	ErrFaildToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer     = errors.New("failed to update the customer in the repository")
)

type Repository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
