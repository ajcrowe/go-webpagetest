package wpt

type Run struct {
	FirstView  RunResult `json:"firstView"`
	RepeatView RunResult `json:"repeatView"`
}

type RunResult struct {
	URL                        string        `json:"URL"`
	LoadTime                   float64       `json:"loadTime"`
	TTFB                       float64       `json:"TTFB"`
	BytesOut                   float64       `json:"bytesOut"`
	BytesOutDoc                float64       `json:"bytesOutDoc"`
	BytesIn                    float64       `json:"bytesIn"`
	BytesInDoc                 float64       `json:"bytesInDoc"`
	Connections                int           `json:"connections"`
	Requests                   interface{}   `json:"requests"`
	RequestsDoc                int           `json:"requestsDoc"`
	Responses200               int           `json:"responses_200"`
	Responses404               int           `json:"responses_404"`
	ResponsesOther             int           `json:"responses_other"`
	Result                     int           `json:"result"`
	Render                     float64       `json:"render"`
	FullyLoaded                float64       `json:"fullyLoaded"`
	Cached                     int           `json:"cached"`
	DocTime                    float64       `json:"docTime"`
	DomTime                    float64       `json:"domTime"`
	ScoreCache                 int           `json:"score_cache"`
	ScoreCDN                   int           `json:"score_cdn"`
	ScoreGzip                  int           `json:"score_gzip"`
	ScoreCookies               int           `json:"score_cookies"`
	ScoreKeepAlive             int           `json:"score_keep-alive"`
	ScoreMinify                int           `json:"score_minify"`
	ScoreCombine               int           `json:"score_combine"`
	ScoreCompress              int           `json:"score_compress"`
	ScoreETags                 int           `json:"score_etags"`
	GzipTotal                  float64       `json:"gzip_total"`
	GzipSaving                 float64       `json:"gzip_saving"`
	MinifyTotal                float64       `json:"minify_total"`
	MinifySaving               float64       `json:"minify_saving"`
	ImageTotal                 float64       `json:"image_total"`
	ImageSaving                float64       `json:"image_saving"`
	OptimizationChecked        int           `json:"optimization_checked"`
	AFT                        int           `json:"aft"`
	DomElements                int           `json:"domElements"`
	Title                      string        `json:"title"`
	TitleTime                  float64       `json:"titleTime"`
	LoadEventStart             float64       `json:"loadEventStart"`
	LoadEventEnd               float64       `json:"loadEventEnd"`
	DomContentLoadedEventStart float64       `json:"domContentLoadedEventStart"`
	DomContentLoadedEventEnd   float64       `json:"domContentLoadedEventEnd"`
	LastVisualChange           float64       `json:"lastVisualChange"`
	BrowserName                string        `json:"browser_name"`
	BrowserVersion             string        `json:"browser_version"`
	ServerCount                int           `json:"server_count"`
	ServerRTT                  float64       `json:"server_rtt"`
	BasePageCDN                string        `json:"base_page_cdn"`
	AdultSite                  int           `json:"adult_site"`
	FixedViewport              int           `json:"fixed_viewport"`
	ScoreProgressiveJPEG       int           `json:"score_progressive_jpeg"`
	FirstPaint                 float64       `json:"firstPaint"`
	DocCPUMS                   float64       `json:"docCPUms"`
	FullyLoadedCPUMS           float64       `json:"fullyLoadedCPUms"`
	DocCPUPCT                  int           `json:"docCPUpct"`
	FullyLoadedDocCPUPCT       int           `json:"fullyLoadedCPUpct"`
	IsResponsive               int           `json:"isResponsive"`
	Date                       Timestamp     `json:"date"`
	SpeedIndex                 float64       `json:"SpeedIndex"`
	VisualComplete             float64       `json:"visualComplete"`
	Run                        int           `json:"run"`
	EffectiveBPS               float64       `json:"effectiveBps"`
	EffectiveBPSDoc            float64       `json:"effectiveBpsDoc"`
	Tester                     string        `json:"tester"`
	AvgRun                     int           `json:"avgRun"`
	Pages                      RunPages      `json:"pages"`
	Thumbnails                 RunThumbnails `json:"thumbnails"`
	Images                     RunImages     `json:"images"`
	RawData                    RunRawData    `json:"rawData"`
}

type RunPages struct {
	Details    string `json:"details"`
	Checklist  string `json:"checklist"`
	Report     string `json:"report"`
	Breakdown  string `json:"breakdown"`
	Domains    string `json:"domains"`
	ScreenShot string `json:"screenShot"`
}
type RunThumbnails struct {
	Waterfall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}
type RunImages struct {
	Waterfall  string `json:"waterfall"`
	Checklist  string `json:"checklist"`
	ScreenShot string `json:"screenShot"`
}
type RunRawData struct {
	Headers      string `json:"headers"`
	PageData     string `json:"pageData"`
	RequestsData string `json:"requestsData"`
}

type RunRequest struct {
	// todo
}
