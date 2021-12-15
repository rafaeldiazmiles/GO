package errors

func ErrorPackage() {
	ErrorInternal := error.New("Internal error.")
	ErrorThirdParty := error.New("Third party error.")
	ErrorOther := error.New("Other error.")
}
