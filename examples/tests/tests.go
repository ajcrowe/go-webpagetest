package main

import (
	"fmt"
	"github.com/ajcrowe/go-webpagetest"
)

func main() {
	config := wpt.Config{
		Host:   "https://www.webpagetest.org",
		APIKey: "xxxxxxxxxxxxx",
	}

	client, _ := wpt.NewClient(config)

	wpt.SetPollingInterval(5)

	params := &wpt.TestParams{
		URL:  "http://www.google.com",
		Runs: 5,
	}

	test, err := wpt.NewTest(params, client)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = test.Run()
	if err != nil {
		fmt.Print(err)

	} else {
		fmt.Printf("wpt-example: watching test %s for status updates...\n", test.RequestID)
		for state := range test.StatusChan {

			fmt.Printf("wpt-example: status: %s\n", state)
		}
		fmt.Println("---")
		fmt.Printf("wpt-example: Test:\t%s\n", test.RequestID)
		fmt.Printf("wpt-example: Load-time:\t%d\nms", test.Results.Data.Average.FirstView.LoadTime)
		fmt.Printf("wpt-example: TTFB:\t%d\nms", test.Results.Data.Average.FirstView.TTFB)
	}

}
