package ip

import (
	"net"
	"testing"
	"time"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/stretchr/testify/assert"
)

func TestGetLocalIP(t *testing.T) {
	ip, err := GetLocalIP()
	if err != nil {
		t.Logf("GetLocalIP() returned error: %v", err)
		if ip != "" {
			t.Errorf("Expected empty IP string when error is not nil, but got: %s", ip)
		}
	} else {
		t.Logf("GetLocalIP() returned IP: %s", ip)
		if ip == "" {
			t.Errorf("Expected non-empty IP string when error is nil, but got empty string")
		}
		if net.ParseIP(ip) == nil {
			t.Errorf("Expected valid IP string, but got: %s", ip)
		}
	}
}

func TestIP(t *testing.T) {
	var dbPath = "./ip2region.xdb"
	searcher, err := xdb.NewWithFileOnly(dbPath)
	if err != nil {
		t.Fatalf("failed to create searcher: %s", err.Error())
	}

	defer searcher.Close()

	// do the search
	var ip = "1.2.3.4"
	var tStart = time.Now()
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		t.Fatalf("failed to SearchIP(%s): %s", ip, err)
	}

	t.Logf("{region: %s, took: %s}\n", region, time.Since(tStart))
	expectedRegion := "美国|0|华盛顿|0|谷歌"
	assert.Equal(t, expectedRegion, region, "Region does not match expected value")

	// 备注：并发使用，每个 goroutine 需要创建一个独立的 searcher 对象。
}
