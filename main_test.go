package main

import (
	"testing"
	//"strings"
	//"github.com/stretchr/testify/assert"
	"fmt"
	//"database/sql"
	//"reflect"
	//"io/ioutil"
	//"bytes"
	//"sync"
	//"os"
	"testing/quick"
)

func TestTruth(t *testing.T)  {
	if true != true {
		t.Error("Everything else is wrong")
	}
}

func TestBroken(t *testing.T) {

	if 1 + 1 != 3 {
		t.Fatal("Can't proceed!")
	}
	if true != false {
		t.Error("expected", true, "got", false)
	}
}

func TestIndexRequestUnmarshalJSON(t *testing.T) {
	json := `{options:`

	t.Run("empty json", func(t *testing.T) {
		a = a + `{}`
		fmt.Println(json)
	})
	t.Run("int value", func(t *testing.T) {
		a = a + `4`
		fmt.Println(json)
	})
}


//func assertEqual(x, y int)  {
//	if x != y {
//		panic(fmt.Sprintf("%d != %d", x, y))
//	}
//}
//
//func TestSplit(t *testing.T)  {
//	words := strings.Split("a,b,c", ",")
//	assertEqual(len(words), 3)
//}

//
//func TestSplit(t *testing.T) {
//	str, sep := "a,b,c", ","
//	words := strings.Split(str, sep)
//	if got, want := len(words), 3; got != want {
//		t.Error("split (%s, %s)expected: %s, actual: %s",
//			str, sep, got, want)
//	}
//}
//
//
//func Order(userInfo string) bool {
//	accessToken := doOauth(userInfo)
//	payment := getPayment(accessToken)
//}
//
//func processCreditCard(cardInfo string) bool {
//
//}
//
//type StatsClient {
//	host String
//}
//
//func (stats StatsClient) Count(name string, value int64, rate float64) {...}
//
//func (f *Fragment) setBit(rowID, columnID uint64) error {
//	// Return row from cache or create a new row
//	bm := f.row(rowID, true, true)
//	err := bm.SetBit(columnID)
//	if err != nil {
//		return err
//	}
//	stats := NewStatsClient(host)
//	stats.Count("setBit", 1, 0.001)
//
//	return nil
//}
//
//func (f *Fragment) setBit(rowID, columnID uint64, stat *StatsClient) error {
//	// Return row from cache or create a new row
//	bm := f.row(rowID, true, true)
//	err := bm.SetBit(columnID)
//	if err != nil {
//		return err
//	}
//	stat.Count("setBit", 1, 0.001)
//	return nil
//}
//
//
//func TestPostIndexRequestUnmarshalJSON(t *testing.T) {
//	t.Run("empty json", func(t *testing.T) {
//		json := `{"options": {}}`
//		err := json.Unmarshal([]byte(test.json), &postIndexRequest{})
//		if err != nil {
//			t.Errorf(err)
//		}
//	})
//
//	t.Run("int value", func(t *testing.T) {
//		json := `{"options": 4}`
//		err := json.Unmarshal([]byte(test.json), &postIndexRequest{})
//		if err != "options is not map[string]interface{}" {
//			t.Errorf(err)
//		}
//	})
//}
//func TestIndexRequestUnmarshalJSON(t *testing.T) {
//	tests := []struct {
//		json     string
//		expected IndexRequest
//		err      string
//	}{
//		{json: `{"options": {}}`, expected: IndexRequest{Options: IndexOptions{}}},
//		{json: `{"options": 4}`, err: "options is not map[string]interface{}"},
//		{json: `{"option": {}}`, err: "Unknown key: option:map[]"}
//	}
//	for _, test := range tests {
//		t.Run(fmt.Sprint(test.json), func(t *testing.T) {
//			err := json.Unmarshal([]byte(test.json), &IndexRequest{})
//			if err != nil {
//				t.Errorf("option: %s, expected error: %v, but got result: %v", test.json, test.err, err)
//			}
//		}
//	}
//}

//func createTempFile() string {
//	tf, err := ioutil.TempFile("", "testFile")
//	if err != nil {
//		panic(err)
//	}
//	tf.Close()
//	return tf.Name()
//}

//func createTempFile(t *testing.T) (string, error) {
//	tf, err := ioutil.TempFile("", "testFile")
//	if err != nil {
//		t.Fatalf("err: %s", err)
//	}
//	tf.Close()
//	return tf.Name(), nil
//}
//
//func TestTmp(t *testing.T) {
//	name, err := createTempFile(t)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(name)
//}
//
//func TestBackupCommand_Run(t *testing.T) {
//	s := test.NewServer()
//	defer s.Close()
//
//	cm := NewBackupCommand(stdin, stdout, stderr)
//	fileName := createTempFile(t)
//
//	err := cm.Run(context.Background())
//	if err != nil {
//		t.Fatalf("Command not working, error: '%s'", err)
//	}
//}
//
//type MockStats struct {
//	mockCount    func(name string, value int64, rate float64)
//}
//
//func (s *MockStats) Count(name string, value int64, rate float64) {
//	if s.mockCount != nil {
//		s.mockCount(name, value, rate)
//		return
//	}
//	return
//}
//
//func TestStatsCount_SetBit(t *testing.T) {
//	s := test.NewServer()
//	defer s.Close()
//	stats := &MockStats{
//		mockCount: func(name string, value int64, rate float64) {
//			if name != "setBit" {
//				t.Errorf("Expected setBit, Results %s", name)
//			}
//			called = true
//			return
//		},
//	}
//	http.DefaultClient.Do(test.MustNewHTTPRequest("POST", s.URL+"/index/i/frame/f", strings.NewReader("SetBit()")))
//	if !called {
//		t.Error("Count isn't called")
//	}
//}
//
//type Fragment struct {
//
//	index string
//	frame string
//	view  string
//	slice uint64
//
//	stats StatsClient
//}
//
// Ensure generated view name can be returned for a given time unit.
func TestViewByTimeUnit(t *testing.T) {
	ts := time.Date(2000, time.January, 2, 3, 4, 5, 6, time.UTC)

	t.Run("Y", func (t *testing.T) {
		if s := pilosa.ViewByTimeUnit("F", ts, 'Y'); s != "F_2000" {
			t.Fatalf("unexpected name: %s", s)
		}
	})
	t.Run("M", func (t *testing.T) {
		if s := pilosa.ViewByTimeUnit("F", ts, 'M'); s != "F_200001" {
			t.Fatalf("unexpected name: %s", s)
		}
	})
	t.Run("D", func (t *testing.T) {
		if s := pilosa.ViewByTimeUnit("F", ts, 'D'); s != "F_20000102" {
			t.Fatalf("unexpected name: %s", s)
		}
	})
}


// Ensure a fragment can set & read a field value.
func TestFragment_SetFieldValue(t *testing.T) {

	t.Run("NotExists", func(t *testing.T) {
		f := test.MustOpenFragment("i", "f", pilosa.ViewStandard, 0, "")
		defer f.Close()

		// Set value.
		f.SetFieldValue(100, 10, 20)

		// Non-existant value.
		if value, exists, err := f.FieldValue(100, 11); err != nil {
			t.Fatal(err)
		} else if value != 0 {
			t.Fatalf("unexpected value: %d", value)
		} else if exists {
			t.Fatal("expected to not exist")
		}
	})
}

	t.Run("QuickCheck", func(t *testing.T) {
		if err := quick.Check(func(bitDepth uint, columnN uint64, values []uint64) bool {
			// Limit bit depth & maximum values.
			bitDepth = (bitDepth % 62) + 1
			columnN = (columnN % 99) + 1

			for i := range values {
				values[i] = values[i] % (1 << bitDepth)
			}

			f := test.MustOpenFragment("i", "f", pilosa.ViewStandard, 0, "")
			defer f.Close()

			// Set values.
			m := make(map[uint64]int64)
			for _, value := range values {
				columnID := value % columnN

				m[columnID] = int64(value)

				if _, err := f.SetFieldValue(columnID, bitDepth, value); err != nil {
					t.Fatal(err)
				}
			}

			// Ensure values are set.
			for columnID, value := range m {
				v, exists, err := f.FieldValue(columnID, bitDepth)
				if err != nil {
					t.Fatal(err)
				} else if value != int64(v) {
					t.Fatalf("value mismatch: column=%d, bitdepth=%d, value: %d != %d", columnID, bitDepth, value, v)
				} else if !exists {
					t.Fatalf("value should exist: column=%d", columnID)
				}
			}

			return true
		}, nil); err != nil {
			t.Fatal(err)
		}
	})
}

