# Text Moderation Package

This package provides a simple solution for text moderation, leveraging Amazon Comprehend to detect sentiment and filtering explicit content based on keywords. The package allows for a combination of keyword-based moderation and sentiment analysis to flag inappropriate or negative content.

## Features

- **Explicit Terms Filtering**: Add and manage a list of explicit terms to flag inappropriate content.
- **Sentiment Analysis**: Analyze text sentiment using Amazon Comprehend, flagging text based on a negative sentiment threshold.
- **Moderation**: Perform comprehensive text moderation to flag texts as explicit or negative.

## Installation

Ensure you have the AWS SDK for Go installed. You can get it with:

```bash
go get -u github.com/aws/aws-sdk-go
```

Add the `textmoderation` package to your project.

## Usage

### 1. Initialize Amazon Comprehend with Credentials

To use the sentiment analysis feature, initialize Amazon Comprehend by specifying your AWS `accessKey`, `secretKey`, and `region` directly. This avoids the need for environment variables and allows credentials to be passed programmatically.

```go
accessKey := "your-access-key"
secretKey := "your-secret-key"
region := "your-region"

comprehendClient, err := textmoderation.InitComprehend(accessKey, secretKey, region)
if err != nil {
    log.Fatal(err)
}
```

> **Note:** Be sure to store AWS credentials securely and avoid hardcoding them directly in your code.

### 2. Adding Explicit Terms

Add explicit terms for keyword-based filtering. You can add individual terms or multiple terms at once.

```go
// Add a single explicit term
textmoderation.AddExplicitTerm("explicitTerm")

// Add multiple explicit terms
textmoderation.AddExplicitTerms([]string{"term1", "term2", "term3"})
```

### 3. Check for Explicit Content

You can use the `CheckForExplicitContent` function to see if any explicit terms are present in a given text:

```go
text := "Some text containing explicitTerm."
if textmoderation.CheckForExplicitContent(text) {
    fmt.Println("Explicit content detected")
}
```

### 4. Sentiment Analysis

Analyze the sentiment of a given text using Amazon Comprehend:

```go
sentiment, negativeScore, err := textmoderation.AnalyzeSentiment(comprehendClient, "This is a sample text.")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Sentiment: %s, Negative Score: %.2f\n", sentiment, negativeScore)
```

### 5. Moderate Text

The `ModerateText` function performs comprehensive moderation, combining explicit content filtering and sentiment analysis. Text is flagged based on detected explicit terms or if the sentiment is negative with a high threshold.

```go
result, err := textmoderation.ModerateText(comprehendClient, "Some text to moderate.")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Moderation result:", result)  // Outputs: "approved", "flagged-explicit", or "flagged-negative"
```

## Moderation Workflow

1. **Explicit Content Check**: The text is first checked for explicit terms.
2. **Sentiment Analysis**: If no explicit content is detected, the sentiment is analyzed. If the sentiment is "NEGATIVE" with a score above 0.7, it is flagged as negative.
3. **Approval**: Text that passes both checks is marked as "approved."

## Configuration

- **Negative Sentiment Threshold**: The current threshold for flagging negative sentiment is 0.7, adjustable within the `ModerateText` function as needed.

## Requirements

- AWS SDK for Go
- Amazon Comprehend credentials and permissions to analyze sentiment

## Security Best Practices

- **Avoid Hardcoding Credentials**: Instead of hardcoding, consider loading credentials securely from a configuration file or using AWS Secrets Manager.
- **Rotate Access Keys**: Regularly rotate AWS access keys to maintain security.
- **Restrict Permissions**: Use the least-privilege principle for the IAM user or role associated with the access keys.

## License

This package is available under the [MIT License](LICENSE). See the `LICENSE` file for details.


```go
// Example:
package main

import (
	"fmt"
	"log"

	textmoderation "github.com/saif-eps/aws-textmoderation"
)

func main() {
	accessKey := ""
	secretKey := ""
	region := "us-west-2"

	// Initialize the Comprehend client
	comprehendClient, err := textmoderation.InitComprehend(accessKey, secretKey, region)
	if err != nil {
		log.Fatal("Failed to initialize Comprehend client:", err)
	}

	// Example text for moderation
	text := "This is some example text that might contain explicit content or negative sentiment."

	// Perform moderation
	result, err := textmoderation.ModerateText(comprehendClient, text)
	if err != nil {
		log.Fatal("Moderation failed:", err)
	}

	fmt.Println("Moderation result:", result) // Outputs: "approved", "flagged-explicit", or "flagged-negative"
}      
```