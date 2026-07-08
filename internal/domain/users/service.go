package users

import (
	"errors"
	"mms-dbsd/internal/domain/tenant"
	"mms-dbsd/internal/domain/users/dto"
	"strconv"
)

type IRegisterService interface {
	Register(request *dto.RegisterRequest) (*dto.UserResponse, error)
}

type registerService struct {
	repo  IRegisterRepository
	tRepo tenant.ITenantRepository
}

func NewRegisterService(repo IRegisterRepository, tRepo tenant.ITenantRepository) IRegisterService {
	return &registerService{
		repo:  repo,
		tRepo: tRepo,
	}
}

func (s *registerService) Register(request *dto.RegisterRequest) (*dto.UserResponse, error) {
	user := &User{
		UserName: request.UserName,
		TenantID: request.TenantID,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
	}

	tenantID, err := strconv.Atoi(request.TenantID)

	if err != nil {
		return nil, errors.New("invalid tenant ID")
	}

	if _, err := s.tRepo.FindByID(uint(tenantID)); err != nil {
		return nil, errors.New("tenant Id " + strconv.Itoa(tenantID) + " not valid for create user")
	}

	if err := s.repo.Register(user); err != nil {
		return nil, err
	}

	return user.ToUserResponse(), nil
}