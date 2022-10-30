package callprefix

type Prefix struct {
	CountryName    string
	CQZone         int
	ITUZone        int
	Continent      string
	Lat            float64
	Lng            float64
	TimeOffset     float64
	Prefix         string
	IsFullCallsign bool
	Flag           string
	CountryCode    string
}
