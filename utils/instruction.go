package utils

import (
	"catalk/instructions"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var instructionMap = make(map[string]*instructions.Instructions)

func ReadInstructions(path string) (*instructions.Instructions, error) {
	if _, ok := instructionMap[path]; ok {
		return instructionMap[path], nil
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
	instructionMap[path] = ins

	return instructionMap[path], nil

}
