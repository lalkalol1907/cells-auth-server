package CustomErrors

import "errors"

var NeedRefreshError = errors.New("Need Refresh")

var NoSession = errors.New("No session")
