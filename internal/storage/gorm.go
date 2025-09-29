package storage

import (
	"TP4/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(path string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//Automigration
	if err := db.AutoMigrate(&models.Contact{}); err != nil {
		return nil, err
	}
	return &GORMStore{db: db}, nil
}

func (gs *GORMStore) Add(contact *models.Contact) error {
	return gs.db.Create(contact).Error
}
func (gs *GORMStore) GetAll() ([]*models.Contact, error) {
	var contacts []*models.Contact
	if err := gs.db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}
func (gs *GORMStore) GetByID(id int) (*models.Contact, error) {
	var contact models.Contact
	if err := gs.db.First(&contact, id).Error; err != nil {
		return nil, err
	}
	return &contact, nil
}
func (gs *GORMStore) Update(id int, newName, newEmail string) error {
	contact, err := gs.GetByID(id)
	if err != nil {
		return err
	}
	contact.Name = newName
	contact.Email = newEmail
	return gs.db.Save(contact).Error
}
func (gs *GORMStore) Delete(id int) error {
	contact, err := gs.GetByID(id)
	if err != nil {
		return err
	}
	return gs.db.Delete(contact).Error
}
func (gs *GORMStore) Close() error {
	sqlDB, err := gs.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
