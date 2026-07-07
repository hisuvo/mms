package tenant

import (
	"errors"

	"gorm.io/gorm"
)

var(
	ErrConflict = errors.New("conflict")
	ErrNotFound = errors.New("tenant not found")
)


type Repository interface {
	CreateTenant( tenant *Tenant) error
	FindByID(tenantId uint) (*Tenant, error)
	FindByEmail(email string) (*Tenant, error)
	FindAll() (*[]Tenant, error)
	Update(tenant *Tenant) error
	Delete( tenantId uint) (*Tenant, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateTenant(tenant *Tenant) error {
	var existing Tenant
	var err error

	if err = r.db.Where("email = ?", tenant.Email).First(&existing).Error; err == nil {
		return errors.New("email already exists")
	}

	if err = r.db.Where("sub_domain = ?", tenant.SubDomain).First(&existing).Error; err == nil {
		return errors.New("sub domain already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err := r.db.Create(tenant).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(tenantId uint) (*Tenant, error){
	var tenant Tenant

	err := r.db.First(&tenant, tenantId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &tenant, nil
}

func (r *repository) FindByEmail(email string) (*Tenant, error){
	var tenant Tenant

	err := r.db.Where("email = ?", email).First(&tenant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &tenant, nil
}

func (r *repository) Update(tenant *Tenant) error {
	if err := r.db.Save(tenant).Error; err != nil {
		return  err
	}

	return nil
}

func (r *repository) FindAll() (*[]Tenant, error){
	var tenants []Tenant

	err := r.db.Find(&tenants).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &tenants, nil
}

func (r *repository) Delete(tenantId uint) (*Tenant, error) {
	var tenant Tenant

	if err := r.db.Where("id = ?", tenantId).First(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, ErrNotFound
		}
		return nil, err
	}

	if err := r.db.Delete(&tenant).Error; err != nil {
		return nil, err
	}

	return &tenant, nil
}