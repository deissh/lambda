package manager

type managerCore struct {
	workers []struct{
		name string
		ip string
		port int
	}
}

func Init () managerCore {

	return managerCore{}
}