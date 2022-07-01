package app

type Application interface{}

type app struct {
	storage Storage
}

type Storage interface{}

func New(storage Storage) Application {
	return &app{
		storage: storage,
	}
}
