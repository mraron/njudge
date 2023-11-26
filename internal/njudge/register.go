package njudge

import "context"

type RegisterRequest struct {
	Name     string
	Email    string
	Password string
}

type RegisterService interface {
	Register(ctx context.Context, req RegisterRequest) (*User, error)
}

type registerService struct {
	u Users
}

func NewRegisterService(u Users) *registerService {
	return &registerService{
		u: u,
	}
}

func (rs *registerService) Register(ctx context.Context, req RegisterRequest) (*User, error) {
	u, err := NewUser(req.Name, req.Email, "user")
	if err != nil {
		return nil, err
	}

	if err := u.SetPassword(req.Password); err != nil {
		return nil, err
	}

	return rs.u.Insert(ctx, *u)
}
