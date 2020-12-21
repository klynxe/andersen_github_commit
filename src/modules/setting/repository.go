package setting

type Repository interface {
	SaveConnect(connect *Connect) error
	GetConnect() (*Connect, error)
}
