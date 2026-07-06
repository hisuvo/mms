package tenant

import (
	"mms-dbsd/internal/domain/tenant/dto"
)

type Service interface {
	CreateTenant( req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
	FindById(tenantId uint) (*dto.TenantResponse, error)
	FindByEmail(email string) (*dto.TenantResponse, error)
	Update(tenantId uint, req dto.UpdateTenantRequest)(*dto.TenantResponse, error)
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

	return tenant.ToTenantResponse(), nil
}

func (s *service) FindById(tenantID uint) (*dto.TenantResponse, error) {
	tenant, err := s.repository.FindByID(tenantID)
	if err != nil {
		return nil, err
	}
	return tenant.ToTenantResponse(), nil
}

func (s *service) FindByEmail(email string) (*dto.TenantResponse, error) {
	tenant, err := s.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return tenant.ToTenantResponse(), nil
}

func (s *service) Update(tenantId uint, req dto.UpdateTenantRequest) (*dto.TenantResponse, error) {

	tenant,err:=s.repository.FindByID(tenantId)
	if err != nil {
		return nil, err
	}

	if req.Email != "" {
		tenant.Email=req.Email
	}

	if req.SubDomain != ""{
		tenant.SubDomain=req.SubDomain
	}

	if req.TenantName != ""{
		tenant.TenantName=req.TenantName
	}

	if err := s.repository.Update(tenant); err != nil {
		return nil, err
	}

	return tenant.ToTenantResponse(), nil
	
}