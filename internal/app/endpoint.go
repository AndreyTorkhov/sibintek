package app

import (
    "fmt"
    "sibintek/internal/data"
    "sibintek/internal/http"
    "sibintek/internal/logger"
    "sibintek/internal/models"
)

func Run(args []string, config models.Config) error {
    logger.LogInfo("Starting Run function")

    if len(args) < 3 || (args[1] == "file" && len(args) < 4) {
        fmt.Println("Usage: go run main.go <source> <output_file> [data_file]")
        return nil
    }

    source := args[1]
    outputFile := args[2]
    var numbers []int
    var err error

    if source == "file" {
        dataFile := args[3]
        logger.LogInfo("Reading numbers from file: " + dataFile)
        numbers, err = data.ReadNumbersFromFile(dataFile)
        if err != nil {
            logger.LogError(fmt.Errorf("failed to read numbers from file: %w", err))
            return err
        }
        logger.LogInfo(fmt.Sprintf("Numbers read from file successfully: %v", numbers))
    } else if source == "stdin" {
        logger.LogInfo("Reading numbers from standard input")
        numbers, err = data.ReadNumbersFromInput()
        if err != nil {
            logger.LogError(fmt.Errorf("failed to read numbers from stdin: %w", err))
            return err
        }
        logger.LogInfo(fmt.Sprintf("Numbers read from stdin successfully: %v", numbers))
    } else {
        fmt.Println("Invalid source or missing data file.")
        return nil
    }

    sum := data.CalculateSum(numbers)
    logger.LogInfo(fmt.Sprintf("Sum of numbers calculated: %d", sum))

    if err := data.SaveResult(outputFile, sum); err != nil {
        logger.LogError(fmt.Errorf("failed to save result: %w", err))
        return err
    }
    logger.LogInfo("Result saved to " + outputFile)

    logger.LogInfo("Performing HTTP GET request to " + config.URL)
    if err := http.CheckURL(config.URL); err != nil {
        logger.LogError(fmt.Errorf("HTTP request failed: %w", err))
    } else {
        logger.LogInfo(fmt.Sprintf("HTTP request to %s returned status 200", config.URL))
    }

    return nil
}
