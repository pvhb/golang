package service

import (
	"encoding/json"
	"errors"
	"os"
)

type Data struct {
	ID   int    `json:"id"`
	Roll string `json:"roll"`
	Name string `json:"name"`
}

type Payload struct {
	Data []Data
}

func fetchData() ([]Data, error) {
	r, err := os.ReadFile("./data.json")
	if err != nil {
		return nil, err
	}

	var pyld Payload
	err = json.Unmarshal(r, &pyld.Data)
	if err != nil {
		return nil, err
	}
	return pyld.Data, nil
}

func GetAll() ([]Data, error) {
	data, err := fetchData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetByID(idx int) (any, error) {
	data, err := fetchData()
	if err != nil {
		return nil, err
	}

	if idx > len(data) {
		return nil, errors.New("your index exceeded the limit")
	}
	return data[idx-1], nil
}
