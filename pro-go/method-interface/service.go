package main

type Service struct {
	description   string
	durationMonth int
	monthlyFee    float64
}

func (service *Service) getName() string {
	return service.description
}

func (service *Service) getCost(recur bool) float64 {
	if recur {
		return service.monthlyFee * float64(service.durationMonth)
	}
	return service.monthlyFee
}
