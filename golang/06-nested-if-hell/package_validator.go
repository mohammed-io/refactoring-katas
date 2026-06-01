package kata

type Package struct {
	Weight              int
	HasWeight           bool
	Hazardous           bool
	Weekend             bool
	TemperatureRequired *int
	RemoteArea          bool
}
type DeliveryResult struct {
	Allowed bool
	Warning string
}

type PackageValidator struct{}

func NewPackageValidator() *PackageValidator {
	return &PackageValidator{}
}

func (pv *PackageValidator) can_deliver(pkg *Package) DeliveryResult {
	if pkg != nil {
		if pkg.HasWeight {
			if pkg.Weight <= 50 {
				if pkg.Hazardous == false {
					if pkg.TemperatureRequired != nil {
						if *pkg.TemperatureRequired >= -20 && *pkg.TemperatureRequired <= 40 {
							return pv.weekend_remote(pkg)
						} else {
							return DeliveryResult{false, "Temperature out of range"}
						}
					} else {
						return pv.weekend_remote(pkg)
					}
				} else {
					return DeliveryResult{false, "Hazardous material"}
				}
			} else {
				return DeliveryResult{false, "Weight exceeded"}
			}
		} else {
			return DeliveryResult{false, "No weight specified"}
		}
	} else {
		return DeliveryResult{false, "No package"}
	}
}
func (pv *PackageValidator) weekend_remote(pkg *Package) DeliveryResult {
	if pkg.Weekend == false {
		if pkg.RemoteArea {
			if pkg.Weight <= 20 {
				return DeliveryResult{true, "Remote surcharge applies"}
			}
			return DeliveryResult{false, "Too heavy for remote"}
		}
		return DeliveryResult{true, ""}
	}
	return DeliveryResult{false, "No weekend delivery"}
}
