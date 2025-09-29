package storage

import (
	"TP4/internal/models"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type JSONStore struct {
	filePath string
	contacts map[int]*models.Contact
	nextID   int
}

// Constructeur
func NewJSONStore(filePath string) (*JSONStore, error) {
	store := &JSONStore{
		filePath: filePath,
		contacts: make(map[int]*models.Contact),
		nextID:   1,
	}
	if err := store.load(); err != nil {
		return nil, err
	}
	return store, nil
}

func (js *JSONStore) load() error {
	data, err := os.ReadFile(js.filePath)
	if os.IsNotExist(err) {
		return nil
	}
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	var loadedContacts []*models.Contact
	if err := json.Unmarshal(data, &loadedContacts); err != nil {
		return err
	}
	for _, contact := range loadedContacts {
		js.contacts[contact.ID] = contact
		if contact.ID >= js.nextID {
			js.nextID = contact.ID + 1
		}
	}
	return nil
}
func (js *JSONStore) save() error {
	var allContacts []*models.Contact
	for _, contact := range js.contacts {
		allContacts = append(allContacts, contact)
	}
	data, err := json.Marshal(allContacts)
	if err != nil {
		return err
	}
	return os.WriteFile(js.filePath, data, 0644)
}
func (js *JSONStore) Add(contact *models.Contact) error {
	contact.ID = js.nextID
	js.contacts[contact.ID] = contact
	js.nextID++
	contact.CreatedAt = time.Now()
	contact.UpdatedAt = time.Now()
	return js.save()
}
func (js *JSONStore) GetAll() ([]*models.Contact, error) {
	var allContacts []*models.Contact
	for _, contact := range js.contacts {
		allContacts = append(allContacts, contact)
	}
	return allContacts, nil
}
func (js *JSONStore) GetByID(id int) (*models.Contact, error) {
	contact, ok := js.contacts[id]
	if !ok {
		return nil, os.ErrNotExist
	}
	return contact, nil
}
func (js *JSONStore) Update(id int, newName, newEmail string) error {
	contact, err := js.GetByID(id)
	if err != nil {
		return err
	}
	if newName == "" && newEmail == "" {
		return errors.New("au moins un des champs 'nom' ou 'email' doit être fourni")
	}
	if newName != "" {
		contact.Name = newName
	}
	if newEmail != "" {
		contact.Email = newEmail
	}
	contact.UpdatedAt = time.Now()
	return js.save()
}
func (js *JSONStore) Delete(id int) error {
	_, err := js.GetByID(id)
	if err != nil {
		return err
	}
	delete(js.contacts, id)
	return js.save()
}
func (js *JSONStore) Close() error {
	return nil // JSONStore n'a pas de ressource à fermer
}
