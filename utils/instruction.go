package utils

import (
	"catalk/instructions"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadInstructions(path string) (*instructions.CatInstructions, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Printf("error open instruction file. Err: %s", err.Error())
		return nil, fmt.Errorf("open instruction file failed")
	}
	defer jsonFile.Close()
	byteVal, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("error io read. Err: %s", err.Error())
		return nil, fmt.Errorf("read file failed")
	}
	instructions := new(instructions.CatInstructions)
	if err := json.Unmarshal(byteVal, &instructions); err != nil {
		log.Printf("error unmarshal body. Err: %s", err.Error())
		return nil, fmt.Errorf("unmarshal failed")
	}

	return instructions, nil

}
