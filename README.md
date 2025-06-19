# Go SDK Documentation

## Overview
The Go SDK for Spartera API provides a convenient way to interact with the Spartera platform from Go applications.

**Install:** `go get github.com/spartera-com/spartera-api-go-sdk`

## Requirements
- Go 1.19 or higher
- Go modules

## 🚀 Sell Your First Data Product in 4 API Calls

Transform your data into revenue in under 5 minutes! Here's how to create and sell a data product on the Spartera marketplace:

```go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    
    sparteraapi "github.com/spartera-com/spartera-api-go-sdk"
)

const (
    COMPANY_ID = "your-company-id"
    USER_ID    = "your-user-id"
)

func main() {
    // Configure client
    config := sparteraapi.NewConfiguration()
    config.Host = "api.spartera.com"
    config.Scheme = "https"
    config.AddDefaultHeader("X-API-Key", "your-api-key-here")
    
    client := sparteraapi.NewAPIClient(config)
    ctx := context.Background()

    if err := sellDataProduct(ctx, client); err != nil {
        log.Fatal(err)
    }
}

func sellDataProduct(ctx context.Context, client *sparteraapi.APIClient) error {
    // Step 1: Discover available data platforms
    fmt.Println("🔍 Step 1: Discovering available platforms...")
    
    engines, _, err := client.CloudProvidersApi.CloudProvidersGet(ctx).Execute()
    if err != nil {
        return fmt.Errorf("failed to get cloud providers: %w", err)
    }
    
    bigqueryEngineID := int32(1) // BigQuery engine ID
    fmt.Printf("✅ Found %d supported platforms\n", len(engines))

    // Step 2: Create a data connection (with credentials in one call!)
    fmt.Println("🔗 Step 2: Creating BigQuery connection...")
    
    // Your BigQuery service account JSON (replace with your actual credentials)
    serviceAccountJSON := map[string]interface{}{
        "type":                        "service_account",
        "project_id":                  "your-project-id",
        "private_key_id":              "key-id",
        "private_key":                 "-----BEGIN PRIVATE KEY-----\nYOUR_PRIVATE_KEY\n-----END PRIVATE KEY-----\n",
        "client_email":                "your-service@your-project.iam.gserviceaccount.com",
        "client_id":                   "client-id",
        "auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
        "token_uri":                   "https://oauth2.googleapis.com/token",
    }
    
    credentialsJSON, _ := json.Marshal(serviceAccountJSON)
    
    connectionData := sparteraapi.Connection{
        CompanyId:             &COMPANY_ID,
        UserId:                &USER_ID,
        EngineId:              &bigqueryEngineID,
        Name:                  sparteraapi.PtrString("My BigQuery Data Warehouse"),
        Description:           sparteraapi.PtrString("Connection to our company's analytics data"),
        Visibility:            sparteraapi.PtrString("PRIVATE"),
        CredentialType:        sparteraapi.PtrString("SERVICE_ACCOUNT"),
        Credentials:           sparteraapi.PtrString(string(credentialsJSON)),
        VerifiedUsageAbility:  sparteraapi.PtrString("true"), // Legal compliance - you have rights to this data
    }

    connection, _, err := client.ConnectionsApi.CompaniesCompanyIdConnectionsPost(ctx, COMPANY_ID).Connection(connectionData).Execute()
    if err != nil {
        return fmt.Errorf("failed to create connection: %w", err)
    }
    
    connectionID := connection.GetConnectionId()
    fmt.Printf("✅ Created connection: %s\n", connectionID)

    // Step 3: Create a marketplace asset
    fmt.Println("📊 Step 3: Creating marketplace asset...")
    
    assetData := sparteraapi.Asset{
        Name:               sparteraapi.PtrString("Average Temperature Analytics"),
        Description:        sparteraapi.PtrString("Real-time weather temperature analytics from our IoT sensors across major cities"),
        CompanyId:          &COMPANY_ID,
        ConnectionId:       &connectionID,
        AssetType:          sparteraapi.PtrString("CALCULATION"),
        SqlLogic:           sparteraapi.PtrString("SELECT AVERAGE(temperature) AS avg_temp, city, COUNT(*) AS readings FROM `your-project.weather.sensor_data` WHERE timestamp >= TIMESTAMP_SUB(CURRENT_TIMESTAMP(), INTERVAL 24 HOUR) GROUP BY city ORDER BY avg_temp DESC"),
        SellInMarketplace:  sparteraapi.PtrString("true"), // 🔥 This makes it available for purchase!
        Source:            sparteraapi.PtrString("MANUAL"),
        Visibility:        sparteraapi.PtrString("PUBLIC"),
    }

    asset, _, err := client.AssetsApi.CompaniesCompanyIdAssetsPost(ctx, COMPANY_ID).Asset(assetData).Execute()
    if err != nil {
        return fmt.Errorf("failed to create asset: %w", err)
    }
    
    assetID := asset.GetAssetId()
    fmt.Printf("✅ Created marketplace asset: %s\n", assetID)

    // Step 4: Set your price and start earning!
    fmt.Println("💰 Step 4: Setting price...")
    
    priceData := sparteraapi.Assetpricehistory{
        PriceUsd: sparteraapi.PtrFloat64(2.00), // $2.00 per analysis (credits calculated automatically)
    }

    price, _, err := client.AssetPriceHistoryApi.CompaniesCompanyIdAssetsAssetIdPricesPost(ctx, COMPANY_ID, assetID).Assetpricehistory(priceData).Execute()
    if err != nil {
        return fmt.Errorf("failed to set price: %w", err)
    }
    
    fmt.Printf("✅ Price set: $%.2f (≈%s credits)\n", price.GetPriceUsd(), price.GetPriceCredits())

    fmt.Println("\n🎉 SUCCESS! Your data product is now live on the Spartera marketplace!")
    fmt.Printf("📈 Asset URL: https://marketplace.spartera.com/assets/%s\n", assetID)
    fmt.Println("💡 Customers can now discover and purchase your analytics!")
    
    return nil
}
```

**That's it!** You're now selling data analytics. Every time someone runs your analysis, you earn money! 

### 🎯 What You Just Built
- **Data Connection**: Secure link to your BigQuery warehouse
- **Analytics Product**: Temperature analysis that buyers can purchase
- **Marketplace Listing**: Your product is discoverable by thousands of potential customers
- **Automated Pricing**: Credits are calculated automatically based on your USD price

### 💰 Revenue Model
- You set the price ($2.00 in this example)
- Customers pay in credits (auto-converted)
- You earn revenue each time someone uses your analytics
- Spartera handles billing, payments, and customer support

---

## 📚 Detailed Documentation

### Authentication

Get your API key from the [Spartera Dashboard](https://app.spartera.com/settings/api-keys):

```go
package main

import (
    "os"
    sparteraapi "github.com/spartera-com/spartera-api-go-sdk"
)

func main() {
    // Option 1: Direct configuration
    config := sparteraapi.NewConfiguration()
    config.Host = "api.spartera.com"
    config.Scheme = "https"
    config.AddDefaultHeader("X-API-Key", "your-api-key-here")
    
    // Option 2: Environment variables (recommended)
    config = sparteraapi.NewConfiguration()
    config.Host = "api.spartera.com"
    config.Scheme = "https"
    config.AddDefaultHeader("X-API-Key", os.Getenv("SPARTERA_API_KEY"))
    
    client := sparteraapi.NewAPIClient(config)
}
```

### Environment Variables

```bash
export SPARTERA_API_KEY="your-api-key"
export SPARTERA_COMPANY_ID="your-company-id"
export SPARTERA_API_BASE_URL="https://api.spartera.com"
```

### Connection Types

Create connections to different data platforms:

```go
// BigQuery
bigqueryCredentials := map[string]interface{}{
    "type":       "service_account",
    "project_id": "your-project",
    // ... your service account JSON
}

// Snowflake
snowflakeConnection := sparteraapi.Connection{
    CredentialType: sparteraapi.PtrString("USERNAME_PASSWORD"),
    Username:       sparteraapi.PtrString("your-username"),
    Password:       sparteraapi.PtrString("your-password"),
}

// API Data Source
apiConnection := sparteraapi.Connection{
    CredentialType: sparteraapi.PtrString("API_KEY"),
    ApiEndpoint:    sparteraapi.PtrString("https://api.yourcompany.com/data"),
    ApiKeyParam:    sparteraapi.PtrString("x-api-key"),
    Credentials:    sparteraapi.PtrString("your-api-key-value"),
}
```

### Asset Types

Create different types of marketplace products:

```go
// SQL-based Analytics
calculationAsset := sparteraapi.Asset{
    AssetType: sparteraapi.PtrString("CALCULATION"),
    SqlLogic:  sparteraapi.PtrString("SELECT COUNT(*) as total_sales, AVG(amount) as avg_order FROM sales WHERE date >= CURRENT_DATE()"),
}

// Visualization/Dashboard
visualizationAsset := sparteraapi.Asset{
    AssetType:           sparteraapi.PtrString("VISUALIZATION"),
    VizChartType:        sparteraapi.PtrString("BAR"),
    VizDepVarColName:    sparteraapi.PtrString("sales_amount"),
    VizIndepVarColName:  sparteraapi.PtrString("month"),
}
```

### Pricing Strategies

Set different pricing models:

```go
// Fixed price per analysis
basicPricing := sparteraapi.Assetpricehistory{
    PriceUsd: sparteraapi.PtrFloat64(1.50),
}

// Premium analytics
premiumPricing := sparteraapi.Assetpricehistory{
    PriceUsd: sparteraapi.PtrFloat64(10.00),
}

// Bulk discount with sales
salePricing := sparteraapi.Assetpricehistory{
    PriceUsd:            sparteraapi.PtrFloat64(5.00),
    DiscountPercentage:  sparteraapi.PtrFloat32(20.0),
    SaleStartDate:       sparteraapi.PtrString("2024-01-01T00:00:00Z"),
    SaleEndDate:         sparteraapi.PtrString("2024-01-31T23:59:59Z"),
}
```

### Marketplace Management

Manage your products after launch:

```go
// Update asset details
updateData := sparteraapi.Asset{
    Description: sparteraapi.PtrString("Updated description with new features"),
}
_, _, err := client.AssetsApi.CompaniesCompanyIdAssetsAssetIdPatch(ctx, COMPANY_ID, assetID).Asset(updateData).Execute()
if err != nil {
    log.Printf("Failed to update asset: %v", err)
}

// Change pricing
newPrice := sparteraapi.Assetpricehistory{
    PriceUsd: sparteraapi.PtrFloat64(3.00),
}
_, _, err = client.AssetPriceHistoryApi.CompaniesCompanyIdAssetsAssetIdPricesPost(ctx, COMPANY_ID, assetID).Assetpricehistory(newPrice).Execute()
if err != nil {
    log.Printf("Failed to update price: %v", err)
}

// Remove from marketplace (but keep private)
marketplaceUpdate := sparteraapi.Asset{
    SellInMarketplace: sparteraapi.PtrString("false"),
}
_, _, err = client.AssetsApi.CompaniesCompanyIdAssetsAssetIdPatch(ctx, COMPANY_ID, assetID).Asset(marketplaceUpdate).Execute()
if err != nil {
    log.Printf("Failed to remove from marketplace: %v", err)
}

// Get sales analytics
analytics, _, err := client.CompaniesApi.CompaniesCompanyIdAnalyticsSalesGet(ctx, COMPANY_ID).Execute()
if err != nil {
    log.Printf("Failed to get analytics: %v", err)
} else {
    fmt.Printf("Total revenue: $%.2f\n", analytics.GetTotalRevenue())
}
```

### Error Handling

```go
import (
    "fmt"
    "net/http"
    sparteraapi "github.com/spartera-com/spartera-api-go-sdk"
)

func handleAPICall(ctx context.Context, client *sparteraapi.APIClient) error {
    asset, response, err := client.AssetsApi.CompaniesCompanyIdAssetsPost(ctx, COMPANY_ID).Asset(assetData).Execute()
    if err != nil {
        fmt.Printf("API Error: %v\n", err)
        
        // Handle specific HTTP status codes
        if response != nil {
            switch response.StatusCode {
            case http.StatusBadRequest:
                fmt.Println("Bad Request: Check your data format")
            case http.StatusUnauthorized:
                fmt.Println("Unauthorized: Check your API key")
            case http.StatusForbidden:
                fmt.Println("Forbidden: Check your permissions")
            case http.StatusNotFound:
                fmt.Println("Not Found: Check your IDs")
            default:
                fmt.Printf("HTTP %d: %v\n", response.StatusCode, err)
            }
        }
        return err
    }
    
    return nil
}
```

### Advanced Features

```go
// Batch operations
var connectionIDs []string
platforms := []string{"bigquery", "snowflake", "redshift"}

for _, platform := range platforms {
    conn, _, err := client.ConnectionsApi.CompaniesCompanyIdConnectionsPost(ctx, COMPANY_ID).Connection(platformConfig).Execute()
    if err != nil {
        return err
    }
    connectionIDs = append(connectionIDs, conn.GetConnectionId())
}

// Asset recommendations
recommendations, _, err := client.AssetsApi.CompaniesCompanyIdAssetsAssetIdRecommendationsGet(ctx, COMPANY_ID, assetID).Limit("10").Execute()
if err != nil {
    return err
}

// Performance analytics
performance, _, err := client.CompaniesApi.CompaniesCompanyIdAnalyticsAssetsGet(ctx, COMPANY_ID).StartDate("2024-01-01").EndDate("2024-12-31").Execute()
if err != nil {
    return err
}
```

## Package Manager
- **Platform**: Go Modules
- **Install**: `go get github.com/spartera-com/spartera-api-go-sdk`
- **Import**: `import sparteraapi "github.com/spartera-com/spartera-api-go-sdk"`

## Publishing
1. Tag: `git tag v1.0.0`
2. Push: `git push origin v1.0.0`
3. Install: `go get github.com/spartera-com/spartera-api-go-sdk@v1.0.0`