package test

import (
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"testing"
)

func TestCheckToken(t *testing.T) {
	jwt := "ZXlKaGJHY2lPaUpTVTBFdFQwRkZVQ0lzSW1WdVl5STZJa0V4TWpoSFEwMGlmUS5SOHFPUnViYkFfY1AwRENRcXVSRWgwb0Q5Rjl1Rms1ZDlxWGNPSXp6Nmk0MS1iVm9EQXNDM2VpSEVlVS1jT2I3Wl9yUkRDSFBadFpUVHY0RGJ4SnhfRU94OUxGOXFfUDV0VDBQQl9TeFhiMzFkOWRvdEx4WDYxcEcxbDlSMW01T2dvOHBGWVc1RU9oamdvY1g2UVlPTGRiQ2xpa2gyQ1lkUG5FVEt0Y2gtTU5QWE03UDdneWQ2eXNBa1JIdWtJRFNnRXFGQzY4bXRid0dndzdidVVERk9jaTB1QmxkRldPOHZHRm5WS04wWENaVnNJNE9fbTJrN08yek9ISlVUQ21ncDcwN1pEZHZXWnBhY3MydzlRa2Mta0d3eXd3U0hKN25aLUVUMUVBM2pjWXJMZEt1RlNpRi1zLV90YXFBRWgtU1RxSVNNc3JPUk5BY0kzMEdmeE5taFEucWNKb1g5YnpXLTZsYlBWNi5IWktJTkxXT2F0X0FtZlY2OFNnTXJtUUxNUkt2QUsydGF0OTdCUngxdTNhNVpaWjBnLXg1dWtlYzJPQm9vTlgxaEVUSWNFVTRHcFBZdUZ4NzBDWXlwYTloU0FBMTNKTVFMb1FUYWxmUm9RcDlLTFFpZGtzVkFuNUJTdmc0eURqLUZBY2hTTlJPS1ZwQ3dqaUdaUW9UUHdpb3MzOTFHd3FpMGkwdEVGRExBWGxpUkN3WFkxb1NvNS1nT1pDcFpzUWZKbXBkVWhMVHQ3YUF4SVFfRFJmOVlIV2N2ZWw0VzF3VjJ2b3pxdGlrckcwbE9BUHhFd3V1S0FqWFFrY3doc2wxTjltVDRzQjlrU2NUcUE1bTd6U1ZOOXdobGFyRkNlZXVHY3czUmdZVDY4TW5Ua240OW5MS2VWLTJVNVRKQkVMTFEtU1l5Y01TQ3A3eXR6ellTa1I3SmZUT2wxeEx2Nm9iYUk2U0F1S05wVGVfVUdWdk1aNnRVRU92T3V5RnZTVnl5U1UzRW9GR0xEbXREOTAueEVJVHBZdFFpX0hJMEtEVVM5a3Bxdw=="

	s := subChecker.NewSubChecker("localhost:8080")
	token := "262403d5e1330d63693919f10584e353b5af9e5762acc7761d0f16d0e240b350"
	apiName := "youtube-download"
	err := s.CheckToken(token, jwt, apiName)
	if err != nil {
		t.Error(err)
	}
}
