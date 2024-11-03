package data

import (
    "encoding/json"
    "sibintek/internal/models"
    "os"
)

func ReadNumbersFromFile(filename string) ([]int, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var data models.NumbersData
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&data)
    return data.Numbers, err
}
