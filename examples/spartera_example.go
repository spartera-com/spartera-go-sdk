package main

import (
    "context"
    "fmt"
    "os"
    
    sparteraapi "github.com/spartera-com/spartera-go-sdk"
)

func main() {
    fmt.Println("🚀 Spartera API Go SDK Example")
    fmt.Println("========================================")
    
    // Configure the client
    configuration := sparteraapi.NewConfiguration()
    
    // Set API base URL
    apiBaseURL := os.Getenv("SPARTERA_API_BASE_URL")
    if apiBaseURL == "" {
        apiBaseURL = "https://api.spartera.com"
    }
    configuration.Host = apiBaseURL
    
    // IMPORTANT: Spartera uses X-API-Key header authentication
    apiKey := os.Getenv("SPARTERA_API_KEY")
    if apiKey == "" {
        apiKey = "your-api-key-here"
        fmt.Println("⚠️  Set SPARTERA_API_KEY environment variable with your actual API key")
    }
    configuration.AddDefaultHeader("X-API-Key", apiKey)
    
    // Create API client
    apiClient := sparteraapi.NewAPIClient(configuration)
    ctx := context.Background()
    
    // Get company ID from environment
    companyId := os.Getenv("SPARTERA_COMPANY_ID")
    if companyId == "" {
        companyId = "your-company-id"
        fmt.Println("⚠️  Set SPARTERA_COMPANY_ID environment variable with your company ID")
    }
    
    fmt.Println()
    fmt.Printf("Configured client for: %s\n", configuration.Host)
    fmt.Printf("Company ID: %s\n", companyId)
    fmt.Printf("API Client created: %v\n", apiClient != nil)
    fmt.Printf("Context created: %v\n", ctx != nil)
    
    fmt.Println("\n✅ SDK basic initialization test completed successfully!")
    fmt.Println("\n📝 To test with real API calls, set your credentials and modify this example")
    fmt.Println("   export SPARTERA_API_KEY='your-actual-api-key'")
    fmt.Println("   export SPARTERA_COMPANY_ID='your-actual-company-id'")
}
