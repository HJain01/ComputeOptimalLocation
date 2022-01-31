package src

type location struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type locationVariance struct {
	Location location `json:"location"`
	Variance float64  `json:"variance"`
}
