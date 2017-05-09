package wpt

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

type Test struct {
	RequestID  string
	StatusChan chan string
	Status     string
	Monitor    bool
	Client     *Client
	Response   *TestRequest
	Results    *TestResults
	Params     *TestParams
}

type TestParams struct {
	URL    string `url:"url"`
	APIKey string `url:"k,omitempty"`
	Label  string `url:"label,omitempty"`
	// locations and connectivity are combined at validation into LocationString
	Location       string `url:"-"`
	Connectivity   string `url:"-"`
	LocationString string `url:"location,omitempty"`
	// Block string array is converted to a space delimited string at validation
	Block       []string `url:"-"`
	BlockString string   `url:"block,omitempty"`

	Login         string `url:"login,omitempty"`
	Password      string `url:"password,omitempty"`
	Notify        string `url:"notify,omitempty"`
	Pingback      string `url:"pingback,omitempty"`
	CMDLine       string `url:"cmdline,omitempty"`
	TSViewID      string `url:"tsview_id,omitempty"`
	Custom        string `url:"custom,omitempty"`
	Tester        string `url:"tester,omitempty"`
	Affinity      string `url:"affinity,omitempty"`
	DomElement    string `url:"domelement,omitempty"`
	Script        string `url:"script,omitempty"`
	TimelineStack int    `url:"timelineStack,omitempty"`
	Runs          int    `url:"runs,omitempty"`
	Connections   int    `url:"connections,omitempty"`
	Authtype      int    `url:"authType,omitempty"`
	BWDown        int    `url:"bwDown,omitempty"`
	BWUp          int    `url:"bwUp,omitempty"`
	Latency       int    `url:"latency,omitempty"`
	PackLossRate  int    `url:"plr,omitempty"`
	IQ            int    `url:"iq,omitempty"`
	// these are integer representations of any booleans
	FVOnly     int `url:"fvonly,omitempty"`
	Private    int `url:"private,omitempty"`
	StopOnLoad int `url:"web10,omitempty"`
	Video      int `url:"video,omitempty"`
	TCPDump    int `url:"tcpdump,omitempty"`
	NoOpt      int `url:"noopt,omitempty"`
	NoImages   int `url:"noimages,omitempty"`
	NoHeaders  int `url:"noheaders,omitempty"`
	PNGSS      int `url:"pngss,omitempty"`
	NoScript   int `url:"noscript,omitempty"`
	ClearCerts int `url:"clearcerts,omitempty"`
	Mobile     int `url:"mobile,omitempty"`
	MV         int `url:"mv,omitempty"`
	HTMLBody   int `url:"htmlbody,omitempty"`
	Timeline   int `url:"timeline,omitempty"`
	IgnoreSSL  int `url:"ignoreSSL,omitempty"`
}

type TestRequest struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       struct {
		TestId     string `json:"testId"`
		OwnerKey   string `json:"ownerKey"`
		JSONUrl    string `json:"jsonUrl"`
		XMLUrl     string `json:"xmlUrl"`
		UserUrl    string `json:"userUrl"`
		SummaryCSV string `json:"summaryCSV"`
		DetailCSV  string `json:"detailCSV"`
	} `json:"data"`
}

type TestStatus struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Id         string `json:"id"`
	Data       struct {
		StatusCode int        `json:"statusCode"`
		StatusText string     `json:"statusText"`
		Id         string     `json:"id"`
		TestInfo   TestParams `json:"testInfo"`
	} `json:"data"`
	TestId          string `json:"testId"`
	Runs            int    `json:"runs"`
	FVOnly          int    `json:"fvonly"`
	Remote          bool   `json:"remote"`
	TestsExpected   int    `json:"testsExpected"`
	Location        string `json:"location"`
	Elapsed         int    `json:"elapsed"`
	BehindCount     int    `json:"behindCount"`
	FVRunsCompleted int    `json:"fvRunsCompleted"`
	RVRunsCompleted int    `json:"rvRunsCompleted"`
}

type TestResults struct {
	StatusCode int    `json:"statusCode"`
	StatusText string `json:"statusText"`
	Data       struct {
		Id               string         `json:"id"`
		URL              string         `json:"url"`
		Summary          string         `json:"summary"`
		TestURL          string         `json:"testUrl"`
		Location         string         `json:"location"`
		From             string         `json:"from"`
		Connectivity     string         `json:"connectivity"`
		BWDown           int            `json:"bwDown"`
		BWUp             int            `json:"bwUp"`
		Latency          int            `json:"latency"`
		PackLossRate     string         `json:"plr"`
		Completed        Timestamp      `json:"completed"`
		Tester           string         `json:"tester"`
		TesterDNS        string         `json:"testerDNS"`
		FVOnly           bool           `json:"fvonly"`
		SuccessfulFVRuns int            `json:"successfulFVRuns"`
		SuccessfulRVRuns int            `json:"successfulRVRuns"`
		Runs             map[string]Run `json:"runs"`
		Average          Run            `json:"average"`
		StdDev           Run            `json:"standardDeviation"`
		Median           Run            `json:"median"`
	} `json:"data"`
}

// NewTest takes a TestParams and Client struct  and returns a Test
func NewTest(tp *TestParams, c *Client) (*Test, error) {
	// validate test params
	err := tp.Validate()
	if err != nil {
		return nil, err
	}

	return &Test{
		Client:  c,
		Monitor: true,
		Params:  tp,
	}, nil
}

func (t *Test) Run() error {
	t.StatusChan = make(chan string)
	// send the request to webpagetest
	err := t.Client.query(urlRunTest, t.Params.getQueryString(), &t.Response)
	if err != nil {
		t.setStatus(testFailed)
		return err
	}
	if t.Response.StatusCode != 200 {
		t.setStatus(testFailed)
		return errors.New(fmt.Sprintf("webpagetest: bad status code %v when submitting test", t.Response.StatusCode))
	}
	// update the Test struct
	t.RequestID = t.Response.Data.TestId
	t.Status = testQueued
	// call the monitor if set in test to update the Test
	if t.Monitor {
		go t.monitor()
	}
	return nil
}

// monitor periodically polls the server for the status of the test
// if the status changes it will transfer the new status over the
// State channel and load the result
func (t *Test) monitor() {
	expired := time.After(maximumMonitorPeriod)
	defer close(t.StatusChan)

	var status string
	for {
		select {
		case <-expired:
			t.setStatus(testTimedOut)
			return
		default:
			// sleep for defined interval
			time.Sleep(pollingInterval)
			// get latest status
			status = t.CheckStatus()
			// only send update if status has changed
			if t.Status != status {
				// send status over the status channel
				t.setStatus(status)
				// if test has finished call function to load results
				if status == testFinished || status == testNotFound || status == testCancelled {
					err := t.LoadResults()
					if err != nil {
						t.setStatus(testFailed)
					} else {
						t.setStatus(testComplete)
					}
					return
				}
			}
		}
	}
}

// CheckStatus returns the current status of the test by quering the testStatus.php endpoint
func (t *Test) CheckStatus() string {
	var status TestStatus
	err := t.Client.query(urlStatus, fmt.Sprintf("test=%s", t.RequestID), &status)
	if err != nil {
		fmt.Printf("webpagetest: error checking status for test: %s\n", t.RequestID)
		return testError
	}

	switch {
	case status.StatusCode == 100:
		return testRunning
	case status.StatusCode == 101:
		return testQueued
	case status.StatusCode == 200:
		return testFinished
	case status.StatusCode == 400:
		return testNotFound
	case status.StatusCode == 402:
		return testCancelled
	default:
		fmt.Printf("webpagetest: unknown status code %d returned for test %s\n", status.StatusCode, t.RequestID)
		return testError
	}
}

// LoadResults is called
func (t *Test) LoadResults() error {
	qs := fmt.Sprintf("test=%s", t.RequestID)
	err := t.Client.query(urlResults, qs, &t.Results)
	if err != nil {
		return errors.New(fmt.Sprintf("webpagetest: error loading results for test: %s", t.RequestID))
	}
	return nil

}

func (t *Test) setStatus(state string) {
	t.Status = state
	go func() {
		// There's no guarantee this channel hasn't already been closed elsewhere.
		defer func() {
			recover()
		}()
		t.StatusChan <- state
	}()
}

// this validates various options present against know valid inputs
func (tp *TestParams) Validate() error {
	// check URL is set
	if tp.URL == "" {
		return errors.New("webpagetest: no url specified")
	}
	// if Connection is specified check it's valid and Location is also specified
	if tp.Connectivity != "" {
		if tp.Location == "" {
			return errors.New("webpagetest: you must set a location to specify a connection profile")
		}
		found := false
		for _, profile := range connectionProfiles {
			if tp.Connectivity == profile {
				found = true
			}
		}
		if !found {
			return errors.New(fmt.Sprintf("webpagetest: invalid connection profile %s", tp.Connectivity))
		}
		// if custom is specified check that the bandwidth params are set
		if (tp.Connectivity == "custom") && ((tp.BWDown == 0) || (tp.BWUp == 0)) {
			return errors.New("webpagetest: you must set BWUp & BWDown to use custom connectivity")
		}
		//
		tp.LocationString = fmt.Sprintf("%s.%s", tp.Location, tp.Connectivity)

	}
	// block set convert string array to space delimited string
	if tp.Block != nil {
		tp.BlockString = strings.Join(tp.Block, " ")
	}
	return nil
}

func (tp *TestParams) getQueryString() string {
	qs, _ := query.Values(tp)
	return qs.Encode()
}

type Timestamp time.Time

func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(t).Unix()
	stamp := fmt.Sprint(ts)
	return []byte(stamp), nil
}
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}
	*t = Timestamp(time.Unix(int64(ts), 0))
	return nil
}
