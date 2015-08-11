package wpt

type Run struct {
	FirstView  RunResult `json:"firstView"`
	RepeatView RunResult `json:"repeatView"`
}

type RunResult struct {
	URL            string        `json:"URL"`
	LoadTime       int           `json:"loadTime"`
	TTFB           int           `json:"TTFB"`
	BytesOut       int           `json:"bytesOut"`
	BytesOutDoc    int           `json:"bytesOutDoc"`
	BytesIn        int           `json:"bytesIn"`
	BytesInDoc     int           `json:"bytesInDoc"`
	Requests       int           `json:"requests"`
	RequestsDoc    int           `json:"requestsDoc"`
	Result         int           `json:"result"`
	Render         int           `json:"render"`
	FullyLoaded    int           `json:"fullyLoaded"`
	Cached         int           `json:"cached"`
	Web            int           `json:"web"`
	DocTime        int           `json:"docTime"`
	DomTime        int           `json:"domTime"`
	ScoreCache     int           `json:"score_cache"`
	ScoreCDN       int           `json:"score_cdn"`
	ScoreGzip      int           `json:"score_gzip"`
	ScoreCookies   int           `json:"score_cookies"`
	ScoreKeepAlive int           `json:"score_keep-alive"`
	ScoreMinify    int           `json:"score_minify"`
	ScoreCombine   int           `json:"score_combine"`
	ScoreCompress  int           `json:"score_compress"`
	ScoreETags     int           `json:"score_etags"`
	Date           int           `json:"date"`
	Pages          RunPages      `json:"pages"`
	Thumbnails     RunThumbnails `json:"thumbnails"`
	Images         RunImages     `json:"images"`
	RawData        RunRawData    `json:"rawData"`
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
