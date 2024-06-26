package sayhi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"
)

// Checks project info by using RegEx
func TestGetProjectInfo(t *testing.T) {

	// Querying the endpoint
	req, err := http.NewRequest("GET", "/versionz", nil)
	if err != nil {
		t.Fatalf("Could not create GET request: %v", err)
	}
	rec := httptest.NewRecorder()
	GetProjectInfo(rec, req)

	// Retrieving result
	res := rec.Result()
	defer res.Body.Close()
	bodyBytes, err := ioutil.ReadAll(res.Body)
	t.Logf("Body converted to string: %s\n", string(bytes.TrimSpace(bodyBytes)))
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}
	var info ProjectInfo
	err = json.Unmarshal(bodyBytes, &info)
	if err != nil {
		t.Fatalf("Could not parse the response: %v", err)
	}

	// Checking the status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: expected %v, got %v", http.StatusOK, status)
	}

	// Checking the git hash
	gitHashRegex, _ := regexp.Compile("^[0-9a-f]{5,40}$")
	if !gitHashRegex.MatchString(info.GitHash) {
		t.Errorf("Invalid git hash: %v", info.GitHash)
	}

	// Checking the project name
	projectNameRegex, _ := regexp.Compile("^[a-zA-Z0-9_.-/]+$")
	if !projectNameRegex.MatchString(info.ProjectName) {
		t.Errorf("Invalid project name: %v", info.ProjectName)
	}

	// Checking the query time
	// Regex: https://www.debuggex.com/r/N4Jk-0WQtcTHFHkM
	timeDateRegex, _ := regexp.Compile("^(?:[1-9]\\d{3}-(?:(?:0[1-9]|1[0-2])-(?:0[1-9]|1\\d|2[0-8])|" +
		"(?:0[13-9]|1[0-2])-(?:29|30)|(?:0[13578]|1[02])-31)|(?:[1-9]\\d(?:0[48]|[2468][048]|[13579][26])|" +
		"(?:[2468][048]|[13579][26])00)-02-29)T(?:[01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d.[0-9]{9}(?:Z|[+-][01]\\d:" +
		"[0-5]\\d)$")
	if info.QueryTime != nil {
		queryTime := *info.QueryTime
		queryTimeString := queryTime.Format(time.RFC3339Nano)
		t.Logf("Query time is: %s", queryTimeString)
		if !timeDateRegex.MatchString(queryTimeString) {
			t.Errorf("Invalid query time: %s", queryTime)
		}
	}
}

// Checks how the service says hi back with various inputs
func TestSayHi(t *testing.T) {

	tests := []struct {
		noparam  bool
		name     string
		person   string
		expected string
		status   int
		error    string
	}{
		{
			name:     "No specified name",
			noparam:  true,
			expected: "Hello Stranger",
		},
		{
			name:     "Empty string name",
			person:   "",
			expected: "Hello Stranger",
		},
		{
			name:     "Requirement example",
			person:   "AlfredENeumann",
			expected: "Hello Alfred E Neumann",
		},
		{
			name:     "UpperCamel case",
			person:   "JackWhite",
			expected: "Hello Jack White",
		},
		{
			name:     "lowerCamel case",
			person:   "jackWhite",
			expected: "Hello jack White",
		},
		{
			name:     "Name with acronym in the middle",
			person:   "FredFDDurst",
			expected: "Hello Fred FD Durst",
		},
		{
			name:   "snake_case",
			person: "jack_white",
			status: http.StatusBadRequest,
			error:  InvalidPersonNameErrorMessage,
		},
		{
			name:   "Invalid name (single digit)",
			person: "3",
			status: http.StatusBadRequest,
			error:  InvalidPersonNameErrorMessage,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			// Querying the endpoint
			endpoint := "/helloworld"
			if !tc.noparam {
				endpoint = endpoint + "?name=" + tc.person
			}
			req, err := http.NewRequest("GET", endpoint, nil)
			if err != nil {
				t.Fatalf("Could not create GET request: %v", err)
			}
			rec := httptest.NewRecorder()
			SayHi(rec, req)

			// Retrieving result
			res := rec.Result()
			defer res.Body.Close()
			bodyBytes, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Could not read response: %v", err)
			}
			body := string(bytes.TrimSpace(bodyBytes))

			// If the expected status code is missing, 200 is implied
			if tc.status == 0 {

				tc.status = http.StatusOK

				// Error code must necessarily be empty
				if tc.error != "" {
					t.Fatalf("Invalid test %v: expected %v, but with error %v", tc.name, tc.status, tc.error)
				}
			}

			// Checking the status code
			if status := rec.Code; status != tc.status {
				t.Errorf("Handler returned wrong status code: expected %v, got %v", tc.status, status)
			}

			// If an error was expected
			if tc.error != "" {

				// Checking the error message, if any
				if tc.status != http.StatusOK && body != tc.error {
					t.Errorf("Expected error message %q, got %q", tc.error, body)
				}
			} else {

				// Checking the response body
				if body != tc.expected {
					t.Errorf("Handler returned unexpected body: expected %v, got %v", tc.expected, rec.Body.String())
				}
			}
		})
	}
}
