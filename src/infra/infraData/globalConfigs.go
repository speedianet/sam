package infraData

type globalConfigs struct {
	PkiConfDir                       string
	DomainOwnershipValidationUrlPath string
}

var GlobalConfigs = globalConfigs{
	PkiConfDir:                       "/app/conf/pki",
	DomainOwnershipValidationUrlPath: "/validateOwnership",
}