package es_sort

import (
	"bytes"
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_close"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"
)

func testSortData(t *testing.T, digits, numLines, numDupLines int, seed int64, desc, uniq bool, sorter Sorter, out *bytes.Buffer) {
	// Prepare test data
	l := esl.Default()
	l.Debug("Conditions",
		esl.Int64("seed", seed),
		esl.Int("digits", digits),
		esl.Int("numLines", numLines),
		esl.Int("numDupLines", numDupLines),
	)
	shuffler := rand.New(rand.NewSource(seed))

	dataFormat := "%0" + strconv.FormatInt(int64(digits), 10) + "d"
	totalLines := numLines + numDupLines
	data0 := make([]string, numLines)
	data := make([]string, totalLines)
	for i := 0; i < numLines; i++ {
		data0[i] = fmt.Sprintf(dataFormat, i)
		data[i] = data0[i]
	}
	for i := numLines; i < totalLines; i++ {
		data[i] = fmt.Sprintf(dataFormat, shuffler.Intn(numLines))
	}

	sort.Slice(data, func(i, j int) bool {
		if desc {
			return strings.Compare(data[i], data[j]) > 0
		} else {
			return strings.Compare(data[i], data[j]) < 0
		}
	})
	sort.Slice(data0, func(i, j int) bool {
		if desc {
			return strings.Compare(data0[i], data0[j]) > 0
		} else {
			return strings.Compare(data0[i], data0[j]) < 0
		}
	})

	var expectedLines int
	var expected []string
	if uniq {
		expectedLines = numLines
		expected = make([]string, expectedLines)
		copy(expected[:], data0[:])
	} else {
		expectedLines = numLines + numDupLines
		expected = make([]string, expectedLines)
		copy(expected[:], data[:])
	}

	shuffler.Shuffle(numLines, func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	// Add to sorter
	for _, line := range data {
		if wrErr := sorter.WriteLine(line); wrErr != nil {
			t.Error(wrErr)
			return
		}
	}

	// flush
	if flErr := sorter.Flush(); flErr != nil {
		t.Error(flErr)
	}

	// verify data
	result := strings.Split(out.String(), "\n")

	if len(result) != expectedLines+1 {
		t.Error(len(result))
	}
	if result[len(result)-1] != "" {
		t.Error(result[len(result)-1])
	}

	for i := 0; i < expectedLines; i++ {
		if result[i] != expected[i] {
			t.Error(i, result[i], expected[i], len(result), len(expected))
		}
	}
}

func testSimple(t *testing.T, data, expected []string, sorter Sorter, out *bytes.Buffer) {
	for _, line := range data {
		if err := sorter.WriteLine(line); err != nil {
			t.Error(err)
			return
		}
	}

	if flErr := sorter.Flush(); flErr != nil {
		t.Error(flErr)
	}

	result := strings.Split(out.String(), "\n")
	if len(result) != len(expected)+1 {
		t.Error(len(result), len(expected))
	}
	if result[len(result)-1] != "" {
		t.Error(result[len(result)-1])
	}

	for i := 0; i < len(expected); i++ {
		if result[i] != expected[i] {
			t.Error(i, result[i], expected[i])
		}
	}
}

func TestOwnComparator(t *testing.T) {
	data := []string{
		"1",
		"10",
		"13",
		"2",
		"21",
	}
	seed := time.Now().Unix()
	l := esl.Default()
	l.Debug("Seed", esl.Int64("seed", seed))
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})

	// Default comparator
	{
		out := &bytes.Buffer{}
		expected := []string{
			"1",
			"10",
			"13",
			"2",
			"21",
		}
		sorter := New(es_close.NewNopWriteCloser(out), Logger(esl.Default()))
		t.Run("default", func(st *testing.T) {
			testSimple(st, data, expected, sorter, out)
		})
	}

	// Numeric comparator
	{
		out := &bytes.Buffer{}
		expected := []string{
			"1",
			"2",
			"10",
			"13",
			"21",
		}
		sorter := New(es_close.NewNopWriteCloser(out), Logger(esl.Default()), Comparator(func(x, y string) int {
			xv, err := strconv.ParseInt(x, 10, 32)
			if err != nil {
				t.Error(x, err)
			}
			yv, err := strconv.ParseInt(y, 10, 32)
			if err != nil {
				t.Error(y, err)
			}
			if xv < yv {
				return -1
			}
			if yv < xv {
				return 1
			}
			return 0
		}))
		t.Run("own", func(st *testing.T) {
			testSimple(st, data, expected, sorter, out)
		})
	}
}

func TestOwnTempFolder(t *testing.T) {
	tf, err := ioutil.TempDir("", "sort")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(tf)
	}()
	out := &bytes.Buffer{}
	sorter := New(es_close.NewNopWriteCloser(out), TempFolder(tf))
	testSortData(t, 10, 100, 5, 0, false, false, sorter, out)
}

func TestWithLogger(t *testing.T) {
	out := &bytes.Buffer{}
	sorter := New(es_close.NewNopWriteCloser(out), Logger(esl.Default()))
	testSortData(t, 10, 100, 5, 0, false, false, sorter, out)
}

func TestDupLines(t *testing.T) {
	out := &bytes.Buffer{}
	sorter := New(es_close.NewNopWriteCloser(out))
	testSortData(t, 10, 100, 5, 0, false, false, sorter, out)
}

func TestConditions(t *testing.T) {
	condOrder := []bool{false, true}
	condUniq := []bool{false, true}
	condCompress := []bool{false, true}
	condDupLines := []int{0, 500}

	for _, cOrder := range condOrder {
		for _, cUniq := range condUniq {
			for _, cCompress := range condCompress {
				for _, cDupLines := range condDupLines {
					seed := time.Now().Unix()
					name := fmt.Sprintf("Order[%t] Uniq[%t] Compress[%t] DupLines[%d] Seed[%d]", cOrder, cUniq, cCompress, cDupLines, seed)
					t.Run(name, func(st *testing.T) {
						{
							out := &bytes.Buffer{}
							sorter := New(es_close.NewNopWriteCloser(out), Desc(cOrder), Uniq(cUniq), TempCompress(cCompress))
							testSortData(st, 10, 1000, cDupLines, seed, cOrder, cUniq, sorter, out)
						}
						{
							out := &bytes.Buffer{}
							sorter := New(es_close.NewNopWriteCloser(out), Desc(cOrder), Uniq(cUniq), TempCompress(cCompress), MemoryLimit(1000))
							testSortData(st, 10, 1000, cDupLines, seed, cOrder, cUniq, sorter, out)
						}
					})
				}
			}
		}
	}
}

func TestAltBehavior(t *testing.T) {
	{
		outData := &bytes.Buffer{}
		out := es_close.NewNopWriteCloser(outData)
		s := New(out)

		if err := s.Flush(); err != nil {
			t.Error(err)
		}

		// 2nd flush should not return err
		if err := s.Flush(); err != nil {
			t.Error(err)
		}
	}

	{
		outData := &bytes.Buffer{}
		out := es_close.NewNopWriteCloser(outData)
		s := New(out)

		if err := s.Flush(); err != nil {
			t.Error(err)
		}

		// should raise an exception
		if err := s.WriteLine("test"); err != ErrorAlreadyClosed {
			t.Error(err)
		}
	}
}
