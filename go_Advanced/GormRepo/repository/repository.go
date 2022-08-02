package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"gorm/model"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error
	GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error
	GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}) error
	HardDelete(uow *UnitOfWork, out interface{}) error
	AddWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error

}

// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{DB: db.New(), committed: false, readOnly: true}
	}
	return &UnitOfWork{DB: db.New().Begin(), committed: false, readOnly: false}
}

// Complete marks end of unit of work
func (uow *UnitOfWork) Complete() {
	if !uow.committed && !uow.readOnly {
		uow.DB.Rollback()
	}
}

// Commit the transaction
func (uow *UnitOfWork) Commit() {
	if !uow.readOnly {
		uow.DB.Commit()
	}
	uow.committed = true
}

// GormRepository implements Repository
type GormRepository struct {
}

// NewRepository returns a new repository object
func NewRepository() Repository {
	return &GormRepository{}
}

// Get a record for specified entity with specific id
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.First(out, "id = ?", id).Error
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Find(out).Error
}

// GetAllForTenant returns all objects of specifeid tenantID
func (repository *GormRepository) GetAllForTenant(uow *UnitOfWork, out interface{}, tenantID uuid.UUID, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Where("tenantID = ?", tenantID).Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Create(entity).Error
}

func (repository *GormRepository) AddWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error {
	return uow.DB.Debug().Omit(omitFields...).Create(entity).Error	
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Model(entity).Update(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Delete(entity).Error
}


func (repository *GormRepository) HardDelete(uow *UnitOfWork, entity interface{}) error {
	var user model.User
	uow.DB.Where("id = ?", "0d1ddf9f-53e2-4440-9545-d72679582fed").Preload("Courses").First(&user)
	uow.DB.Debug().Model(entity).Association("Courses").Delete(user.Courses)//this is deleting the association of user_courses
	return uow.DB.Debug().Unscoped().Delete(entity).Error
	//this is hard deleting the user and the hobbies but not the entries from user_courses

	/*
	Delete Associations
	Remove relationship between source & argument objects, only delete the reference, won’t delete those objects from DB.
	db.Model(&user).Association("Languages").Delete([]Language{languageZH, languageEN})
	db.Model(&user).Association("Languages").Delete(languageZH, languageEN)
	*/
}


