# CDK DynamoDB Table (Go + AWS CDK)

This project is a small Infrastructure as Code (IaC) example built with the AWS Cloud Development Kit (CDK) using Go.

It defines and deploys a DynamoDB table using AWS CDK constructs, demonstrating how cloud resources can be created programmatically instead of manually in the AWS console.

---

## What This Project Does

1. Defines an AWS CDK stack in Go
2. Creates a DynamoDB table
3. Uses on-demand (pay-per-request) billing
4. Synthesizes the CloudFormation template via CDK

The project focuses only on table creation to keep the example simple and easy to understand.

---

## Why This Exists

This project was built to practice:

* Infrastructure as Code (IaC) concepts
* Using AWS CDK with Go
* Defining AWS resources programmatically
* Creating DynamoDB tables with best-practice defaults

It demonstrates how application developers can manage cloud infrastructure alongside application code.

---

## Project Structure

```
cdk-dynamodb-table
├── main/
│   └── main.go          # CDK app entry point
├── cdktable.go          # Stack definition and DynamoDB table
├── cdktable_test.go     # Basic CDK test
├── cdk.json             # CDK configuration
├── cdk.context.json     # CDK context data
├── cdk.out/             # Synthesized CloudFormation output
├── stacks.csv           # Stack name reference
├── go.mod
└── go.sum
```

---

## DynamoDB Table Configuration

The stack creates a DynamoDB table with the following settings:

* **Table name:** `barjokes-cdk`
* **Partition key:** `NAME` (string)
* **Billing mode:** Pay-per-request (on-demand)

This configuration avoids capacity planning and is suitable for variable or unpredictable workloads.

---

## How It Works (High Level)

* The CDK app is initialized in `main.go`
* A stack is defined in `cdktable.go`
* The DynamoDB table is created using CDK constructs
* `cdk synth` generates a CloudFormation template
* `cdk deploy` can be used to deploy the stack to AWS

---

## Requirements

* Go installed
* AWS CDK installed
* AWS credentials configured locally

---

## How to Use

Synthesize the CloudFormation template:

```bash
cdk synth
```

Deploy the stack:

```bash
cdk deploy
```

---

## Notes

* The AWS environment (account and region) can be configured in `main.go`
* This project focuses on infrastructure definition, not application logic
* The example is intentionally minimal for learning purposes

---

## Purpose of This Project

This project is intentionally small.

It is meant to demonstrate:

* Infrastructure as Code using AWS CDK
* DynamoDB table creation
* Go-based cloud infrastructure definitions

It is not intended to be a full production infrastructure setup.
