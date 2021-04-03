package crotal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func removeConfigFile() {
	_ = os.Remove("crotal.json")
}

func TestCrotalHandleMissingConf(t *testing.T) {
	removeConfigFile()
	account := "012345678910"
	ret, err := Crotal(account)
	if ret != -1 {
		t.Errorf("Crotal(%s) = %d; want -1", account, ret)
	}
	if err == nil {
		t.Errorf("Crotal(%s) = %d; want %s", account, ret, err)
	}
}

func TestCrotalHandleMissingKeys(t *testing.T) {
	data := map[string]int{
		"012345678910": 5,
		"121231231231": 1,
		"534531231212": 3,
	}
	f, _ := json.Marshal(data)
	_ = ioutil.WriteFile("crotal.json", f, 0644)

	defer removeConfigFile()

	var accounts = []struct {
		account string
		limit   int
		want    int
	}{
		{"1", 5, -1},
		{"2", 1, -1},
		{"3", 3, -1},
		{"0", 5, -1},
	}

	for _, tt := range accounts {
		testname := fmt.Sprintf("%s,%d", tt.account, tt.limit)
		t.Run(testname, func(t *testing.T) {
			ans, _ := Crotal(tt.account)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestCrotalProcessRequests(t *testing.T) {
	data := map[string]int{
		"112345678910": 1,
		"212345678910": 4,
		"312345678910": 6,
		"412345678910": 9,
	}
	f, _ := json.Marshal(data)
	_ = ioutil.WriteFile("crotal.json", f, 0644)

	defer removeConfigFile()

	var accounts = []struct {
		account string
		limit   int
		want    int
	}{
		{"112345678910", 1, 1},
		{"212345678910", 4, 4},
		{"312345678910", 6, 6},
		{"412345678910", 9, 9},
	}

	for _, tt := range accounts {
		testname := fmt.Sprintf("%s,%d", tt.account, tt.limit)
		t.Run(testname, func(t *testing.T) {
			ans, _ := Crotal(tt.account)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
