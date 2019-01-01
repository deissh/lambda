package scaling

type ServiceQuery struct {
	// название образа
	Name 			  string
	// текущее кол-во реплик
	Replicas          int
	// максимальное кол-во реплик
	MaxReplicas       int
	// минимальное кол-во реплик
	MinReplicas       int
}
