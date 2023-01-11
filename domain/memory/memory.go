package memory

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/soysan/go_ddd_practice/aggregate"
	"github.com/soysan/go_ddd_practice/domain/customer"
	"sync"
)

type Repository struct {
	Customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *Repository {
	return &Repository{
		Customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func (mr *Repository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.Customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (mr *Repository) Add(c aggregate.Customer) error {
	if mr.Customers == nil {
		mr.Lock()
		mr.Customers = make(map[uuid.UUID]aggregate.Customer)
		mr.Unlock()
	}

	if _, ok := mr.Customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists: %w", customer.ErrFaildToAddCustomer)
	}

	mr.Lock()
	mr.Customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *Repository) Update(c aggregate.Customer) error {
	if _, ok := mr.Customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists %w", customer.ErrUpdateCustomer)
	}
	mr.Lock()
	mr.Customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
