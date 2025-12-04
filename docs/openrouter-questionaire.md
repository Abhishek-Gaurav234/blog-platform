I'll elaborate on the API Endpoint Requirements section with more technical details:

```markdown name=OpenRouter-API-Integration-Checklist. md
## 3. API Endpoint Requirements

### Endpoint Configuration

#### Base URL & Routing
- [ ] **Stable Base URL**
  - [ ] Production endpoint is stable and versioned (e.g., `https://api.yourservice.com/v1`)
  - [ ] No breaking changes without version increment
  - [ ] Endpoint has >99. 9% uptime SLA
  - [ ] CDN/load balancing configured for reliability
  - [ ] Geographic distribution for low latency

- [ ] **Chat Completions Endpoint**
  - [ ] Primary endpoint: `/chat/completions` or equivalent
  - [ ] Accepts POST requests
  - [ ] Content-Type: `application/json`
  - [ ] Returns JSON responses
  - [ ] Supports both streaming and non-streaming modes

#### Example Request Structure
```json
POST /v1/chat/completions
Content-Type: application/json
Authorization: Bearer YOUR_API_KEY

{
  "model": "your-model-name",
  "messages": [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Hello!"}
  ],
  "temperature": 0.7,
  "max_tokens": 150,
  "stream": false
}
```

---

### Authentication

#### API Key Implementation
- [ ] **Authentication Method**
  - [ ] Bearer token authentication in Authorization header
  - [ ] Alternative: `api-key` header support
  - [ ] API keys are UUIDs or cryptographically secure random strings
  - [ ] Minimum key length: 32 characters

- [ ] **Key Management**
  - [ ] Key generation endpoint or dashboard
  - [ ] Key revocation capability
  - [ ] Multiple keys per account support
  - [ ] Key usage tracking and analytics
  - [ ] Key rotation policy documented

- [ ] **Security Best Practices**
  - [ ] HTTPS/TLS 1.2+ enforced (no HTTP)
  - [ ] Keys encrypted at rest
  - [ ] Keys never logged or exposed in responses
  - [ ] IP whitelisting option (optional but recommended)
  - [ ] Request signing support (optional)

- [ ] **Authentication Errors**
  - [ ] `401 Unauthorized` for missing/invalid keys
  - [ ] Clear error messages:
    ```json
    {
      "error": {
        "message": "Invalid API key provided",
        "type": "invalid_request_error",
        "code": "invalid_api_key"
      }
    }
    ```

---

### Rate Limiting

#### Rate Limit Configuration
- [ ] **Limits Defined**
  - [ ] Requests per minute (RPM) limit specified
  - [ ] Tokens per minute (TPM) limit specified
  - [ ] Requests per day (RPD) limit (if applicable)
  - [ ] Different tiers for different pricing plans
  - [ ] Burst allowance documented

- [ ] **Rate Limit Headers**
  - [ ] `X-RateLimit-Limit-Requests` - Total requests allowed
  - [ ] `X-RateLimit-Limit-Tokens` - Total tokens allowed
  - [ ] `X-RateLimit-Remaining-Requests` - Requests remaining
  - [ ] `X-RateLimit-Remaining-Tokens` - Tokens remaining
  - [ ] `X-RateLimit-Reset-Requests` - Time when limit resets (Unix timestamp)
  - [ ] `X-RateLimit-Reset-Tokens` - Token limit reset time

#### Example Response Headers
```
X-RateLimit-Limit-Requests: 3500
X-RateLimit-Limit-Tokens: 90000
X-RateLimit-Remaining-Requests: 3499
X-RateLimit-Remaining-Tokens: 89800
X-RateLimit-Reset-Requests: 1733356800
X-RateLimit-Reset-Tokens: 1733356800
```

- [ ] **Rate Limit Exceeded Response**
  - [ ] HTTP Status: `429 Too Many Requests`
  - [ ] `Retry-After` header with seconds until reset
  - [ ] Clear error message:
    ```json
    {
      "error": {
        "message": "Rate limit exceeded. Please retry after 45 seconds.",
        "type": "rate_limit_error",
        "code": "rate_limit_exceeded"
      }
    }
    ```

- [ ] **Graceful Degradation**
  - [ ] Soft limits with warnings before hard limits
  - [ ] Queue requests when possible instead of rejecting
  - [ ] Priority handling for different request types

---

### Error Handling

#### Standard HTTP Status Codes
- [ ] **Success Codes**
  - [ ] `200 OK` - Successful non-streaming response
  - [ ] `201 Created` - Resource created (if applicable)

- [ ] **Client Error Codes (4xx)**
  - [ ] `400 Bad Request` - Malformed request, invalid parameters
  - [ ] `401 Unauthorized` - Missing or invalid authentication
  - [ ] `403 Forbidden` - Valid auth but insufficient permissions
  - [ ] `404 Not Found` - Model or endpoint not found
  - [ ] `413 Payload Too Large` - Request exceeds size limits
  - [ ] `422 Unprocessable Entity` - Validation errors
  - [ ] `429 Too Many Requests` - Rate limit exceeded

- [ ] **Server Error Codes (5xx)**
  - [ ] `500 Internal Server Error` - Unexpected server error
  - [ ] `502 Bad Gateway` - Upstream service error
  - [ ] `503 Service Unavailable` - Temporary downtime
  - [ ] `504 Gateway Timeout` - Request timeout

#### Error Response Format
- [ ] **Consistent Error Structure**
  ```json
  {
    "error": {
      "message": "Human-readable error description",
      "type": "error_category",
      "code": "specific_error_code",
      "param": "problematic_parameter"
    }
  }
  ```

- [ ] **Error Types**
  - [ ] `invalid_request_error` - Bad request format/parameters
  - [ ] `authentication_error` - Auth issues
  - [ ] `rate_limit_error` - Rate limits exceeded
  - [ ] `api_error` - Server-side errors
  - [ ] `timeout_error` - Request timeout
  - [ ] `context_length_exceeded` - Prompt too long

#### Detailed Error Examples

**Invalid Parameter Error:**
```json
{
  "error": {
    "message": "Invalid value for 'temperature': must be between 0 and 2",
    "type": "invalid_request_error",
    "code": "invalid_parameter_value",
    "param": "temperature"
  }
}
```

**Context Length Exceeded:**
```json
{
  "error": {
    "message": "This model's maximum context length is 4096 tokens. Your messages resulted in 5234 tokens.",
    "type": "invalid_request_error",
    "code": "context_length_exceeded",
    "param": "messages"
  }
}
```

**Model Not Found:**
```json
{
  "error": {
    "message": "The model 'invalid-model-name' does not exist",
    "type": "invalid_request_error",
    "code": "model_not_found",
    "param": "model"
  }
}
```

**Service Unavailable:**
```json
{
  "error": {
    "message": "The server is temporarily unable to handle the request.  Please try again later.",
    "type": "api_error",
    "code": "service_unavailable"
  }
}
```

---

### Request Validation

- [ ] **Input Validation**
  - [ ] Validate all required fields present
  - [ ] Check parameter types (string, number, boolean, array)
  - [ ] Validate parameter ranges (e.g., temperature 0-2)
  - [ ] Sanitize inputs to prevent injection attacks
  - [ ] Validate message structure and roles
  - [ ] Check total token count before processing

- [ ] **Request Size Limits**
  - [ ] Maximum request body size defined (e.g., 10MB)
  - [ ] Maximum number of messages per request
  - [ ] Maximum tokens per message
  - [ ] Maximum total context length enforced

---

### Response Format

#### Non-Streaming Response
- [ ] **Standard Response Structure**
  ```json
  {
    "id": "chatcmpl-123abc",
    "object": "chat.completion",
    "created": 1733356800,
    "model": "your-model-name",
    "choices": [
      {
        "index": 0,
        "message": {
          "role": "assistant",
          "content": "Response text here"
        },
        "finish_reason": "stop"
      }
    ],
    "usage": {
      "prompt_tokens": 20,
      "completion_tokens": 50,
      "total_tokens": 70
    }
  }
  ```

- [ ] **Required Fields**
  - [ ] `id` - Unique identifier for the completion
  - [ ] `object` - Object type (e.g., "chat.completion")
  - [ ] `created` - Unix timestamp
  - [ ] `model` - Model used for generation
  - [ ] `choices` - Array of completion choices
  - [ ] `usage` - Token usage statistics

- [ ] **Finish Reasons**
  - [ ] `stop` - Natural completion
  - [ ] `length` - Max tokens reached
  - [ ] `content_filter` - Content filtered
  - [ ] `function_call` - Function call triggered (if supported)

#### Streaming Response
- [ ] **SSE Format**
  - [ ] Content-Type: `text/event-stream`
  - [ ] Each chunk prefixed with `data: `
  - [ ] Chunks are valid JSON objects
  - [ ] Final message is `data: [DONE]`
  - [ ] Keep-alive messages sent if needed

- [ ] **Streaming Chunk Format**
  ```json
  data: {"id":"chatcmpl-123","object":"chat.completion. chunk","created":1733356800,"model":"your-model","choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"},"finish_reason":null}]}

  data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1733356800,"model":"your-model","choices":[{"index":0,"delta":{"content":" world"},"finish_reason":null}]}

  data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1733356800,"model":"your-model","choices":[{"index":0,"delta":{},"finish_reason":"stop"}]}

  data: [DONE]
  ```

---

### Timeout & Performance

- [ ] **Timeout Configuration**
  - [ ] Request timeout clearly defined (e.g., 60 seconds)
  - [ ] Streaming timeout per chunk (e.g., 30 seconds of silence)
  - [ ] Configurable timeout for long-running requests
  - [ ] Graceful timeout handling with partial results if possible

- [ ] **Performance Requirements**
  - [ ] Time to first token (TTFT) < 2 seconds for streaming
  - [ ] Average response time < 5 seconds for short prompts
  - [ ] P95 latency documented
  - [ ] Throughput capacity specified (requests/second)

---

### Health & Status Endpoints

- [ ] **Health Check Endpoint**
  - [ ] `/health` or `/status` endpoint
  - [ ] Returns 200 OK when service is healthy
  - [ ] Includes basic system status
  - [ ] No authentication required for health checks

- [ ] **Example Health Response**
  ```json
  {
    "status": "operational",
    "version": "1.0.0",
    "uptime": 99.95,
    "timestamp": 1733356800
  }
  ```

---

### CORS & Headers

- [ ] **CORS Configuration** (if supporting browser requests)
  - [ ] `Access-Control-Allow-Origin` properly configured
  - [ ] `Access-Control-Allow-Methods` includes POST, OPTIONS
  - [ ] `Access-Control-Allow-Headers` includes Authorization, Content-Type
  - [ ] Preflight requests handled

- [ ] **Response Headers**
  - [ ] `Content-Type: application/json` for JSON responses
  - [ ] `Content-Type: text/event-stream` for streaming
  - [ ] Rate limit headers included
  - [ ] Request ID header for tracking (e.g., `X-Request-ID`)

---

### Logging & Debugging

- [ ] **Request Logging**
  - [ ] Every request logged with unique ID
  - [ ] Timestamp, endpoint, status code logged
  - [ ] Response time tracked
  - [ ] NO sensitive data (API keys, user content) in logs

- [ ] **Debug Support**
  - [ ] Request ID returned in response headers
  - [ ] Request ID included in error responses
  - [ ] Ability to trace requests through system
  - [ ] Detailed error logs for troubleshooting

---

### Versioning

- [ ] **API Versioning Strategy**
  - [ ] Version included in URL path (e.g., `/v1/`)
  - [ ] Backward compatibility maintained within major versions
  - [ ] Deprecation notices provided 6+ months in advance
  - [ ] Multiple versions supported simultaneously during transition
  - [ ] Clear migration guides between versions
```

This elaboration includes:

1. **Detailed endpoint configuration** with example request structures
2. **Comprehensive authentication** implementation details
3. **Advanced rate limiting** with headers and examples
4. **Extensive error handling** with specific error types and examples
5. **Request validation** requirements
6. **Complete response formats** for both streaming and non-streaming
7. **Performance and timeout** specifications
8. **Health check endpoints**
9. **CORS and header** configurations
10. **Logging and debugging** best practices
11. **API versioning** strategy

Would you like me to:
1.  Elaborate on any other sections?
2. Create code examples for implementing these requirements?
3. Generate test cases for API validation? 
