package tenant

import (
	"errors"

	"gorm.io/gorm"
)

var ErrConflict = errors.New("conflict")


type Repository interface {
	CreateTenant( tenant *Tenant) error
	// FindByID(ctx context.Context, id uuid.UUID) (*Tenant, error)
	// FindByDatabaseName(ctx context.Context, dbName string) (*Tenant, error)
	// Update(ctx context.Context, tenant *Tenant) error
	// Delete(ctx context.Context, id uuid.UUID) error
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