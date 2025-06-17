package dc_log

import (
	"testing"
	"time"

	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

func TestRecord_Structure(t *testing.T) {
	// Test Record struct creation and fields
	now := time.Now()
	record := &Record{
		Timestamp:        now,
		ReqMethod:        "GET",
		ReqUrl:           "https://api.example.com/v1/users",
		ResCode:          200,
		ResContentLength: 1024,
		Latency:          0.125,
	}

	if record.Timestamp != now {
		t.Error("Timestamp not set correctly")
	}
	if record.ReqMethod != "GET" {
		t.Errorf("Expected ReqMethod 'GET', got '%s'", record.ReqMethod)
	}
	if record.ReqUrl != "https://api.example.com/v1/users" {
		t.Errorf("Expected ReqUrl 'https://api.example.com/v1/users', got '%s'", record.ReqUrl)
	}
	if record.ResCode != 200 {
		t.Errorf("Expected ResCode 200, got %d", record.ResCode)
	}
	if record.ResContentLength != 1024 {
		t.Errorf("Expected ResContentLength 1024, got %d", record.ResContentLength)
	}
	if record.Latency != 0.125 {
		t.Errorf("Expected Latency 0.125, got %f", record.Latency)
	}
}

func TestUrlFormat(t *testing.T) {
	testCases := []struct {
		name     string
		reqUrl   string
		shorten  bool
		expected string
	}{
		{
			name:     "Full URL not shortened",
			reqUrl:   "https://api.example.com/v1/users?id=123",
			shorten:  false,
			expected: "https://api.example.com/v1/users?id=123",
		},
		{
			name:     "Full URL shortened to path",
			reqUrl:   "https://api.example.com/v1/users?id=123",
			shorten:  true,
			expected: "/v1/users",
		},
		{
			name:     "Path only",
			reqUrl:   "/api/v1/endpoint",
			shorten:  true,
			expected: "/api/v1/endpoint",
		},
		{
			name:     "Invalid URL not shortened",
			reqUrl:   "not a valid url",
			shorten:  true,
			expected: "not a valid url",
		},
		{
			name:     "Complex path shortened",
			reqUrl:   "https://api.dropbox.com/2/files/list_folder?param=value",
			shorten:  true,
			expected: "/2/files/list_folder",
		},
		{
			name:     "Root path",
			reqUrl:   "https://api.example.com/",
			shorten:  true,
			expected: "/",
		},
		{
			name:     "Empty URL",
			reqUrl:   "",
			shorten:  true,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := UrlFormat(tc.reqUrl, tc.shorten)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s'", tc.expected, result)
			}
		})
	}
}

func TestPopulation_Structure(t *testing.T) {
	pop := &Population{
		Url:        "/api/v1/users",
		Code:       200,
		Population: 100,
		Proportion: 0.75,
	}

	if pop.Url != "/api/v1/users" {
		t.Errorf("Expected Url '/api/v1/users', got '%s'", pop.Url)
	}
	if pop.Code != 200 {
		t.Errorf("Expected Code 200, got %d", pop.Code)
	}
	if pop.Population != 100 {
		t.Errorf("Expected Population 100, got %d", pop.Population)
	}
	if pop.Proportion != 0.75 {
		t.Errorf("Expected Proportion 0.75, got %f", pop.Proportion)
	}
}

func TestLatency_Structure(t *testing.T) {
	lat := &Latency{
		Url:        "/api/v1/users",
		Code:       200,
		Population: 100,
		Mean:       0.125,
		Median:     0.100,
		P50:        0.100,
		P70:        0.150,
		P90:        0.200,
	}

	if lat.Url != "/api/v1/users" {
		t.Errorf("Expected Url '/api/v1/users', got '%s'", lat.Url)
	}
	if lat.Code != 200 {
		t.Errorf("Expected Code 200, got %d", lat.Code)
	}
	if lat.Population != 100 {
		t.Errorf("Expected Population 100, got %d", lat.Population)
	}
	if lat.Mean != 0.125 {
		t.Errorf("Expected Mean 0.125, got %f", lat.Mean)
	}
	if lat.Median != 0.100 {
		t.Errorf("Expected Median 0.100, got %f", lat.Median)
	}
	if lat.P50 != 0.100 {
		t.Errorf("Expected P50 0.100, got %f", lat.P50)
	}
	if lat.P70 != 0.150 {
		t.Errorf("Expected P70 0.150, got %f", lat.P70)
	}
	if lat.P90 != 0.200 {
		t.Errorf("Expected P90 0.200, got %f", lat.P90)
	}
}

func TestTimeSeries_Structure(t *testing.T) {
	ts := &TimeSeries{
		Time:      "2024-01-01T00:00:00Z",
		Url:       "/api/v1/users",
		Code2xx:   100,
		Code3xx:   10,
		Code4xx:   5,
		Code429:   2,
		Code5xx:   1,
		CodeOther: 0,
	}

	if ts.Time != "2024-01-01T00:00:00Z" {
		t.Errorf("Expected Time '2024-01-01T00:00:00Z', got '%s'", ts.Time)
	}
	if ts.Url != "/api/v1/users" {
		t.Errorf("Expected Url '/api/v1/users', got '%s'", ts.Url)
	}
	if ts.Code2xx != 100 {
		t.Errorf("Expected Code2xx 100, got %d", ts.Code2xx)
	}
	if ts.Code3xx != 10 {
		t.Errorf("Expected Code3xx 10, got %d", ts.Code3xx)
	}
	if ts.Code4xx != 5 {
		t.Errorf("Expected Code4xx 5, got %d", ts.Code4xx)
	}
	if ts.Code429 != 2 {
		t.Errorf("Expected Code429 2, got %d", ts.Code429)
	}
	if ts.Code5xx != 1 {
		t.Errorf("Expected Code5xx 1, got %d", ts.Code5xx)
	}
	if ts.CodeOther != 0 {
		t.Errorf("Expected CodeOther 0, got %d", ts.CodeOther)
	}
}

func TestUrlCode_Structure(t *testing.T) {
	uc := &UrlCode{
		ReqUrl:  "/api/v1/users",
		ResCode: 200,
	}

	if uc.ReqUrl != "/api/v1/users" {
		t.Errorf("Expected ReqUrl '/api/v1/users', got '%s'", uc.ReqUrl)
	}
	if uc.ResCode != 200 {
		t.Errorf("Expected ResCode 200, got %d", uc.ResCode)
	}
}

func TestCaptureAggregatorOpts_Apply(t *testing.T) {
	// Test with no options
	opts := CaptureAggregatorOpts{
		TimeIntervalSeconds: 1800,
		Shorten:             false,
	}
	result := opts.Apply([]CaptureAggregatorOpt{})
	if result.TimeIntervalSeconds != 1800 {
		t.Errorf("Expected TimeIntervalSeconds 1800, got %d", result.TimeIntervalSeconds)
	}
	if result.Shorten {
		t.Error("Expected Shorten to be false")
	}

	// Test with single option
	result = opts.Apply([]CaptureAggregatorOpt{OptTimeInterval(7200)})
	if result.TimeIntervalSeconds != 7200 {
		t.Errorf("Expected TimeIntervalSeconds 7200, got %d", result.TimeIntervalSeconds)
	}

	// Test with multiple options
	result = opts.Apply([]CaptureAggregatorOpt{
		OptTimeInterval(900),
		OptShorten(true),
	})
	if result.TimeIntervalSeconds != 900 {
		t.Errorf("Expected TimeIntervalSeconds 900, got %d", result.TimeIntervalSeconds)
	}
	if !result.Shorten {
		t.Error("Expected Shorten to be true")
	}
}

func TestOptTimeInterval(t *testing.T) {
	opts := CaptureAggregatorOpts{}
	optFunc := OptTimeInterval(300)
	result := optFunc(opts)

	if result.TimeIntervalSeconds != 300 {
		t.Errorf("Expected TimeIntervalSeconds 300, got %d", result.TimeIntervalSeconds)
	}
}

func TestOptShorten(t *testing.T) {
	opts := CaptureAggregatorOpts{}
	
	// Test enabling
	optFunc := OptShorten(true)
	result := optFunc(opts)
	if !result.Shorten {
		t.Error("Expected Shorten to be true")
	}

	// Test disabling
	optFunc = OptShorten(false)
	result = optFunc(opts)
	if result.Shorten {
		t.Error("Expected Shorten to be false")
	}
}

func TestNewCaptureAggregator(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatal(err)
		}

		// Test with default options
		ca := NewCaptureAggregator(db, ctl)
		if ca == nil {
			t.Error("Expected non-nil CaptureAggregator")
		}

		// Test with custom options
		ca2 := NewCaptureAggregator(db, ctl, OptTimeInterval(300), OptShorten(true))
		if ca2 == nil {
			t.Error("Expected non-nil CaptureAggregator")
		}

		// Verify it implements the interface
		var _ CaptureAggregator = ca
		var _ CaptureAggregator = ca2
	})
}

func TestCaImpl_BasicFunctionality(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatal(err)
		}

		// Migrate the Record table
		err = db.AutoMigrate(&Record{})
		if err != nil {
			t.Fatal(err)
		}

		ca := NewCaptureAggregator(db, ctl)

		// Test AddById with non-existent job
		err = ca.AddById("non-existent-job")
		if err == nil {
			t.Error("Expected error for non-existent job")
		}

		// Test AddByCliPath with empty path
		err = ca.AddByCliPath("")
		if err == nil {
			t.Error("Expected error for empty cli path")
		}
	})
}

func TestCaImpl_WithTestData(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatal(err)
		}

		// Migrate the Record table
		err = db.AutoMigrate(&Record{})
		if err != nil {
			t.Fatal(err)
		}

		// Add test records
		now := time.Now()
		testRecords := []Record{
			{
				Timestamp: now,
				ReqMethod: "GET",
				ReqUrl:    "/api/v1/users",
				ResCode:   200,
				Latency:   0.100,
			},
			{
				Timestamp: now.Add(1 * time.Second),
				ReqMethod: "GET",
				ReqUrl:    "/api/v1/users",
				ResCode:   200,
				Latency:   0.150,
			},
			{
				Timestamp: now.Add(2 * time.Second),
				ReqMethod: "GET",
				ReqUrl:    "/api/v1/users",
				ResCode:   404,
				Latency:   0.050,
			},
			{
				Timestamp: now.Add(3 * time.Second),
				ReqMethod: "POST",
				ReqUrl:    "/api/v1/users",
				ResCode:   201,
				Latency:   0.200,
			},
			{
				Timestamp: now.Add(4 * time.Second),
				ReqMethod: "GET",
				ReqUrl:    "/api/v1/users",
				ResCode:   429,
				Latency:   0.010,
			},
			{
				Timestamp: now.Add(5 * time.Second),
				ReqMethod: "GET",
				ReqUrl:    "/api/v1/users",
				ResCode:   500,
				Latency:   0.500,
			},
		}

		for _, record := range testRecords {
			err = db.Create(&record).Error
			if err != nil {
				t.Fatal(err)
			}
		}

		ca := NewCaptureAggregator(db, ctl, OptTimeInterval(3600))

		// Test AggregatePopulation
		populationCount := 0
		err = ca.AggregatePopulation(func(r *Population) {
			populationCount++
			if r.Url == "" {
				t.Error("Expected non-empty URL in population aggregate")
			}
			if r.Population <= 0 {
				t.Error("Expected positive population count")
			}
			if r.Proportion < 0 || r.Proportion > 1 {
				t.Errorf("Expected proportion between 0 and 1, got %f", r.Proportion)
			}
		})
		if err != nil {
			t.Errorf("AggregatePopulation failed: %v", err)
		}
		if populationCount == 0 {
			t.Error("Expected at least one population aggregate")
		}

		// Test AggregateLatency
		latencyCount := 0
		err = ca.AggregateLatency(func(r *Latency) {
			latencyCount++
			if r.Url == "" {
				t.Error("Expected non-empty URL in latency aggregate")
			}
			if r.Population <= 0 {
				t.Error("Expected positive population in latency aggregate")
			}
			if r.Mean <= 0 {
				t.Error("Expected positive mean latency")
			}
		})
		if err != nil {
			t.Errorf("AggregateLatency failed: %v", err)
		}
		if latencyCount == 0 {
			t.Error("Expected at least one latency aggregate")
		}

		// Test AggregateTimeSeries
		timeSeriesCount := 0
		err = ca.AggregateTimeSeries(func(r *TimeSeries) {
			timeSeriesCount++
			if r.Url == "" {
				t.Error("Expected non-empty URL in time series")
			}
			if r.Time == "" {
				t.Error("Expected non-empty time in time series")
			}
			// Check that at least one code count is positive
			totalCodes := r.Code2xx + r.Code3xx + r.Code4xx + r.Code429 + r.Code5xx + r.CodeOther
			if totalCodes == 0 {
				t.Error("Expected at least one response code in time series")
			}
		})
		if err != nil {
			t.Errorf("AggregateTimeSeries failed: %v", err)
		}
		if timeSeriesCount == 0 {
			t.Error("Expected at least one time series aggregate")
		}
	})
}

func TestCaImpl_EmptyDatabase(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Fatal(err)
		}

		// Migrate the Record table
		err = db.AutoMigrate(&Record{})
		if err != nil {
			t.Fatal(err)
		}

		ca := NewCaptureAggregator(db, ctl)

		// Test AggregatePopulation with empty database
		called := false
		err = ca.AggregatePopulation(func(r *Population) {
			called = true
		})
		if err != nil {
			t.Errorf("AggregatePopulation should not error on empty database: %v", err)
		}
		if called {
			t.Error("Handler should not be called for empty database")
		}

		// Test AggregateLatency with empty database
		called = false
		err = ca.AggregateLatency(func(r *Latency) {
			called = true
		})
		if err != nil {
			t.Errorf("AggregateLatency should not error on empty database: %v", err)
		}
		if called {
			t.Error("Handler should not be called for empty database")
		}

		// Test AggregateTimeSeries with empty database
		called = false
		err = ca.AggregateTimeSeries(func(r *TimeSeries) {
			called = true
		})
		if err != nil {
			t.Errorf("AggregateTimeSeries should not error on empty database: %v", err)
		}
		if called {
			t.Error("Handler should not be called for empty database")
		}
	})
}

func TestResponseCodeCategorization(t *testing.T) {
	testCases := []struct {
		code     int
		is2xx    bool
		is3xx    bool
		is4xx    bool
		is429    bool
		is5xx    bool
		isOther  bool
	}{
		{200, true, false, false, false, false, false},
		{201, true, false, false, false, false, false},
		{204, true, false, false, false, false, false},
		{301, false, true, false, false, false, false},
		{302, false, true, false, false, false, false},
		{400, false, false, true, false, false, false},
		{404, false, false, true, false, false, false},
		{429, false, false, false, true, false, false},
		{500, false, false, false, false, true, false},
		{503, false, false, false, false, true, false},
		{100, false, false, false, false, false, true},
		{600, false, false, false, false, false, true},
	}

	for _, tc := range testCases {
		// Test categorization logic
		is2xx := tc.code >= 200 && tc.code < 300
		is3xx := tc.code >= 300 && tc.code < 400
		is4xx := tc.code >= 400 && tc.code < 500 && tc.code != 429
		is429 := tc.code == 429
		is5xx := tc.code >= 500 && tc.code < 600
		isOther := tc.code < 200 || tc.code >= 600

		if is2xx != tc.is2xx {
			t.Errorf("Code %d: expected is2xx=%v, got %v", tc.code, tc.is2xx, is2xx)
		}
		if is3xx != tc.is3xx {
			t.Errorf("Code %d: expected is3xx=%v, got %v", tc.code, tc.is3xx, is3xx)
		}
		if is4xx != tc.is4xx {
			t.Errorf("Code %d: expected is4xx=%v, got %v", tc.code, tc.is4xx, is4xx)
		}
		if is429 != tc.is429 {
			t.Errorf("Code %d: expected is429=%v, got %v", tc.code, tc.is429, is429)
		}
		if is5xx != tc.is5xx {
			t.Errorf("Code %d: expected is5xx=%v, got %v", tc.code, tc.is5xx, is5xx)
		}
		if isOther != tc.isOther {
			t.Errorf("Code %d: expected isOther=%v, got %v", tc.code, tc.isOther, isOther)
		}
	}
}