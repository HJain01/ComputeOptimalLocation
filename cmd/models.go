package cmd

type Location struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

type LocationVariance struct {
	Location Location `json:"location"`
	Variance float64  `json:"variance"`
}
