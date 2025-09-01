# 🔍 AWS Comprehend Demo - Text Analysis

A comprehensive demonstration of AWS Comprehend's natural language processing capabilities using the AWS SDK for Go v2. This project showcases various text analysis features including sentiment analysis, entity detection, language detection, and more.

## 🚀 Features

This demo demonstrates the following AWS Comprehend capabilities:

- **🌍 Language Detection** - Automatically detect the language of input text
- **😊 Sentiment Analysis** - Analyze emotional tone (positive, negative, neutral, mixed)
- **🏷️ Entity Detection** - Identify people, places, organizations, and other entities
- **🔑 Key Phrase Extraction** - Extract important phrases from text
- **🔒 PII Detection** - Identify personally identifiable information (emails, SSNs, phone numbers, etc.)
- **📝 Syntax Analysis** - Parse parts of speech and grammatical structure

## 📋 Prerequisites

- Go 1.23.1 or later
- AWS account with Comprehend service access
- AWS credentials configured (via AWS CLI, environment variables, or IAM roles)

## 🛠️ Installation

1. Clone this repository:
```bash
git clone <repository-url>
cd amazon-comprehend
```

2. Install dependencies:
```bash
go mod download
```

3. Configure AWS credentials using one of these methods:
   - AWS CLI: `aws configure`
   - Environment variables: `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`
   - IAM roles (if running on EC2)

## 🎯 Usage

Run the demo:
```bash
go run main.go
```

The program will analyze three sample texts and display comprehensive results for each AWS Comprehend feature.

## 📊 Sample Output

```
🔍 AWS Comprehend Demo - Text Analysis
=====================================

📝 Sample Text 1:
I love this new product! It's absolutely amazing and works perfectly. John Smith from New York called me at 555-123-4567.

🌍 Language Detection:
  Language: en (Confidence: 98.50%)

😊 Sentiment Analysis:
  Sentiment: POSITIVE
  Positive: 99.97% | Negative: 0.00% | Neutral: 0.03% | Mixed: 0.00%

🏷️  Entity Detection:
  PERSON: "John Smith" (Confidence: 99.92%)
  LOCATION: "New York" (Confidence: 99.76%)
  OTHER: "555-123-4567" (Confidence: 99.88%)

🔑 Key Phrases:
  "this new product" (Confidence: 100.00%)
  "John Smith" (Confidence: 99.99%)
  "New York" (Confidence: 100.00%)
  "555-123-4567" (Confidence: 99.57%)

🔒 PII Detection:
  NAME: "John Smith" (Confidence: 99.97%)
  ADDRESS: "New York" (Confidence: 97.84%)
  PHONE: "555-123-4567" (Confidence: 99.96%)

📝 Syntax Analysis (First 5 tokens):
  "I" -> PRON
  "love" -> VERB
  "this" -> DET
  "new" -> ADJ
  "product" -> NOUN
  ... and 19 more tokens
```

## 🏗️ Project Structure

```
aws-comprehend/
├── main.go           # Main application with all Comprehend demos
├── go.mod           # Go module definition
├── go.sum           # Go module checksums
└── README.md        # This file
```

## 📦 Dependencies

- `github.com/aws/aws-sdk-go-v2/config` - AWS SDK configuration
- `github.com/aws/aws-sdk-go-v2/service/comprehend` - AWS Comprehend service client

## 🔐 AWS Permissions

Your AWS credentials need the following Comprehend permissions:
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "comprehend:DetectDominantLanguage",
                "comprehend:DetectSentiment",
                "comprehend:DetectEntities",
                "comprehend:DetectKeyPhrases",
                "comprehend:DetectPiiEntities",
                "comprehend:DetectSyntax"
            ],
            "Resource": "*"
        }
    ]
}
```

## 🌟 Customization

To analyze your own text, modify the `sampleTexts` slice in `main.go`:

```go
var sampleTexts = []string{
    "Your custom text here...",
    "Another text to analyze...",
}
```
