package main

import (
    "fmt"
    "path/filepath"
    "sibintek/internal/app"
    "sibintek/internal/logger"
    "sibintek/internal/pkg"
    "os"
)

func main() {
    projectRoot, err := os.Getwd()
    if err != nil {
        logger.LogError(fmt.Errorf("failed to determine project root: %w", err))
        return
    }

    configPath := filepath.Join(projectRoot, "config.json")
    config, err := pkg.ReadConfig(configPath)
    if err != nil {
        logger.LogError(fmt.Errorf("failed to read config file at %s: %w", configPath, err))
        return
    }

    if err := app.Run(os.Args, config); err != nil {
        logger.LogError(err)
    }
}
