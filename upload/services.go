package upload

type OnlineService string

const (
	OnlineServiceSM   OnlineService = "sm"
	OnlineServiceQRZ  OnlineService = "qrz"
	OnlineServiceLoTW OnlineService = "lotw"
	OnlineServiceEQSL OnlineService = "eqsl"
	OnlineServiceClub OnlineService = "clublog"
)

var AllowedUploadServices = map[OnlineService]struct{}{
	OnlineServiceSM:   {},
	OnlineServiceQRZ:  {},
	OnlineServiceLoTW: {},
	OnlineServiceEQSL: {},
	OnlineServiceClub: {},
}

func (s OnlineService) Valid() bool {
	_, ok := AllowedUploadServices[s]
	return ok
}
