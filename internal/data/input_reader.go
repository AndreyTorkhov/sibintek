package data

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func ReadNumbersFromInput() ([]int, error) {
    fmt.Print("Enter numbers separated by spaces: ")
    scanner := bufio.NewScanner(os.Stdin)
    if scanner.Scan() {
        input := scanner.Text()
        return parseInput(input)
    }
    return nil, nil
}

func parseInput(input string) ([]int, error) {
    var numbers []int
    fields := strings.Fields(input)
    for _, field := range fields {
        num, err := strconv.Atoi(field)
        if err != nil {
            return nil, err
        }
        numbers = append(numbers, num)
    }
    return numbers, nil
}
