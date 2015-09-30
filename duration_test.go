package duration

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseString(t *testing.T) {
	t.Parallel()

	// test with bad format
	_, err := ParseString("asdf")
	assert.Equal(t, err, ErrBadFormat)

	// test with good full string
	dur, err := ParseString("P1Y2M3DT4H5M6S")
	assert.Nil(t, err)
	assert.Equal(t, float64(1), dur.Years())
	assert.Equal(t, float64(30*2+3), dur.Days())
	assert.Equal(t, float64(4), dur.Hours())
	assert.Equal(t, float64(5), dur.Minutes())
	assert.Equal(t, float64(6), dur.Seconds())

	// test with good week string
	dur, err = ParseString("P1W")
	assert.Nil(t, err)
	assert.Equal(t, float64(1), dur.Weeks())
}

func TestString(t *testing.T) {
	t.Parallel()

	// test empty
	d := Duration{}
	assert.Equal(t, d.String(), "P")

	// test only larger-than-day
	p, _ := time.ParseDuration("10272h")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1Y63D")

	// test only smaller-than-day
	p, _ = time.ParseDuration("1h2m3s")
	d = Duration{p}
	assert.Equal(t, d.String(), "PT1H2M3S")

	// test full format
	p, _ = time.ParseDuration("10276h5m6s")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1Y63DT4H5M6S")

	// test week format
	p, _ = time.ParseDuration("168h")
	d = Duration{p}
	assert.Equal(t, d.String(), "P1W")
}

func TestToDuration(t *testing.T) {
	t.Parallel()

	p, _ := time.ParseDuration("8760h")
	d := Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*365)

	p, _ = time.ParseDuration("720h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*30)

	p, _ = time.ParseDuration("168h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24*7)

	p, _ = time.ParseDuration("24h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour*24)

	p, _ = time.ParseDuration("1h")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Hour)

	p, _ = time.ParseDuration("1m")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Minute)

	p, _ = time.ParseDuration("1s")
	d = Duration{p}
	assert.Equal(t, d.ToDuration(), time.Second)
}

func TestJSON(t *testing.T) {
	t.Parallel()

	// test JSON Marshal
	dur, _ := ParseString("P1Y2M3DT4H5M6S")
	bytes, err := json.Marshal(dur)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// test JSON Unmarshal
	dur = nil
	err = json.Unmarshal(bytes, &dur)
	assert.Nil(t, err)
	assert.Equal(t, "P1Y63DT4H5M6S", dur.String())
}
