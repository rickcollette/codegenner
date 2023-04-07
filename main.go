package main

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strings"

    "github.com/sashabaranov/go-openai"
)

type Function struct {
    Name        string
    Description string
}

type Overview struct {
    Type        string
    Description string
}

func main() {
    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set the OPENAI_API_KEY environment variable.")
        os.Exit(1)
    }
    client := openai.NewClient(apiKey)

    if len(os.Args) < 2 {
        fmt.Println("Usage: codegenner <filename>")
        os.Exit(1)
    }
    filename := os.Args[1]

    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var overview Overview
    var functions []Function
    var language string
    for scanner.Scan() {
        line := scanner.Text()
        if strings.HasPrefix(line, "OVERVIEW:") {
            overview.Description = strings.TrimSpace(strings.TrimPrefix(line, "OVERVIEW:"))
        } else if strings.HasPrefix(line, "LANGUAGE:") {
            language = strings.TrimSpace(strings.TrimPrefix(line, "LANGUAGE:"))
        } else if strings.HasPrefix(line, "FUNCTION_NAME:") {
            functionName := strings.TrimSpace(strings.TrimPrefix(line, "FUNCTION_NAME:"))
            scanner.Scan()
            functionDescriptionLine := scanner.Text()
            functionDescription := strings.TrimSpace(strings.TrimPrefix(functionDescriptionLine, "DESCRIPTION:"))
            function := Function{
                Name:        functionName,
                Description: functionDescription,
            }
            functions = append(functions, function)
        }
    }

    fmt.Printf("Overview:\nType: %s\nDescription: %s\n", overview.Type, overview.Description)

    for _, function := range functions {
        prompt := fmt.Sprintf("Write a %s function for the following:\nFunction Name: %s\nDescription: %s", language, function.Name, function.Description)
        params := &openai.CompletionRequest{
            Prompt:     prompt,
            Model:      "text-davinci-002",
            MaxTokens:  64,
            Temperature: 0.7,
            TopP: 1,
            N: 1,
            Stop: []string{"*/"},
        }

        response, err := client.CreateCompletion(context.Background(), *params)

        if err != nil {
            fmt.Printf("Error generating code for function %s: %v\n", function.Name, err)
            continue
        }

        code := response.Choices[0].Text
        fmt.Printf("\nFunction:\n%s", code)
    }
}
