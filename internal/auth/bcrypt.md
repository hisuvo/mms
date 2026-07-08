A typical **Clean Architecture** in Go keeps password hashing in the **service/use case layer**, not in the handler or repository.

## Project Structure

```text
internal/
├── domain/
│   └── user/
│       ├── entity.go
│       ├── repository.go
│       ├── service.go
│       └── password.go
├── infrastructure/
│   └── password/
│       └── bcrypt.go
├── handler/
│   └── user_handler.go
└── dto/
    ├── request.go
    └── response.go
```

---

## 1. Entity

`internal/domain/user/entity.go`

```go
package user

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}
```

---

## 2. Password Interface

`internal/domain/user/password.go`

```go
package user

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}
```

---

## 3. Repository Interface

`internal/domain/user/repository.go`

```go
package user

type Repository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}
```

---

## 4. Service

`internal/domain/user/service.go`

```go
package user

type Service struct {
	repo   Repository
	hasher PasswordHasher
}

func NewService(repo Repository, hasher PasswordHasher) *Service {
	return &Service{
		repo:   repo,
		hasher: hasher,
	}
}

func (s *Service) Register(name, email, password string) error {
	hash, err := s.hasher.Hash(password)
	if err != nil {
		return err
	}

	u := &User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return s.repo.Create(u)
}

func (s *Service) Login(email, password string) error {
	u, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	return s.hasher.Compare(u.Password, password)
}
```

---

## 5. Bcrypt Implementation

`internal/infrastructure/password/bcrypt.go`

```go
package password

import "golang.org/x/crypto/bcrypt"

type Bcrypt struct{}

func New() *Bcrypt {
	return &Bcrypt{}
}

func (b *Bcrypt) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func (b *Bcrypt) Compare(hash, password string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)
}
```

---

## 6. Handler

```go
func (h *Handler) Register(c echo.Context) error {
	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := h.service.Register(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		return err
	}

	return c.JSON(201, "User created")
}
```

---

## 7. Dependency Injection

```go
repo := repository.New(db)
hasher := password.New()

service := user.NewService(repo, hasher)

handler := handler.New(service)
```

## Request

```json
{
  "name": "Suvo",
  "email": "suvo@gmail.com",
  "password": "123456"
}
```

## Database

```text
ID:        1
Name:      Suvo
Email:     suvo@gmail.com
Password:  $2a$10$M4cFv3A0jv5d...
```

The original password (`123456`) is **never stored**.

### Request (Login)

```json
{
  "email": "suvo@gmail.com",
  "password": "123456"
}
```

Flow:

```text
HTTP Request
      │
      ▼
Handler
      │
      ▼
Service
      │
      ├── FindByEmail()
      ▼
Repository
      │
      ▼
Database
      │
      ▼
Service
      │
      ├── bcrypt.CompareHashAndPassword()
      ▼
JWT Generation (if password matches)
      │
      ▼
HTTP Response
```

This design follows Clean Architecture because:

* **Handler**: Receives HTTP requests and responses only.
* **Service (Use Case)**: Contains business logic (registration, login, password hashing/comparison).
* **Repository**: Handles database operations only.
* **PasswordHasher**: Defined as an interface in the domain, implemented by bcrypt in the infrastructure layer, allowing you to replace bcrypt without changing business logic.
