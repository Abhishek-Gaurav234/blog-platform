# OpenRouter Provider Listing Requirements Checklist

A comprehensive guide for model providers to get listed on OpenRouter. 

## 📋 Table of Contents
- [Overview](#overview)
- [Pre-Registration Checklist](#pre-registration-checklist)
- [API Compatibility Requirements](#api-compatibility-requirements)
- [Model Information Requirements](#model-information-requirements)
- [Technical Implementation](#technical-implementation)
- [Documentation Requirements](#documentation-requirements)
- [Quality Standards](#quality-standards)
- [Integration Process](#integration-process)
- [Testing & Validation](#testing--validation)

---

## Overview

This checklist covers all requirements for listing your AI models on OpenRouter. Complete each section to ensure a smooth integration process.

---

## Pre-Registration Checklist

### Basic Requirements
- [ ] Models support OpenAI-compatible API format OR
- [ ] Models support Anthropic-compatible API format OR
- [ ] Models support Google-compatible API format OR
- [ ] Models support Server-Sent Events (SSE) streaming
- [ ] API is production-ready and stable
- [ ] You have pricing structure defined
- [ ] You can provide API documentation

---

## API Compatibility Requirements

### Supported API Formats
- [ ] **OpenAI Format**: `/v1/chat/completions` endpoint
- [ ] **Anthropic Format**: `/v1/messages` endpoint
- [ ] **Google Format**: `generateContent` endpoint
- [ ] **Streaming**: SSE (Server-Sent Events) support

### Parameter Support
- [ ] `model` - Model identifier
- [ ] `messages` - Conversation history (OpenAI format)
- [ ] `max_tokens` - Maximum completion tokens
- [ ] `temperature` - Sampling temperature (0-2)
- [ ] `top_p` - Nucleus sampling parameter
- [ ] `top_k` - Top-k sampling parameter (if applicable)
- [ ] `frequency_penalty` - Frequency penalty (-2 to 2)
- [ ] `presence_penalty` - Presence penalty (-2 to 2)
- [ ] `stop` - Stop sequences
- [ ] `stream` - Streaming support (true/false)
- [ ] `tools` / `functions` - Function calling support (if applicable)
- [ ] `response_format` - JSON mode support (if applicable)

### Error Handling
- [ ] Returns standard HTTP status codes (4xx, 5xx)
- [ ] Error responses include descriptive messages
- [ ] Rate limit errors return `429` status
- [ ] Authentication errors return `401` status
- [ ] Invalid requests return `400` status with details

---

## Model Information Requirements

### Model Identification
- [ ] Unique model ID (e.g., `provider-name/model-name-version`)
- [ ] Display name for the model
- [ ] Model version or release date
- [ ] Model family/architecture (e.g., GPT, Claude, Llama)

### Pricing Information
- [ ] Prompt token price (per 1M tokens)
- [ ] Completion token price (per 1M tokens)
- [ ] Any special pricing tiers or discounts
- [ ] Image input pricing (if applicable)
- [ ] Image output pricing (if applicable)

### Model Specifications
- [ ] Context length (input + output tokens)
- [ ] Maximum output tokens
- [ ] Training data cutoff date
- [ ] Supported languages
- [ ] Modality (text, vision, audio, etc.)

### Model Capabilities
- [ ] Function/tool calling support (yes/no)
- [ ] JSON mode support (yes/no)
- [ ] Vision capabilities (yes/no)
- [ ] Streaming support (yes/no)
- [ ] System message support (yes/no)

---

## Technical Implementation

### API Endpoint Configuration
- [ ] Production API base URL
- [ ] Staging/testing endpoint (if available)
- [ ] Endpoint supports HTTPS
- [ ] API versioning strategy

### Authentication
- [ ] Authentication method (API key, OAuth, etc.)
- [ ] API key format and location (header, query param)
- [ ] Sample authentication header example
- [ ] Key rotation policy (if applicable)

### Rate Limiting
- [ ] Requests per minute (RPM) limit
- [ ] Tokens per minute (TPM) limit
- [ ] Rate limit headers in responses
- [ ] Rate limit error handling

### Response Format
- [ ] OpenAI-compatible response structure
- [ ] Streaming response format (if applicable)
- [ ] Usage statistics in response
  - [ ] `prompt_tokens`
  - [ ] `completion_tokens`
  - [ ] `total_tokens`
- [ ] Finish reason indicators (`stop`, `length`, `tool_calls`)

---

## Documentation Requirements

### Model Description
- [ ] Clear, concise model description (2-3 sentences)
- [ ] Primary use cases
- [ ] Model strengths and capabilities
- [ ] Any limitations or restrictions

### Use Cases & Examples
- [ ] At least 3 example use cases
- [ ] Sample prompts and expected outputs
- [ ] Best practices for prompt engineering
- [ ] Parameter recommendations for different scenarios

### API Documentation
- [ ] Complete endpoint documentation
- [ ] Request/response examples
- [ ] Parameter descriptions
- [ ] Error codes and meanings
- [ ] cURL examples
- [ ] SDK examples (Python, JavaScript, etc.)

### Integration Guide
- [ ] Quick start guide
- [ ] Authentication setup
- [ ] First API call example
- [ ] Common issues and troubleshooting

---

## Quality Standards

### Performance Benchmarks
- [ ] Average response latency (ms)
- [ ] P95 latency (ms)
- [ ] Streaming time to first token (ms)
- [ ] Tokens per second (for streaming)

### Reliability
- [ ] Uptime SLA (target: 99.9%)
- [ ] Planned maintenance notification process
- [ ] Incident response time
- [ ] Status page URL

### Response Quality
- [ ] Model performance benchmarks (MMLU, HumanEval, etc.)
- [ ] Quality assurance process
- [ ] Moderation/safety measures
- [ ] Content filtering policies

---

## Integration Process

### Step 1: Initial Registration
- [ ] Submit provider application to OpenRouter
- [ ] Provide company/organization details
- [ ] Share API documentation
- [ ] Define pricing structure

### Step 2: Technical Review
- [ ] OpenRouter team reviews API documentation
- [ ] Discuss any compatibility issues
- [ ] Confirm parameter support
- [ ] Review error handling

### Step 3: API Testing
- [ ] Provide test API key to OpenRouter
- [ ] OpenRouter conducts integration testing
- [ ] Fix any identified issues
- [ ] Validate streaming functionality
- [ ] Confirm pricing calculations

### Step 4: Documentation Review
- [ ] Submit model descriptions
- [ ] Provide use case examples
- [ ] Share pricing details
- [ ] Include any special requirements

### Step 5: Staging Validation
- [ ] Test in OpenRouter staging environment
- [ ] Validate all endpoints
- [ ] Confirm parameter transformations
- [ ] Test error scenarios
- [ ] Verify billing calculations

### Step 6: Go-Live Preparation
- [ ] Final security review
- [ ] Confirm production API endpoint
- [ ] Set up monitoring and alerts
- [ ] Prepare support documentation
- [ ] Define escalation procedures

### Step 7: Launch
- [ ] Production API integration
- [ ] Model listing goes live on OpenRouter
- [ ] Monitor initial traffic
- [ ] Collect user feedback
- [ ] Address any issues promptly

---

## Testing & Validation

### Functional Testing
- [ ] Basic completion request works
- [ ] Streaming responses work correctly
- [ ] All supported parameters function properly
- [ ] Error handling works as expected
- [ ] Rate limiting behaves correctly

### Load Testing
- [ ] API handles expected traffic volume
- [ ] Latency remains acceptable under load
- [ ] Rate limits are properly enforced
- [ ] No degradation in response quality

### Edge Cases
- [ ] Very long prompts (near context limit)
- [ ] Empty or minimal prompts
- [ ] Special characters and Unicode
- [ ] Concurrent requests
- [ ] Invalid parameters

### Compatibility Testing
- [ ] Works with OpenRouter's transformation layer
- [ ] Compatible with popular SDKs
- [ ] Functions in different programming languages
- [ ] Mobile and web client compatibility

---

## Post-Launch Checklist

### Monitoring
- [ ] Set up uptime monitoring
- [ ] Monitor API latency
- [ ] Track error rates
- [ ] Monitor token usage and billing

### Support
- [ ] Designate technical contact
- [ ] Set up support email/channel
- [ ] Create FAQ document
- [ ] Join OpenRouter provider community

### Maintenance
- [ ] Regular API health checks
- [ ] Plan for model updates
- [ ] Pricing review process
- [ ] Documentation updates

### Optimization
- [ ] Collect performance metrics
- [ ] Gather user feedback
- [ ] Identify improvement areas
- [ ] Plan feature enhancements

---

## Resources

- **OpenRouter Documentation**: [https://openrouter.ai/docs](https://openrouter.ai/docs)
- **API Reference**: [https://openrouter. ai/docs/api-reference](https://openrouter. ai/docs/api-reference)
- **Provider Portal**: Contact OpenRouter team for access
- **Support**: support@openrouter.ai

---

## Notes

### Tips for Success
1. **Start with one model**: Begin with your most stable model before adding others
2. **Test thoroughly**: Use the checklist to ensure nothing is missed
3. **Clear documentation**: The better your docs, the smoother the integration
4. **Monitor closely**: Watch your API during initial launch period
5. **Communicate**: Keep OpenRouter team updated on any changes

### Common Issues to Avoid
- ❌ Incomplete parameter support
- ❌ Inconsistent error responses
- ❌ Missing usage statistics in responses
- ❌ Incorrect pricing calculations
- ❌ Poor documentation
- ❌ Inadequate rate limiting

### Best Practices
- ✅ Follow OpenAI API standards closely
- ✅ Provide detailed error messages
- ✅ Include usage statistics in every response
- ✅ Test with various parameter combinations
- ✅ Document any deviations from standards
- ✅ Set up comprehensive monitoring

---

## Contact

For questions or to begin the integration process:
- **Email**: support@openrouter. ai
- **Website**: https://openrouter.ai

---

**Version**: 1.0  
**Last Updated**: December 2025
