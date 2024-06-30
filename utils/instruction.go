package utils

import (
	"catalk/instructions"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var instructionsInstance *instructions.Instructions

func ReadInstructions(path string) (*instructions.Instructions, error) {
	if instructionsInstance != nil {
		return instructionsInstance, nil
	}
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
	ins := new(instructions.Instructions)
	if err := json.Unmarshal(byteVal, &ins); err != nil {
		log.Printf("error unmarshal body. Err: %s", err.Error())
		return nil, fmt.Errorf("unmarshal failed")
	}
	instructionsInstance = ins

	return instructionsInstance, nil

}
