package httpcord

import (
	"math"
	"strconv"
	"time"
)

type Snowflake string

const DiscordEpoch = 1420070400000

func (s Snowflake) CreatedAt() Time {
	timestamp := (int64(s.Uint64()) >> 22) + DiscordEpoch
	return Time{time.Unix(0, timestamp*int64(time.Millisecond))}
}

func (s Snowflake) Uint64() uint64 {
	r, _ := strconv.ParseUint(s.String(), 10, 64)

	return r
}

func (s Snowflake) String() string {
	return string(s)
}

func (s Snowflake) Valid(id Snowflake) bool {
	return validSnowflake(id)
}

// Epoch Get epoch of id
func Epoch(id Snowflake) float64 {
	if !validSnowflake(id) {
		panic("Invalid snowflake for Epoch")
	}

	r, _ := strconv.ParseFloat(string(id), 64)
	return math.Floor(r / 4194304)
}

func validSnowflake(snowflake Snowflake) bool {
	if _, err := strconv.ParseUint(string(snowflake), 10, 0); err != nil {
		return false
	}

	return true
}

func StringArrayToSnowflakeArray(arr []string) []*Snowflake {
	var snowflakeArr []*Snowflake

	for _, snowflake := range arr {
		s := Snowflake(snowflake)
		snowflakeArr = append(snowflakeArr, &s)
	}

	return snowflakeArr
}
