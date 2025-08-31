package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/comprehend"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
)

var (
	ctx    context.Context
	client *comprehend.Client
)

// Sample texts to analyze
var sampleTexts = []string{
	"I love this new product! It's absolutely amazing and works perfectly. John Smith from New York called me at 555-123-4567.",
	"This service is terrible. I'm very disappointed and frustrated with the poor quality.",
	"The weather today is nice. My email is john.doe@example.com and my SSN is 123-45-6789.",
}

func main() {
	fmt.Println("🔍 AWS Comprehend Demo - Text Analysis")
	fmt.Println("=====================================\n")

	for i, text := range sampleTexts {
		fmt.Printf("📝 Sample Text %d:\n%s\n\n", i+1, text)

		// Analyze each feature
		detectLanguage(text)
		detectSentiment(text)
		detectEntities(text)
		detectKeyPhrases(text)
		detectPII(text)
		detectSyntax(text)

		fmt.Println("----------------------------------------\n")
	}
}

func init() {
	ctx = context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	client = comprehend.NewFromConfig(cfg)
}

// Detect language of the text
func detectLanguage(text string) {
	fmt.Println("🌍 Language Detection:")

	input := &comprehend.DetectDominantLanguageInput{
		Text: &text,
	}

	result, err := client.DetectDominantLanguage(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, lang := range result.Languages {
		fmt.Printf("  Language: %s (Confidence: %.2f%%)\n", *lang.LanguageCode, *lang.Score*100)
	}
	fmt.Println()
}

// Detect sentiment of the text
func detectSentiment(text string) {
	fmt.Println("😊 Sentiment Analysis:")

	input := &comprehend.DetectSentimentInput{
		Text:         &text,
		LanguageCode: types.LanguageCodeEn,
	}

	result, err := client.DetectSentiment(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("  Sentiment: %s\n", result.Sentiment)
	fmt.Printf("  Positive: %.2f%% | Negative: %.2f%% | Neutral: %.2f%% | Mixed: %.2f%%\n",
		*result.SentimentScore.Positive*100,
		*result.SentimentScore.Negative*100,
		*result.SentimentScore.Neutral*100,
		*result.SentimentScore.Mixed*100)
	fmt.Println()
}

// Detect entities in the text
func detectEntities(text string) {
	fmt.Println("🏷️  Entity Detection:")

	input := &comprehend.DetectEntitiesInput{
		Text:         &text,
		LanguageCode: types.LanguageCodeEn,
	}

	result, err := client.DetectEntities(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Entities) == 0 {
		fmt.Println("  No entities found")
	} else {
		for _, entity := range result.Entities {
			fmt.Printf("  %s: \"%s\" (Confidence: %.2f%%)\n",
				entity.Type, *entity.Text, *entity.Score*100)
		}
	}
	fmt.Println()
}

// Detect key phrases in the text
func detectKeyPhrases(text string) {
	fmt.Println("🔑 Key Phrases:")

	input := &comprehend.DetectKeyPhrasesInput{
		Text:         &text,
		LanguageCode: types.LanguageCodeEn,
	}

	result, err := client.DetectKeyPhrases(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.KeyPhrases) == 0 {
		fmt.Println("  No key phrases found")
	} else {
		for _, phrase := range result.KeyPhrases {
			fmt.Printf("  \"%s\" (Confidence: %.2f%%)\n", *phrase.Text, *phrase.Score*100)
		}
	}
	fmt.Println()
}

// Detect PII (Personally Identifiable Information)
func detectPII(text string) {
	fmt.Println("🔒 PII Detection:")

	input := &comprehend.DetectPiiEntitiesInput{
		Text:         &text,
		LanguageCode: types.LanguageCodeEn,
	}

	result, err := client.DetectPiiEntities(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(result.Entities) == 0 {
		fmt.Println("  No PII found")
	} else {
		for _, entity := range result.Entities {
			fmt.Printf("  %s: \"%s\" (Confidence: %.2f%%)\n",
				entity.Type, text[*entity.BeginOffset:*entity.EndOffset], *entity.Score*100)
		}
	}
	fmt.Println()
}

// Detect syntax and parts of speech
func detectSyntax(text string) {
	fmt.Println("📝 Syntax Analysis (First 5 tokens):")

	input := &comprehend.DetectSyntaxInput{
		Text:         &text,
		LanguageCode: types.SyntaxLanguageCode(types.LanguageCodeEn),
	}

	result, err := client.DetectSyntax(ctx, input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Show only first 5 tokens to keep output clean
	maxTokens := 5
	if len(result.SyntaxTokens) < maxTokens {
		maxTokens = len(result.SyntaxTokens)
	}

	for i := 0; i < maxTokens; i++ {
		token := result.SyntaxTokens[i]
		fmt.Printf("  \"%s\" -> %s\n", *token.Text, token.PartOfSpeech.Tag)
	}

	if len(result.SyntaxTokens) > 5 {
		fmt.Printf("  ... and %d more tokens\n", len(result.SyntaxTokens)-5)
	}
	fmt.Println()
}
