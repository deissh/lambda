package typings

type FunctionInfo struct {
	Name    string `json:"name" binding:"required"`
	Uuid    string `json:"uuid" binding:"required"`
	Runtime struct{
		Executor string `json:"executor" binding:"required"`
		Cmd string `json:"cmd" binding:"required"`
		Params string `json:"params"`
		Env []struct{
			Name string `json:"name"`
			Value string `json:"value"`
		} `json:"env"`
	} `json:"runtime" binding:"required"`
	Repository struct{
		Url string `json:"url" binding:"required"`
	} `json:"repository" binding:"required"`
}