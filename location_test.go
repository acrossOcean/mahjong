package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocation_String(t *testing.T) {
	var east, south, west, north, nowhere Location
	east = LocationEast
	south = LocationSouth
	west = LocationWest
	north = LocationNorth

	assert.Equal(t, fmt.Sprint(east), "东")
	assert.Equal(t, fmt.Sprint(south), "南")
	assert.Equal(t, fmt.Sprint(west), "西")
	assert.Equal(t, fmt.Sprint(north), "北")
	assert.Equal(t, fmt.Sprint(nowhere), "未知")
}
