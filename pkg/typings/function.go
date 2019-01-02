package typings

type FunctionInfo struct {
	Name    string `json:"name" binding:"required"`
	Uuid    string `json:"uuid" binding:"required"`
	Runtime struct {
		Executor string `json:"executor" binding:"required"`
		Cmd      string `json:"cmd"`
		Params   string `json:"params"`
		Env      []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"env"`
	} `json:"runtime" binding:"required"`
	Repository struct {
		ImageName string `json:"image" binding:"required"`
	} `json:"repository" binding:"required"`
	Service struct {
		Port string `json:"port" binding:"required"`
		Host string `json:"host" binding:"required"`
	} `json:"service" binding:"required"`
}
