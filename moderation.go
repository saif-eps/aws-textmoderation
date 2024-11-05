package textmoderation

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

// explicitTerms holds the list of explicit terms to filter, initially empty
var explicitTerms = []string{}

// InitComprehend initializes the Amazon Comprehend client with explicit credentials and region
func InitComprehend(accessKey, secretKey, region string) (*comprehend.Comprehend, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		return nil, err
	}
	return comprehend.New(sess), nil
}

// AddExplicitTerm adds a new term to the explicit terms list
func AddExplicitTerm(term string) {
	explicitTerms = append(explicitTerms, term)
}

// AddExplicitTerms allows adding multiple explicit terms at once
func AddExplicitTerms(terms []string) {
	explicitTerms = append(explicitTerms, terms...)
}

// CheckForExplicitContent checks for explicit content based on keywords
func CheckForExplicitContent(text string) bool {
	textLower := strings.ToLower(text)
	for _, term := range explicitTerms {
		if strings.Contains(textLower, term) {
			return true
		}
	}
	return false
}

// AnalyzeSentiment analyzes the sentiment using Amazon Comprehend
func AnalyzeSentiment(svc *comprehend.Comprehend, text string) (string, float64, error) {
	input := &comprehend.DetectSentimentInput{
		Text:         aws.String(text),
		LanguageCode: aws.String("en"),
	}

	result, err := svc.DetectSentiment(input)
	if err != nil {
		return "", 0, err
	}

	// Get the sentiment score
	sentiment := *result.Sentiment
	negativeScore := *result.SentimentScore.Negative
	return sentiment, negativeScore, nil
}

// ModerateText performs moderation by checking for explicit content and analyzing sentiment
func ModerateText(svc *comprehend.Comprehend, text string) (string, error) {
	// Step 1: Check for explicit content
	if CheckForExplicitContent(text) {
		return "flagged-explicit", nil
	}

	// Step 2: Analyze sentiment
	sentiment, negativeScore, err := AnalyzeSentiment(svc, text)
	if err != nil {
		return "", err
	}

	// Flag content based on negative sentiment threshold
	if sentiment == "NEGATIVE" && negativeScore > 0.7 {
		return "flagged-negative", nil
	}

	return "approved", nil
}
