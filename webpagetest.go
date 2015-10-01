package wpt

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	defaultBaseURL = "http://www.webpagetest.org"
	// Test status
	testNew      = "new"
	testFailed   = "failed"
	testQueued   = "queued"
	testRunning  = "running"
	testFinished = "finished"
	testComplete = "complete"
	testNotFound = "notfound"
	testTimedOut = "timeout"
	testError    = "error"
	// URL endpoints
	urlLocations = "getLocations.php"
	urlRunTest   = "runtest.php"
	urlCancel    = "cancelTest.php"
	urlStatus    = "testStatus.php"
	urlResults   = "jsonResult.php"
)

var (
	serverURL            *url.URL
	pollingInterval      time.Duration
	maximumMonitorPeriod time.Duration
	// Valid Connection Profiles
	connectionProfiles = []string{
		"DSL",
		"Cable",
		"FIOS",
		"Dial",
		"3G",
		"3GFast",
		"Native",
		"custom",
	}
)

// setup the defaults for WebPageTest server and polling intervals
func init() {
	serverURL, _ = url.Parse(defaultBaseURL)
	pollingInterval = 60 * time.Second
	maximumMonitorPeriod = 1800 * time.Second // 30 minutes
}

// function to set your own private WebPageTest server url
// this is validated and stored as a url.URL struct
func SetURL(host string) error {
	var err error
	serverURL, err = url.Parse(host)
	if err != nil {
		return errors.New("webpagetest: invalid server url")
	}
	return nil
}

// function to overwrite the default 60 second polling interval for
// checking test status.
func SetPollingInterval(period int64) {
	pollingInterval = time.Duration(period) * time.Second
}

// function to overwrite the default 600 second maximum period to monitor
// tests for status changes
func SetMaximumMonitorPeriod(period int64) {
	maximumMonitorPeriod = time.Duration(period) * time.Second
}

type Config struct {
	// Host specifies the host to send http requests to
	Host string
	// APIKey to send with requests to the WebPageTest instance
	APIKey string
	// Timeout for all http requests
	// if not specified this defaults to 0 and requests will not timeout
	Timeout time.Duration
}

type Client struct {
	// URL stores the Host and standard query params use for all
	// requests to the WPT server
	URL        url.URL
	APIKey     string
	HTTPClient *http.Client
}

// NewClient parses the provided Config and returns a Client struct pointer
func NewClient(c Config) (*Client, error) {
	u, err := url.Parse(c.Host)
	if err != nil {
		return nil, errors.New("webpagetest: error parsing configuration")
	}
	v := url.Values{}
	v.Add("f", "json")
	if c.APIKey != "" {
		v.Add("k", c.APIKey)
	}
	u.RawQuery = v.Encode()
	client := Client{
		URL:        *u,
		HTTPClient: &http.Client{Timeout: c.Timeout},
	}

	return &client, nil
}

// Locations queries webpagetest using getLocation.php and parses the returned data into
// an array of Location structs and returns this.
func (c *Client) Locations() (locations Locations, err error) {

	var resp LocationsResponse
	err = c.query(urlLocations, "", &resp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, location := range resp.Data {
		locations = append(locations, location)
	}

	return locations, nil
}

// Get the test results for a specific test ID.
func (c *Client) GetResults(id string) (results TestResults, err error) {
	qs := fmt.Sprintf("test=%s", id)
	err = c.query(urlResults, qs, &results)
	if err != nil {
		fmt.Printf("webpagetest: error loading results for test: %s", id)
		return TestResults{}, err
	}
	return results, nil

}

// query takes a given path and query string and calls the WPT host. The response is
// then UnMarshalled into the provided struct pointer
func (c *Client) query(path string, qs string, respStruct interface{}) error {
	u := c.URL
	u.Path = path
	u.RawQuery = fmt.Sprintf("%s&%s", u.RawQuery, qs)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return errors.New("webpagetest: error creating request")
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return errors.New("webpagetest: error querying webpagetest")
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("webpagetest: bad response code from host")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("webpagetest: error ready response body")
	}

	err = json.Unmarshal(body, &respStruct)

	return nil
}
