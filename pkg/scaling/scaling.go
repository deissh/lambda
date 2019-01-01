package scaling

type ScalingConfig struct {
	ServiceQuery ServiceQuery
	// кол-во попыток запустить еще функции
	SetScaleRetries uint
}

func Create (service ServiceQuery) ScalingConfig {

	return ScalingConfig{
		ServiceQuery: service,
		SetScaleRetries: 3,
	}
}