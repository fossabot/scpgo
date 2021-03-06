// Copyright © 2017 Ricardo Aravena <raravena@branch.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scp

import (
	"fmt"
	"io"
	"os"
	"time"
)

// DEFAULTFORMAT for progressbar
const DEFAULTFORMAT = "\r%s   % 3d %%  %d kb %0.2f kb/s %v      "

// ProgressBar Struct for Progress Bar
type ProgressBar struct {
	Out       io.Writer
	Format    string
	Subject   string
	StartTime time.Time
	Size      int64
}

// NewProgressBarTo Instantiatiates a new Progress Bar To
func NewProgressBarTo(subject string, size int64, outPipe io.Writer) ProgressBar {
	return ProgressBar{outPipe, DEFAULTFORMAT, subject, time.Now(), size}
}

// NewProgressBar Instantiatiates a new Progress Bar
func NewProgressBar(subject string, size int64) ProgressBar {
	return NewProgressBarTo(subject, size, os.Stdout)
}

// Update Updates the Progress Bar
func (pb ProgressBar) Update(tot int64) {
	percent := int64(0)
	if pb.Size > int64(0) {
		percent = (int64(100) * tot) / pb.Size
	}
	totTime := time.Now().Sub(pb.StartTime)
	spd := float64(tot/1000) / totTime.Seconds()
	//TODO put kb size into format string
	fmt.Fprintf(pb.Out, pb.Format, pb.Subject, percent, tot, spd, totTime)

}
