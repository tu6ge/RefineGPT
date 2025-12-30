# Example

This example demonstrates a full loop:

LLM (mock) → JSON candidate → validator → feedback → retry

## What it shows

- How to implement a business State
- How to implement a Validator
- How to plug in a Generator + PromptAdapter
- How Engine orchestrates everything

Run:

```bash
go run ./example