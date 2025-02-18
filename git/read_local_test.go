// Copyright 2022-2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package git

import (
	"github.com/pb33f/openapi-changes/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLocalRepoAvailable(t *testing.T) {
	assert.True(t, CheckLocalRepoAvailable("./"))
	assert.False(t, CheckLocalRepoAvailable("/tmp/made-up"))
}

func TestExtractHistoryFromFile(t *testing.T) {

	c := make(chan *model.ProgressUpdate)
	e := make(chan model.ProgressError)
	d := make(chan bool)
	go func() {
		iterations := 0
		for iterations < 23 {
			select {
			case <-c:
				iterations++
			case <-e:
				iterations++
			}
		}
		d <- true
	}()

	history, _ := ExtractHistoryFromFile("./", "read_local.go", c, e)
	<-d
	assert.NotNil(t, history)
	assert.Equal(t, "adding read_local.go to test repo code", history[len(history)-1].Message)
}

func TestExtractHistoryFromFile_Fail(t *testing.T) {

	c := make(chan *model.ProgressUpdate)
	e := make(chan model.ProgressError)
	d := make(chan bool)
	go func() {
		select {
		case <-c:
			d <- true
		case <-e:
			d <- true
		}
	}()

	history, _ := ExtractHistoryFromFile("./", "no_file_nope", c, e)
	<-d
	assert.Len(t, history, 0)
}

func TestExtractPathAndFile(t *testing.T) {
	var dir, file string
	dir, file = ExtractPathAndFile("/some/place/thing.html")
	assert.Equal(t, "/some/place", dir)
	assert.Equal(t, "thing.html", file)
	dir, file = ExtractPathAndFile("../../thing.html")
	assert.Equal(t, "../..", dir)
	assert.Equal(t, "thing.html", file)
}

func TestGetTopLevel(t *testing.T) {
	str, err := GetTopLevel("./")
	assert.NoError(t, err)
	assert.NotEmpty(t, str)
}
