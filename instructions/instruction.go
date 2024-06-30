package instructions

type Instructions struct {
	MainInstruction   string            `json:"mainInstruction"`
	BreedsInstruction map[string]string `json:"breeds"`
}
