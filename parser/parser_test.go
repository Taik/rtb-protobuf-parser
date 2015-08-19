package parser

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var decodeTests = []struct {
	rawB64Encoded string
	expected      string
}{
	{
		`EhBV1MzFAAuRQwofLodLCJL1IgNCV1IysgFNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgNC40LjI7IFNNLVQyMTdTIEJ1aWxkL0tPVDQ5SCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzMwLjAuMC4wIFNhZmFyaS81MzcuMzYgKE1vYmlsZTsgYWZtYS1zZGstYS12Nzg5OTAwMC43NTcxMDAwLjEpWjlodHRwczovL3BsYXkuZ29vZ2xlLmNvbS9zdG9yZS9hcHBzL2RldGFpbHM/aWQ9bW9iaS5pZnVubnliAmVuaggItgEVAACAPmoICNUIFQAAgD5qCAi1CxUAAIA+agcIAxUAAIA+cuABCAEQwAIYMiIFRhweIhYyiQEKHCorMzw/XF5xfoABggGQAZEBkgGcAbMBtgHGAcwB4QHnAegB6QHsAe0B7gHyAf8BhAKLAp0CrwK7AsMCxQLLAswCzgLPAtYCngOwA7kDvQPYA9oD3APdA+AD4QPlA+YD6QPqA/ED+AORBJYEmQSeBJ8EpgSnBKwEsgS0BLYEvAS9BL4EvwTBBEoQEKeihtpSKPDQSDoEMICXIlIWcGFjay1icmFuZC1pRnVubnk6OkFsbGABcAF51qnrNLKcQE55xoAhVeTAJB2IAQCQAQB4AKABAaoBG0NBRVNFS3FqZDUwQzdjdG1Bay12NWhuSEpjRcABAsgBkP7/////////AdIBASriAYoBGgdhbmRyb2lkMgttb2JpLmlmdW5ueTgBQAJIAVjm1ANY4NQDWPyeBGIHc2Ftc3VuZ2oIc20tdDIxN3NyBAgEEAR42ASAAYAIkAHLowSYAegHogEkVdTMxQALkUMKHy6HAQiS9SoVfK7W4IWzlHzZ8XU8k5VktGR7wgEJaUZ1bm55IDopzQE+EItA+AGAjokBigIFMjAwMDG4AqXjpQTIArow0QJJB8psWzWCWQ==`,
		`{"id":"VdTMxQALkUMKHy6HSwiS9Q==","ip":"QldS","google_user_id":"CAESEKqjd50C7ctmAk-v5hnHJcE","cookie_version":1,"cookie_age_seconds":2246400,"user_agent":"Mozilla/5.0 (Linux; Android 4.4.2; SM-T217S Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Safari/537.36 (Mobile; afma-sdk-a-v7899000.7571000.1)","geo_criteria_id":9007525,"postal_code":"20001","timezone_offset":-240,"seller_network_id":6202,"url":"https://play.google.com/store/apps/details?id=mobi.ifunny","detected_language":["en"],"detected_vertical":[{"id":182,"weight":0.25},{"id":1109,"weight":0.25},{"id":1461,"weight":0.25},{"id":3,"weight":0.25}],"vertical_dictionary_version":2,"detected_content_label":[42],"mobile":{"platform":"android","brand":"samsung","model":"sm-t217s","os_version":{"os_version_major":4,"os_version_minor":4},"carrier_id":70091,"is_app":true,"app_id":"mobi.ifunny","mobile_device_type":2,"screen_width":600,"screen_height":1024,"screen_orientation":1,"app_category_ids":[60006,60000,69500],"device_pixel_ratio_millis":1000,"encrypted_advertising_id":"VdTMxQALkUMKHy6HAQiS9SoVfK7W4IWzlHzZ8XU8k5VktGR7","app_name":"iFunny :)","app_rating":4.3457327},"publisher_settings_list_id":6449776283131447113,"adslot":[{"id":1,"ad_block_key":1,"targetable_channel":["pack-brand-iFunny::All"],"width":[320],"height":[50],"excluded_attribute":[70,28,30,34,22],"allowed_vendor_type":[10,28,42,43,51,60,63,92,94,113,126,128,130,144,145,146,156,179,182,198,204,225,231,232,233,236,237,238,242,255,260,267,285,303,315,323,325,331,332,334,335,342,414,432,441,445,472,474,476,477,480,481,485,486,489,490,497,504,529,534,537,542,543,550,551,556,562,564,566,572,573,574,575,577],"matching_ad_data":[{"adgroup_id":22200553767,"minimum_cpm_micros":1190000,"pricing_rule":[{"minimum_cpm_micros":560000}]}],"publisher_settings_list_id":[5638679022673832406,2100015413174829254],"slot_visibility":1}],"is_test":false}`,
	},
}

func TestDecode(t *testing.T) {
	for _, entry := range decodeTests {
		raw, _ := base64.StdEncoding.DecodeString(entry.rawB64Encoded)
		actual, err := Decode(raw)
		assert.Nil(t, err)
		assert.Equal(t, entry.expected, string(actual), "Unexpected Decode() result")
	}
}

func BenchmarkDecode(b *testing.B) {
	raw, _ := base64.StdEncoding.DecodeString("EhBV1MzFAAuRQwofLodLCJL1IgNCV1IysgFNb3ppbGxhLzUuMCAoTGludXg7IEFuZHJvaWQgNC40LjI7IFNNLVQyMTdTIEJ1aWxkL0tPVDQ5SCkgQXBwbGVXZWJLaXQvNTM3LjM2IChLSFRNTCwgbGlrZSBHZWNrbykgVmVyc2lvbi80LjAgQ2hyb21lLzMwLjAuMC4wIFNhZmFyaS81MzcuMzYgKE1vYmlsZTsgYWZtYS1zZGstYS12Nzg5OTAwMC43NTcxMDAwLjEpWjlodHRwczovL3BsYXkuZ29vZ2xlLmNvbS9zdG9yZS9hcHBzL2RldGFpbHM")
	for n := 0; n < b.N; n++ {
		Decode(raw)
	}
}
