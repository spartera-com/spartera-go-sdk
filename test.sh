#!/bin/bash
# Test script for Go SDK

echo "🧪 Testing Go SDK..."

# Install dependencies
echo "Installing dependencies..."
go mod tidy

if [ $? -eq 0 ]; then
    echo "✅ Dependencies installed successfully"
    
    # Try to build our custom example
    echo "Building example..."
    cd examples
    go build -o spartera_example spartera_example.go
    
    if [ $? -eq 0 ]; then
        echo "✅ Example built successfully"
        echo "📝 To test with real API calls:"
        echo "   export SPARTERA_API_KEY='your-api-key'"
        echo "   export SPARTERA_COMPANY_ID='your-company-id'"
        echo "   cd examples && ./spartera_example"
        
        # Clean up the built example
        rm -f spartera_example
        cd ..
    else
        echo "❌ Example build failed"
        cd ..
        exit 1
    fi
else
    echo "❌ Dependencies installation failed"
    exit 1
fi
