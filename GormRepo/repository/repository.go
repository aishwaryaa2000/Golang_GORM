package repository

import (
	"gorm.io/gorm"
	"github.com/satori/go.uuid"
)

type Repository interface {
	Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string, q ...QueryProcessor) error
	GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	Delete(uow *UnitOfWork, out interface{}) error
	AddWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error
	RemoveAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) error 
	UpdateWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error

}

type QueryProcessor func(db *gorm.DB, out interface{}) (*gorm.DB, error)


// UnitOfWork represents a connection
type UnitOfWork struct {
	DB        *gorm.DB
	committed bool
	readOnly  bool
}

// NewUnitOfWork creates new UnitOfWork
func NewUnitOfWork(db *gorm.DB, readOnly bool) *UnitOfWork {
	if readOnly {
		return &UnitOfWork{DB: db, committed: false, readOnly: true}
	}
	return &UnitOfWork{DB: db.Begin(), committed: false, readOnly: false}
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
func (repository *GormRepository) Get(uow *UnitOfWork, out interface{}, id uuid.UUID, preloadAssociations []string,q ...QueryProcessor) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	
	if q != nil {
		var err error
		for _, queryProcessor := range q {
			db, err = queryProcessor(db, out)
			if err != nil {
				return err
			}
		}
	}

	return db.First(out).Error
}

// Filter will filter the results
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Where(condition, args...)
		return db, nil
	}
}

// GetAll retrieves all the records for a specified entity and returns it
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, preloadAssociations []string) error {
	db := uow.DB
	for _, association := range preloadAssociations {
		db = db.Preload(association)
	}
	return db.Find(out).Error
}

// Add specified Entity
func (repository *GormRepository) Add(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Create(entity).Error
}

// AddWithOmit add specified Entity by omitting passed fields
func (repository *GormRepository) AddWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error {
	return uow.DB.Debug().Omit(omitFields...).Create(entity).Error	
}

// Update specified Entity
func (repository *GormRepository) Update(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Model(entity).Updates(entity).Error
}

// Delete specified Entity
func (repository *GormRepository) Delete(uow *UnitOfWork, entity interface{}) error {
	return uow.DB.Delete(entity).Error
}

// RemoveAssociations removes associations from the given out entity
func (repository *GormRepository) RemoveAssociations(uow *UnitOfWork, out interface{}, associationName string, associations ...interface{}) error {
	if err := uow.DB.Model(out).Association(associationName).Delete(associations...); err != nil {
		return err
	}
	return nil
}

// UpdateWithOmit updates specified Entity by omitting passed fields
func (repository *GormRepository) UpdateWithOmit(uow *UnitOfWork, entity interface{}, omitFields []string) error {
	if err := uow.DB.Model(entity).Omit(omitFields...).Updates(entity).Error; err != nil {
		return err
	}
	return nil
}