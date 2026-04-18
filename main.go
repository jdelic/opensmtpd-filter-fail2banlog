package main

import (
	"log"
	"os"

	"github.com/jdelic/opensmtpd-filters-go"
)

type Fail2banFilter struct {
	opensmtpd.SessionTrackingMixin
}

func (f *Fail2banFilter) GetName() string {
	return "opensmtpd fail2ban auth reporter"
}

// link-auth params according to the framework's session tracker:
//
//	params[0] = username
//	params[1] = result
//
// The mixin treats result == "pass" as success and stores the username only then.
// For fail2ban we only care about non-pass results.
func (f *Fail2banFilter) LinkAuth(fw opensmtpd.FilterWrapper, ev opensmtpd.FilterEvent) {
	params := ev.GetParams()
	if len(params) != 2 {
		log.Printf("opensmtpd-f2b: malformed link-auth session=%s params=%q",
			ev.GetSessionId(), params)
		return
	}

	username := params[0]
	result := params[1]

	if result == "pass" {
		return
	}

	ip := "unknown"
	rdns := ""

	if s := f.GetSession(ev.GetSessionId()); s != nil {
		if s.SrcIp != "" {
			ip = s.SrcIp
		}
		rdns = s.Rdns
	}

	// Single stable line for fail2ban to match.
	log.Printf("opensmtpd-f2b: auth-failure rip=%s user=%q rdns=%q result=%q session=%s",
		ip, username, rdns, result, ev.GetSessionId())
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	filter := opensmtpd.NewFilter(&Fail2banFilter{})
	opensmtpd.Run(filter)
}
