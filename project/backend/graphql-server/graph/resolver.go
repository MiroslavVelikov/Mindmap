package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type ServiceProjectInterface interface {
}

type ServiceUserInterface interface {
}

type ServiceNodeInterface interface {
}

type Resolver struct{
	
}

func NewResolver() *Resolver {
	return &Resolver{}
}
