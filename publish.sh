#!/bin/bash
# =============================================================================
# Go SDK Publisher
# =============================================================================
# 
# PREREQUISITES:
# 1. Push code to GitHub: git push origin main
# 2. Create release tag: git tag v1.0.0 && git push origin v1.0.0
# 
# Go modules are automatically available once pushed to GitHub with proper tags.
# No separate package manager registration required.
# =============================================================================

echo "🐹 Preparing Go SDK for publishing..."

# Check if we're in a git repository
if [ ! -d ".git" ]; then
    echo "❌ Not in a git repository. Initialize with: git init"
    exit 1
fi

# Tidy up dependencies
echo "📚 Tidying dependencies..."
go mod tidy

if [ $? -ne 0 ]; then
    echo "❌ go mod tidy failed"
    exit 1
fi

# Run tests if available
echo "🧪 Running tests..."
if go test ./... >/dev/null 2>&1; then
    echo "✅ Tests passed"
else
    echo "⚠️  Tests failed or no tests found"
fi

# Verify module
echo "🔍 Verifying module..."
go mod verify

if [ $? -ne 0 ]; then
    echo "❌ Module verification failed"
    exit 1
fi

# Get version from config or use default
VERSION="v$API_VERSION"
if [ -z "$API_VERSION" ]; then
    VERSION="v1.0.0"
fi

echo ""
echo "📋 To publish Go SDK:"
echo "1. Commit all changes: git add . && git commit -m 'Release $VERSION'"
echo "2. Push to GitHub: git push origin main"
echo "3. Create version tag: git tag $VERSION"
echo "4. Push tag: git push origin $VERSION"
echo "5. SDK will be available as: go get github.com/spartera-com/spartera-go-sdk@$VERSION"
echo ""

read -p "🚀 Create and push git tag $VERSION? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    # Check if we have uncommitted changes
    if ! git diff-index --quiet HEAD --; then
        echo "⚠️  You have uncommitted changes. Committing them first..."
        git add .
        git commit -m "Release $VERSION"
    fi
    
    # Push current branch
    echo "📤 Pushing to GitHub..."
    git push origin $(git branch --show-current)
    
    # Create and push tag
    echo "🏷️  Creating tag $VERSION..."
    git tag "$VERSION"
    git push origin "$VERSION"
    
    echo "✅ Tag created and pushed!"
    echo "🎉 Go SDK is now available at:"
    echo "   go get github.com/spartera-com/spartera-go-sdk@$VERSION"
    echo ""
    echo "🔗 Users can import it as:"
    echo "   import sparteraapi \"github.com/spartera-com/spartera-go-sdk\""
else
    echo "❌ Tag creation cancelled"
    echo ""
    echo "💡 Manual steps:"
    echo "   git tag $VERSION"
    echo "   git push origin $VERSION"
fi
