package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Request struct {
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Body    string            `yaml:"body,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
	Params  map[string]string `yaml:"params,omitempty"`
}

type Response struct {
	Status  int               `yaml:"status,omitempty"`
	Body    string            `yaml:"body,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

type Criteria struct {
	Type     string   `yaml:"type"`
	Response Response `yaml:"reponse"`
}

type Param struct {
	Type string   `yaml:"type"`
	From int      `yaml:"from,omitempty"`
	To   int      `yaml:"to,omitempty"`
	File string   `yaml:"file,omitempty"`
	Dict []string `yaml:"dict,omitempty"`
}

func httpRes2CustomRes(httpRes *http.Response) *Response {
	res := new(Response)

	res.Status = httpRes.StatusCode
	body, err := io.ReadAll(httpRes.Body)

	if err != nil {
		return nil
	}

	res.Body = string(body)

	// TODO: test this
	for name, values := range httpRes.Header {
		res.Headers[name] = strings.Join(values, ", ")
	}

	return res
}

func makeRequest(request Request, paramName string, paramValue string) *http.Response {
	URL := strings.ReplaceAll(request.URL, fmt.Sprintf("$%s$", paramName), paramValue)
	body := strings.ReplaceAll(request.Body, fmt.Sprintf("$%s$", paramName), paramValue)

	preparedUrl, err := url.Parse(URL)
	if err != nil {
		return nil
	}

	// Parse params
	params := preparedUrl.Query()
	for name, value := range request.Params {
		val := strings.ReplaceAll(value, fmt.Sprintf("$%s$", paramName), paramValue)
		params.Add(name, val)
	}
	preparedUrl.RawQuery = params.Encode()

	// Create request
	req, err := http.NewRequest(request.Method, preparedUrl.String(), strings.NewReader(body))
	if err != nil {
		return nil
	}

	// Parse headers
	for name, value := range request.Headers {
		val := strings.ReplaceAll(value, fmt.Sprintf("$%s$", paramName), paramValue)
		req.Header.Set(name, val)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil
	}

	fmt.Printf("[%s: %s] %s\n", paramName, paramValue, res.Status)
	return res
}


func checkCriteria(response *http.Response, criteria *Response) bool {
	if criteria.Status != 0 {
		if criteria.Status != response.StatusCode {
			return false
		}
	}

	if criteria.Body != "" {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			return false
		}

		match, err := regexp.MatchString(criteria.Body, string(body))
		if err != nil || !match {
			return false
		}
	}

	if len(criteria.Headers) > 0 {
		for name, value := range criteria.Headers {
			resVal, ok := response.Header[name] 
			if !ok {
				return false
			}

			// TODO: Test this
			if strings.Join(resVal, ", ") != value {
				return false
			}
		}
	}

	return true
}