package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"
	"time"
)

type Service interface {
	CreateTenant( req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) CreateTenant( req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {
	tenant := &Tenant{
		TenantName: req.TenantName,
		Email:      req.Email,
		SubDomain:  req.SubDomain,
	}

	// TODO: create tenant database
	if err := s.repository.CreateTenant(tenant); err != nil {
		return nil, err
	}

	return &dto.TenantResponse{
		ID:         uint64(tenant.ID),
		TenantName: tenant.TenantName,
		Email:      tenant.Email,
		SubDomain:  tenant.SubDomain,
		CreatedAt:  tenant.CreatedAt.Format(time.RFC3339),
	}, nil
}