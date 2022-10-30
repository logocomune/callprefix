package callprefix

import "strings"

func Info(callsign string) (Prefix, bool) {

	callsign = strings.ToUpper(strings.TrimSpace(callsign))

	if info, ok := prefixMapping[callsign]; ok && info.IsFullCallsign {
		return info, true
	}

	for _, size := range prefixSizes {
		if len(callsign) < size {
			continue
		}
		p := callsign[0:size]
		if info, ok := prefixMapping[p]; ok {
			return info, true
		}
	}

	return Prefix{}, false
}
