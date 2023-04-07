# codegenner
Generates functional code using OpenAI's API (you need an api key)

The 'codegenner' application is a code generation tool that uses OpenAI's API to generate code snippets based on user input. The input is provided in a text file with a specific format, and the application outputs code snippets for each specified function.

## build it

```
git clone git@github.com:rickcollette/codegenner.git
go mod init codegenner
go mod tidy
go build -o codegenner main.go
```

## Usage:

1. Set the OPENAI_API_KEY environment variable with your OpenAI API key.
2. Create an input file with the following format:

```
OVERVIEW: <Overview of the code>
LANGUAGE: <Target programming language>
FUNCTION_NAME: <Function name>
DESCRIPTION: <Function description>
```

3. Run the application with the input file as a command-line argument:

```
./codegenner <input_filename>
```

4. The application will output generated code snippets for each specified function.

## The application has the following structure:

1. Import necessary packages, including the OpenAI API client library.
2. Define two structs: Function and Overview.
   - Function represents a single function, with a Name and Description.
   - Overview represents the overall purpose of the generated code, with a Type and Description.
3. Define the main function, which is the entry point of the application.
   - Read the API key from the OPENAI_API_KEY environment variable.
   - Check if the required command-line argument (filename) is provided.
   - Open the input file and defer closing it.
   - Initialize a bufio.Scanner to read the input file line by line.
   - Parse the input file to extract the Overview, Language, and a list of Functions.
   - Print the Overview to the console.
   - Iterate through the list of Functions and perform the following steps for each:
   - Create a prompt string by combining the language, function name, and function description.
   - Create an openai.CompletionRequest with the required parameters.
   - Call the OpenAI API with the generated prompt using the client.CreateCompletion method.
   - Check for errors and print an error message if the API call failed, then continue to the next function.
   - Print the generated code snippet for the function to the console.
