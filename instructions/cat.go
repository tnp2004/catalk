package instructions

type CatInstructions struct {
	MainInstruction string            `json:"mainInstruction"`
	Breeds          map[string]string `json:"breeds"`
}
