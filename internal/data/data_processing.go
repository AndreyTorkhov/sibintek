package data

import (
    "fmt"
    "os"
)

func CalculateSum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }
    return sum
}

func SaveResult(filename string, sum int) error {
    resultFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer resultFile.Close()

    _, err = resultFile.WriteString(fmt.Sprintf("Total sum of numbers: %d\n", sum))
    return err
}
