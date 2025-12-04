# OpenRouter API Integration Checklist

## 📋 Overview
This checklist is designed for LLM providers preparing to integrate with OpenRouter.  Complete all sections to ensure compliance with OpenRouter's API requirements. 

---

## 1. Company & Provider Information

### Basic Information
- [ ] Company name provided
- [ ] Company email provided (will be added to Slack Connect channel)
- [ ] Privacy Policy URL published and accessible

### Distinguishing Features
Select and document what distinguishes your service:
- [ ] Low Latency
- [ ] High Throughput
- [ ] Unique Models (list models not available elsewhere)
- [ ] Low Pricing (provide comparative pricing)
- [ ] Volume/Committed Discounts (specify discount structure)
- [ ] Unique Infrastructure (describe technical advantages)
- [ ] Decentralized Architecture
- [ ] Strategic Partnership Opportunities
- [ ] Unusually high performance on specific models
- [ ] Other architectural innovations (hardware, caching, etc.)

**Extra Details:**
```
[Document your unique value proposition and technical advantages]
```

---

## 2. Core API Endpoints (REQUIRED)

### 2.1 POST /chat/completions
- [ ] **Endpoint URL provided:** `_______________________________`
- [ ] OpenAI-compliant API format
- [ ] Returns `usage` object for **non-stream** requests
- [ ] Returns `usage` object for **stream** requests (CRITICAL)
- [ ] Returns cache read/write usage (if caching available)
- [ ] Supports streaming (`stream: true`)
- [ ] Supports non-streaming (`stream: false`)
- [ ] Tested with various payload sizes
- [ ] Error handling implemented for malformed requests

**Testing Checklist:**
- [ ] Successfully returns responses for simple queries
- [ ] Streaming works with SSE (Server-Sent Events) format
- [ ] `data: [DONE]` signal sent at stream end
- [ ] Usage tokens returned in final stream chunk
- [ ] Non-streaming returns complete response with usage

---

### 2.2 POST /completions
- [ ] **Endpoint URL provided:** `_______________________________`
- [ ] OpenAI-compliant API format
- [ ] Returns `usage` object for **non-stream** requests
- [ ] Returns `usage` object for **stream** requests (CRITICAL)
- [ ] Returns cache read/write usage (if caching available)
- [ ] Supports streaming (`stream: true`)
- [ ] Supports non-streaming (`stream: false`)
- [ ] Tested with text completion scenarios

---

### 2.3 GET /models
- [ ] **Endpoint URL provided:** `_______________________________`
- [ ] Returns up-to-date model information
- [ ] Schema similar to [OpenAI models endpoint](https://platform.openai.com/docs/api-reference/models)
- [ ] Shows **maximum output tokens** per model
- [ ] Shows **context length** per model
- [ ] Includes model IDs matching what's used in completions endpoints
- [ ] Returns pricing information per model (if applicable)
- [ ] Updates dynamically when models are added/removed

**Required Fields per Model:**
- [ ] `id` (model identifier)
- [ ] `object` (should be "model")
- [ ] `created` (timestamp)
- [ ] `owned_by` (your company name)
- [ ] `context_length` (integer)
- [ ] `max_output_tokens` (integer)

---

## 3. Tokenization

- [ ] **Tokenizer used:** `_______________________________`
- [ ] Token counting method documented
- [ ] Consistent with reported usage in API responses
- [ ] Handles special tokens appropriately
- [ ] Documented edge cases (if any)

**Documentation:**
```
[Explain how you count tokens, which tokenizer you use, and any specifics]
```

---

## 4. Required Parameters Support

All endpoints must support these parameters:

- [ ] `max_tokens` - Maximum tokens to generate
- [ ] `temperature` - Sampling temperature (0. 0 to 2.0 typical range)
- [ ] `top_p` - Nucleus sampling parameter
- [ ] `stop` - Stop sequences (string or array)
- [ ] `seed` - Random seed for reproducibility

**Parameter Testing:**
- [ ] Each parameter tested individually
- [ ] Parameters tested in combination
- [ ] Default values documented
- [ ] Range limits documented

---

## 5. Optional Parameters Support

Mark which optional parameters you support:

### Sampling Parameters
- [ ] `top_k` - Top-K sampling
- [ ] `frequency_penalty` - Penalize frequent tokens
- [ ] `presence_penalty` - Penalize present tokens
- [ ] `repetition_penalty` - Alternative repetition control
- [ ] `min_p` - Minimum probability threshold
- [ ] `top_a` - Top-A sampling

### Response Control
- [ ] `logit_bias` - Modify token probabilities
- [ ] `logprobs` - Return log probabilities
- [ ] `top_logprobs` - Number of top logprobs to return

### Structured Outputs
- [ ] `response_format` - Control output format (e.g., JSON)
- [ ] `structured_outputs` - JSON schema validation

### Tool Calling
- [ ] `tools` - Function/tool definitions
- [ ] `tool_choice` - Control tool selection

**Additional Parameters:**
```
[List any extra sampling or control parameters you support]
```

---

## 6. Tool Calling (If Supported)

- [ ] Tool calling supported
- [ ] Works with **streaming** requests
- [ ] Well-tested and production-ready
  - [ ] Production-ready
  - [ ] Beta (experimental)
- [ ] Follows OpenAI function calling schema
- [ ] Returns tool calls in stream chunks
- [ ] Handles multiple tool calls in single request
- [ ] Handles parallel tool calls

**Testing Status:**
```
[Describe testing coverage and known limitations]
```

---

## 7. Structured Outputs (If Supported)

- [ ] JSON schema support implemented
- [ ] Works with **streaming** requests
- [ ] Always returns valid JSON matching schema
- [ ] Well-tested and production-ready
  - [ ] Production-ready
  - [ ] Beta (experimental)
- [ ] Handles complex nested schemas
- [ ] Provides validation errors when schema violated

**Schema Support Details:**
```
[Describe JSON schema capabilities and limitations]
```

---

## 8. Multi-Modal Support (If Applicable)

### Supported Image Types
- [ ] `image/jpeg` (JPEG)
- [ ] `image/png` (PNG)
- [ ] `image/gif` (GIF)
- [ ] `image/webp` (WebP)
- [ ] Other: `_______________________________`

### Image Input Methods
- [ ] Base64 encoded images
- [ ] Image URLs
- [ ] Maximum image size: `_____ MB`
- [ ] Maximum images per request: `_____`

### Multi-Modal Models
```
[List models that support vision/multi-modal inputs]
```

---

## 9. Pricing & Payment

### Pricing Transparency
- [ ] Publicly available pricing in USD per M tokens
- [ ] Pricing broken down by:
  - [ ] Input tokens
  - [ ] Output tokens
  - [ ] Cache read tokens (if applicable)
  - [ ] Cache write tokens (if applicable)
- [ ] Pricing matches what's returned in `/models` endpoint

### Volume Discounts
- [ ] Volume discounts available
- [ ] Discount structure documented:
```
[Describe volume tiers and discount percentages]
```

### Payment Methods
- [ ] Credit card payments accepted
- [ ] Automated payment supported (auto top-up)
- [ ] Invoicing available
- [ ] Payment frequency: `_______________________________`

**Invoicing Details:**
```
[Describe invoicing process and payment terms]
```

---

## 10. Rate Limits

- [ ] **Initial rate limits documented:**
  - Requests per minute: `_____`
  - Requests per day: `_____`
  - Tokens per minute: `_____`
  - Concurrent requests: `_____`

- [ ] Rate limits can be raised
- [ ] Process for raising limits documented:
```
[Describe how OpenRouter can request rate limit increases]
```

- [ ] Rate limit headers returned in responses:
  - [ ] `x-ratelimit-limit-requests`
  - [ ] `x-ratelimit-remaining-requests`
  - [ ] `x-ratelimit-reset-requests`

---

## 11. Failure States & Billing

### Mid-Request Cancellations
When clients disconnect/abort mid-stream:
- [ ] **Do you charge? **
  - [ ] Yes, we charge for tokens generated up to cancellation
  - [ ] No, we do not charge for cancelled requests

**Cancellation Policy:**
```
[Explain your cancellation billing policy in detail]
```

---

### Model/Engine Failures
When models fail or engine errors occur:
- [ ] **Do you charge?**
  - [ ] Yes, we charge even on failures
  - [ ] No, we do not charge for failed requests

**Failure Billing Policy:**
```
[Explain when you do/don't charge for failures]
```

---

### Error Response Format
- [ ] OpenAI-compliant error format
- [ ] Includes `error. message`
- [ ] Includes `error.type`
- [ ] Includes `error.code` (if applicable)

**Special Error Shapes:**
```
[Document any custom error codes or finish_reason values]
```

**Custom Finish Reasons:**
- [ ] `stop` - Natural completion
- [ ] `length` - Max tokens reached
- [ ] `content_filter` - Content filtered
- [ ] `tool_calls` - Tool calling completed
- [ ] Other: `_______________________________`

---

## 12. Data Privacy & Retention

### Training Policy
- [ ] **Do you train on prompts? **
  - [ ] Yes
  - [ ] No

- [ ] **Do you train on completions?**
  - [ ] Yes
  - [ ] No

### Data Retention
- [ ] Logging policy documented
- [ ] Retention period: `_______________________________`
- [ ] Data deletion process available
- [ ] Compliance certifications:
  - [ ] GDPR compliant
  - [ ] SOC 2
  - [ ] HIPAA
  - [ ] Other: `_______________________________`

**Data Policy Details:**
```
[Provide comprehensive data handling and privacy policy]
```

---

## 13. Infrastructure & Location

### Inference Locations
Document country codes where inference will be served:
- [ ] `_______________________________`
- [ ] `_______________________________`
- [ ] `_______________________________`

### Performance Considerations
- [ ] Can handle streaming-only requests from OpenRouter
- [ ] Can handle requirement to always return `stream_options. include_usage: true`
- [ ] No latency issues with streaming requirements
- [ ] No QoS degradation expected

**Performance Notes:**
```
[Document any concerns about streaming performance or latency]
```

---

## 14. Testing & Validation

### Basic Integration Tests
- [ ] `/chat/completions` non-stream request successful
- [ ] `/chat/completions` stream request successful
- [ ] `/completions` non-stream request successful
- [ ] `/completions` stream request successful
- [ ] `/models` returns valid data
- [ ] Usage tokens returned in all responses

### Streaming Tests
- [ ] Stream chunks arrive incrementally
- [ ] `data: [DONE]` signal received
- [ ] Usage included in stream (with `stream_options.include_usage: true`)
- [ ] Connection handles long responses (10K+ tokens)
- [ ] Cancellation mid-stream works gracefully

### Error Handling Tests
- [ ] Invalid API key returns 401
- [ ] Invalid model returns 404
- [ ] Malformed JSON returns 400
- [ ] Rate limit returns 429
- [ ] Server errors return 500+ with details

### Load Tests
- [ ] Sustained concurrent requests handled
- [ ] Rate limits enforced correctly
- [ ] No degradation at advertised throughput

---

## 15. Final Confirmation

- [ ] All required endpoints implemented and tested
- [ ] All required parameters supported
- [ ] Usage tokens always returned (stream and non-stream)
- [ ] OpenAI compatibility verified
- [ ] Documentation complete and accurate
- [ ] Privacy policy published
- [ ] Pricing published
- [ ] Payment method established
- [ ] **Confirmed: We can handle streaming-only requests**
- [ ] **Confirmed: We always return `stream_options.include_usage: true`**

---

## 16. Submission Notes

**Final Notes for OpenRouter Team:**
```
[Add any additional context, concerns, or information about latency/QoS
with streaming requirements, special considerations, or unique features]
```

---

## Contact & Support

- **Technical Contact:** `_______________________________`
- **Support Email:** `_______________________________`
- **Documentation URL:** `_______________________________`
- **Status Page URL:** `_______________________________`
- **Slack Channel:** (Will be created by OpenRouter)

---

## Appendix: OpenAI Compatibility Reference

For detailed parameter specifications, refer to:
- Chat Completions: https://platform. openai.com/docs/api-reference/chat
- Completions: https://platform.openai.com/docs/api-reference/completions
- Models: https://platform.openai.com/docs/api-reference/models

---

**Version:** 1.0  
**Last Updated:** 2025-12-04  
**Based on:** OpenRouter Provider Onboarding Questionnaire
