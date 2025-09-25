package storage

import "TP4/internal/models"

type Storer interface {
	Add(contact *models.Contact) error
	GetAll() ([]*models.Contact, error)
	GetByID(id int) (*models.Contact, error)
	Update(id int, newName, newEmail string) error
	Delete(id int) error
}
