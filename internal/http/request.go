package http

import (
    "fmt"
    "net/http"
)

func CheckURL(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        return fmt.Errorf("failed to perform HTTP GET request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        return nil
    }
    return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}
