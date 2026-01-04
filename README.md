# LLM Decision Loop Framework

A **Golang-based framework** that combines the reasoning capability of **Large Language Models (LLMs)** with a **deterministic rule engine** to produce **valid, controllable, and auditable decisions** through an iterative validation loop.

---

## ✨ Motivation

Large Language Models are powerful at reasoning and decision-making, but their outputs often suffer from:

- Violating business or compliance rules
- Unstable or non-deterministic formats
- Lack of strong guarantees and verifiability

This framework is built around a simple idea:

> **Let the LLM decide, let the rule engine validator.**

By validating LLM-generated decisions with a rule engine and feeding structured errors back into the model, the system enables **automatic self-correction** until a valid decision is produced.

---

## 🧠 Core Concept

1. LLM generates a decision
2. Decision is returned as structured JSON
3. Rule engine validates the decision
4. If invalid, structured error feedback is generated
5. Error feedback is sent back to the LLM
6. The loop continues until the decision is valid or a max iteration limit is reached

---

## 🏗️ Architecture

```text
+------------------------------------------------------+
|                  Decision Loop Engine                |
|                                                      |
|  +-------------------+       +-------------------+   |
|  |                   |       |                   |   |
|  |        LLM         |       |   Rule Engine     |   |
|  |   (Decision Maker)|       |   (Validator)     |   |
|  |                   |       |                   |   |
|  +---------+---------+       +---------+---------+   |
|            |                               ^          |
|            | Decision (JSON)               |          |
|            v                               |          |
|  +------------------------------------------------+  |
|  |                Loop Controller                 |  |
|  |                                                |  |
|  |  - Iteration control                           |  |
|  |  - JSON parsing & schema check                 |  |
|  |  - Error aggregation                           |  |
|  |  - Prompt construction                         |  |
|  +------------------------------------------------+  |
|            |                               ^          |
|            | Validation Errors (JSON)      |          |
|            +-------------------------------+          |
|                                                      |
+------------------------------------------------------+
```

---

## 🔁 Decision Loop

Start
↓
LLM generates decision (JSON)
↓
Rule engine validates
↓
Is decision valid?
├── Yes → Return final decision
└── No  → Return errors → Feed back to LLM → Next iteration

---

## 📦 Decision Format (Example)

### LLM Output

```json
{
  "action": "create_order",
  "amount": 1200,
  "currency": "CNY",
  "user_level": "vip"
}
```

---

### Rule Engine Validation Error

```json
{
  "error_code": "AMOUNT_LIMIT_EXCEEDED",
  "message": "Order amount exceeds the maximum allowed for the user level",
  "rule": "vip_user_max_amount_1000"
}
```

---

### Feedback to LLM (Prompt Example)

```text
The previous decision failed validation:

- Error Code: AMOUNT_LIMIT_EXCEEDED
- Rule: vip_user_max_amount_1000
- Message: Order amount exceeds the allowed maximum

Please generate a new decision JSON that satisfies all rules.
```

# 🚀 Features
- ✅ Clear separation between LLM reasoning and rule enforcement
-	✅ Deterministic validation with a rule engine
-	✅ Structured JSON-based communication
-	✅ Automatic multi-round self-correction
-	✅ Configurable maximum iteration limit
-	✅ Auditable and traceable decision process
-	✅ Easy integration with any LLM provider

# 🛠️ Usage Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/tu6ge/RefineGPT/engine"
	"github.com/tu6ge/RefineGPT/generator"
)

func main() {
	ctx := context.Background()

	// 1️⃣ State
	state := ExampleState{
		Task: "decide whether to allow access",
	}

	// 3️⃣ Generator
	gen := &generator.LLMGenerator{
		Client:  &MockLLM{},
		Adapter: generator.NewDefaultPromptAdapter(),
		Parser:  &CandidateFactory{},
		Schema: `
{
  "type": "object",
  "properties": {
    "action": {
      "type": "string",
      "enum": ["allow", "deny"]
    }
  },
  "required": ["action"]
}
`,
	}

	// 4️⃣ Engine
	e := &engine.Engine{
		Generator: gen,
		Validator: &ExampleValidator{},
		Policy: engine.LoopPolicy{
			MaxIteration: 5,
			StopOnFatal:  true,
		},
	}

	// 5️⃣ Run
	result, feedbacks, err := e.Run(ctx, state)
	if err != nil {
		panic(err)
	}

	fmt.Println("Final Candidate:", string(result.Raw()))
	fmt.Println("Feedback History:", feedbacks)
}
```