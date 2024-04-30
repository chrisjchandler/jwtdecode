package main

import (
    "bufio"
    "encoding/base64"
    "encoding/json"
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the file path containing the JWT: ")
    filePath, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading file path:", err)
        return
    }

    filePath = strings.TrimSpace(filePath)
    tokenString, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading token from file:", err)
        return
    }

    // Split token
    parts := strings.Split(string(tokenString), ".")
    if len(parts) < 3 {
        fmt.Println("Error: JWT does not have three parts")
        return
    }

    // Decode payload
    payload, err := decodeSegment(parts[1])
    if err != nil {
        fmt.Println("Error decoding payload:", err)
        return
    }

    fmt.Println("Payload:", string(payload))

    var claims jwt.MapClaims
    err = json.Unmarshal(payload, &claims)
    if err != nil {
        fmt.Println("Error unmarshalling payload:", err)
        return
    }

    if exp, ok := claims["exp"].(float64); ok {
        expTime := time.Unix(int64(exp), 0)
        fmt.Printf("Expiration date: %s\n", expTime)
    } else {
        fmt.Println("Expiration date not found in the token")
    }
}

// Help 4 base64 segments
func decodeSegment(seg string) ([]byte, error) {
    if l := len(seg) % 4; l > 0 {
        seg += strings.Repeat("=", 4-l)
    }
    return base64.URLEncoding.DecodeString(seg)
}
