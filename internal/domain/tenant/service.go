package tenant

import (
	"mms-dbsd/internal/auth"
	"mms-dbsd/internal/domain/tenant/dto"
)

type Service interface {
	CreateTenant( req *dto.CreateTenantRequest) (*dto.TenantResponse, error)
	FindById(tenantId uint) (*dto.TenantResponse, error)
	FindByEmail(email string) (*dto.TenantResponse, error)
	Update(tenantId uint, req dto.UpdateTenantRequest)(*dto.TenantResponse, error)
	FindAll() (*[]dto.TenantResponse, error)
	Delete(tenantId uint) (*dto.TenantResponse, error)
}

type service struct {
	repository ITenantRepository
	hasher auth.PasswordHasher
	generateCode auth.GenerateCode
}

func NewService(repository ITenantRepository) Service {
	return &service{
		repository: repository,
		hasher: auth.NewPassowrdHasher(),
		generateCode: auth.NewCodeGenerator(),
	}
}

func (s *service) CreateTenant( req *dto.CreateTenantRequest) (*dto.TenantResponse, error) {

	hashedPassword, err := s.hasher.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	tenantCode, err := s.generateCode.GenerateToken("TEN", 6)
	if err != nil {
		return nil, err
	}

	tenant := &Tenant{
		TenantName: req.TenantName,
		TenantCode: tenantCode,
		Email:      req.Email,
		Password: hashedPassword,
		SubDomain:  req.SubDomain,
	}

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

func (s *service) FindAll() (*[]dto.TenantResponse, error) {
	tenants, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}
	var tenantResponses []dto.TenantResponse

	for _, tenant := range *tenants {
		tenantResponses = append(tenantResponses, *tenant.ToTenantResponse())
	}
	return &tenantResponses, nil
}

func (s *service) Delete(tenantID uint) (*dto.TenantResponse, error) {
	t, err := s.repository.Delete(tenantID)

	if err != nil {
		return nil, err
	}

	// TODO: delete tenant database
	return t.ToTenantResponse(), nil
}