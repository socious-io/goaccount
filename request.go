package goaccount

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type HTTPRequestMethodType string

const (
	MethodGet    HTTPRequestMethodType = "GET"
	MethodPost   HTTPRequestMethodType = "POST"
	MethodPut    HTTPRequestMethodType = "PUT"
	MethodDelete HTTPRequestMethodType = "DELETE"
)

type RequestOptions struct {
	Method   HTTPRequestMethodType
	Endpoint string
	Token    *SessionToken
	Headers  map[string]string
	Query    map[string]string
	Body     interface{}
}

// HTTPRequest sends an HTTP request with given method, URL, headers, query parameters, and body as map[string]any.
func Request(options RequestOptions) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Build the full URL with query parameters
	reqURL, err := url.Parse(options.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	query := reqURL.Query()
	for key, value := range options.Query {
		query.Set(key, value)
	}
	reqURL.RawQuery = query.Encode()

	// Convert body map to JSON if provided
	var reqBody io.Reader
	if options.Body != nil {
		jsonBody, err := json.Marshal(options.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	// Create request with body
	req, err := http.NewRequest(string(options.Method), reqURL.String(), reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}

	// Ensure Content-Type is set when sending a JSON body
	if options.Body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

func RequestMultipart(options RequestOptions) ([]byte, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Build URL with query parameters
	reqURL, err := url.Parse(options.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}
	query := reqURL.Query()
	for key, value := range options.Query {
		query.Set(key, value)
	}
	reqURL.RawQuery = query.Encode()

	// Create multipart form
	var bodyBuffer bytes.Buffer
	writer := multipart.NewWriter(&bodyBuffer)

	// Process form fields
	for key, value := range options.Body.(map[string]any) {
		switch v := value.(type) {
		case string:
			// Add text field
			_ = writer.WriteField(key, v)

		case multipart.File:
			// Create form file field
			part, err := writer.CreateFormFile(key, "uploaded_file")
			if err != nil {
				return nil, fmt.Errorf("failed to create form file: %w", err)
			}

			// Copy file data to form field
			_, err = io.Copy(part, v)
			if err != nil {
				return nil, fmt.Errorf("failed to copy file data: %w", err)
			}

		default:
			return nil, fmt.Errorf("unsupported form field type for key %s", key)
		}
	}

	// Close writer
	writer.Close()

	// Create request
	req, err := http.NewRequest(string(options.Method), reqURL.String(), &bodyBuffer)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	for key, value := range options.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Perform request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}
