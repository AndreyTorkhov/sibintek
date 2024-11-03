package pkg

import (
    "encoding/json"
    "sibintek/internal/models"
    "os"
)

func ReadConfig(filename string) (models.Config, error) {
    file, err := os.Open(filename)
    if err != nil {
        return models.Config{}, err
    }
    defer file.Close()

    var config models.Config
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&config)
    return config, err
}
