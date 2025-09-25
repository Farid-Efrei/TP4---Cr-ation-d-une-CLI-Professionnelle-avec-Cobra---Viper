package storage

import (
	"TP4/internal/models"
	"errors"
)

// MemoryStore est une implémentation en mémoire de Storer
type MemoryStore struct {
	contacts map[int]*models.Contact
	nextID   int
}

// Constructeur
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: make(map[int]*models.Contact),
		nextID:   1,
	}
}

func (ms *MemoryStore) Add(contact *models.Contact) error {
	contact.ID = ms.nextID
	ms.contacts[contact.ID] = contact
	ms.nextID++
	return nil
}
func (ms *MemoryStore) GetAll() ([]*models.Contact, error) {
	var allContacts []*models.Contact
	for _, contact := range ms.contacts {
		allContacts = append(allContacts, contact)
	}
	return allContacts, nil
}
func (ms *MemoryStore) GetByID(id int) (*models.Contact, error) {
	contact, ok := ms.contacts[id]
	if !ok {
		return nil, errors.New("contact not found")
	}
	return contact, nil
}
func (ms *MemoryStore) Update(id int, newName, newEmail string) error {
	contact, err := ms.GetByID(id)
	if err != nil {
		return err
	}
	contact.Name = newName
	contact.Email = newEmail
	return nil
}
func (ms *MemoryStore) Delete(id int) error {
	_, err := ms.GetByID(id)
	if err != nil {
		return err
	}
	delete(ms.contacts, id)
	return nil
}
