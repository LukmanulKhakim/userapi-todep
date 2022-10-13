package domain

type Core struct {
	ID     uint
	Name   string
	HP     string
	Addres string
}

type Repository interface {
	Insert(newUser Core) (Core, error)
	Update(updatedData Core, ID uint) (Core, error)
	Get(ID uint) (Core, error)
	GetAll() ([]Core, error)
	Delete(ID uint) (Core, error)
	// GetUser(newUser Core) (Core, error)
}

type Service interface {
	AddUser(newUser Core) (Core, error)
	UpdateProfile(updatedData Core, ID uint) (Core, error)
	Profile(ID uint) (Core, error)
	ShowAllUser() ([]Core, error)
	DeleteUser(ID uint) (Core, error)
	// Login(newUser Core) (Core, error)
}
