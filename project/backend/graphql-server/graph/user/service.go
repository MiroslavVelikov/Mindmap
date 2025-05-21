package user

type RepositoryUserInterface interface {
}

type ServiceUser struct {
	repository RepositoryUserInterface
}

func NewServiceUser(repository RepositoryUserInterface) *ServiceUser {
	return &ServiceUser{
		repository: repository,
	}
}
