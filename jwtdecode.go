package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

func main() {
    // reader for stdin
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter the file path containing the JWT: ")
    filePath, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading file path:", err)
        return
    }

    filePath = strings.TrimSpace(filePath)

    // Read the JWT 
    tokenString, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println("Error reading token from file:", err)
        return
    }

    // Parse 
    token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
        // Return nil b/c not valid signature here
        return nil, nil
    })
    if err != nil {
        fmt.Println("Error parsing token:", err)
        return
    }

    // validate & extract claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // print expiration date
        if exp, ok := claims["exp"].(float64); ok {
            expTime := time.Unix(int64(exp), 0)
            fmt.Printf("Expiration date: %s\n", expTime)
        } else {
            fmt.Println("Expiration date not found in the token")
        }
    } else {
        fmt.Println("Invalid token")
    }
}
