package titleservice

// Codes used in the MMS TitleService API
const (
	Adults   = "V"
	Children = "B"
	Swedish  = "S"
	Foreign  = "U"
)

// Endpoint type
type Endpoint string

// Endpoints in the MMS TitleService API
const (
	RegisterSeriesEndpoint  Endpoint = "RegisterSeries"
	RegisterEpisodeEndpoint Endpoint = "RegisterEpisode"
	RegisterClipEndpoint    Endpoint = "RegisterClip"
)

// CategoryID type
type CategoryID int

// CategoryIDs
const (
	TvProgram        CategoryID = 1
	TvSegment        CategoryID = 2
	TvExtra          CategoryID = 3
	Webisode         CategoryID = 4
	WebSegment       CategoryID = 5
	WebExtra         CategoryID = 6
	WebClip          CategoryID = 7
	Simulcast        CategoryID = 8
	ChannelSimulcast CategoryID = 9
	WebLiveBroadcast CategoryID = 10
)

func validCategoryID(id CategoryID) bool {
	switch id {
	case TvProgram, TvSegment, TvExtra,
		Webisode, WebSegment, WebExtra, WebClip,
		Simulcast, ChannelSimulcast, WebLiveBroadcast:
		return true
	}

	return false
}

// LiveChannelID type used to identify Live TV Broadcast Channels
type LiveChannelID int

// LiveChannelIDs
const (
	SVT1                   LiveChannelID = 1001
	SVT2                   LiveChannelID = 1002
	Discovery              LiveChannelID = 1005
	Kanal5                 LiveChannelID = 1015
	TV3                    LiveChannelID = 1028
	TV4                    LiveChannelID = 1029
	Kanal9                 LiveChannelID = 1043
	SVTB                   LiveChannelID = 1050
	Sjuan                  LiveChannelID = 1061
	TV4Film                LiveChannelID = 1070
	TV6                    LiveChannelID = 1074
	TV4Sport               LiveChannelID = 1089
	TV4Fakta               LiveChannelID = 1103
	TV4Guld                LiveChannelID = 1153
	TV4Komedi              LiveChannelID = 1155
	TV8                    LiveChannelID = 1197
	SVT24                  LiveChannelID = 1211
	TV12                   LiveChannelID = 1508
	Kunskapskanalen        LiveChannelID = 1671
	TV10                   LiveChannelID = 1683
	TLC                    LiveChannelID = 1733
	InvestigationDiscovery LiveChannelID = 1736
	Eurosport              LiveChannelID = 1906
	Kanal11                LiveChannelID = 2025
	Eurosport2Sweden       LiveChannelID = 2047
)

// LookupLiveChannelID using the channel name as specified in the MMS TitleService v1.2 documentation
func LookupLiveChannelID(name string) (LiveChannelID, bool) {
	id, ok := channelLookupTable[name]

	return id, ok
}

func validLiveChannelID(id LiveChannelID) bool {
	switch id {
	case
		SVT1, SVT2, Discovery, Kanal5, TV3, TV4, Kanal9, SVTB, Sjuan, TV4Film, TV6,
		TV4Sport, TV4Fakta, TV4Guld, TV4Komedi, TV8, SVT24, TV12, Kunskapskanalen,
		TV10, TLC, InvestigationDiscovery, Eurosport, Kanal11, Eurosport2Sweden:
		return true
	}

	return false
}

var channelLookupTable = map[string]LiveChannelID{
	"svt1":            SVT1,
	"svt2":            SVT2,
	"Discovery":       Discovery,
	"Kanal5":          Kanal5,
	"TV3":             TV3,
	"TV4":             TV4,
	"Kanal9":          Kanal9,
	"svtB":            SVTB,
	"Sjuan":           Sjuan,
	"TV4 Film":        TV4Film,
	"TV6":             TV6,
	"TV4 Sport":       TV4Sport,
	"TV4 Fakta":       TV4Fakta,
	"TV4 Guld":        TV4Guld,
	"TV4 Komedi":      TV4Komedi,
	"TV8":             TV8,
	"svt24":           SVT24,
	"TV12":            TV12,
	"Kunskapskanalen": Kunskapskanalen,
	"TV10":            TV10,
	"TLC":             TLC,
	"Investigation Discovery": InvestigationDiscovery,
	"Eurosport":               Eurosport,
	"Kanal 11":                Kanal11,
	"Eurosport 2 Sweden":      Eurosport2Sweden,
}
