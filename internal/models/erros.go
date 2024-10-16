package models

import "errors"

var ErrNoRecord = errors.New("models : no matching recored found")
var ErrInsert = errors.New("models : error inserting data")
