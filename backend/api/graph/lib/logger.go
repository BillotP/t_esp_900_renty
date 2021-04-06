package lib

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/getsentry/sentry-go"
)

// TMSTP is ...
var TMSTP = time.Now().Format(time.RFC3339)

func init() {
	// Init Sentry.io errors reporting service
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         ServerConf.SentryDSN,
		DebugWriter: os.Stderr,
		Debug:       true,
		Environment: ServerConf.Stage,
		Release:     ServerConf.Release,
		SampleRate:  0.5,
	}); err != nil {
		LogError("main", err.Error())
		os.Exit(1)
	}
}

// LogInfo log info with timestamp on stdout
func LogInfo(orig string, message string) string {
	logmsg := fmt.Sprintf("[%s] 💡 Info(%s): %s\n", TMSTP, orig, message)
	fmt.Fprint(os.Stdout, logmsg)
	return logmsg
}

// LogError log error with timestamp on stderr
func LogError(orig string, message string) string {
	logmsg := fmt.Sprintf("[%s] 🚨  Error(%s): %s\n", TMSTP, orig, message)
	if ServerConf.Stage != "dev" &&
		!strings.Contains(message, "Token is expired") {
		sentry.CaptureMessage(logmsg)
	}
	fmt.Fprint(os.Stderr, logmsg)
	return logmsg
}
