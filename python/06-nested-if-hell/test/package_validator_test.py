from src.package_validator import PackageValidator

def test_normal_package():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": False})
    assert result["allowed"] is True
    assert result["warning"] is None

def test_overweight():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 60, "hazardous": False, "weekend": False})
    assert result["allowed"] is False
    assert result["warning"] == "Weight exceeded"

def test_hazardous():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": True, "weekend": False})
    assert result["allowed"] is False
    assert result["warning"] == "Hazardous material"

def test_weekend():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": True})
    assert result["allowed"] is False
    assert result["warning"] == "No weekend delivery"

def test_extreme_temperature():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": 50})
    assert result["allowed"] is False
    assert result["warning"] == "Temperature out of range"

def test_valid_temperature():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": 20})
    assert result["allowed"] is True

def test_remote_small():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 15, "hazardous": False, "weekend": False, "remoteArea": True})
    assert result["allowed"] is True
    assert result["warning"] == "Remote surcharge applies"

def test_remote_heavy():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 25, "hazardous": False, "weekend": False, "remoteArea": True})
    assert result["allowed"] is False
    assert result["warning"] == "Too heavy for remote"

def test_null_package():
    validator = PackageValidator()
    result = validator.can_deliver(None)
    assert result["allowed"] is False
    assert result["warning"] == "No package"

def test_missing_weight():
    validator = PackageValidator()
    result = validator.can_deliver({"hazardous": False})
    assert result["allowed"] is False
    assert result["warning"] == "No weight specified"


def test_weight_50_allowed():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 50, "hazardous": False, "weekend": False})
    assert result["allowed"] is True
    assert result["warning"] is None


def test_temperature_40_allowed():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": 40})
    assert result["allowed"] is True


def test_temperature_minus_20_allowed():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": -20})
    assert result["allowed"] is True


def test_remote_weight_20_allowed():
    validator = PackageValidator()
    result = validator.can_deliver({"weight": 20, "hazardous": False, "weekend": False, "remoteArea": True})
    assert result["allowed"] is True
    assert result["warning"] == "Remote surcharge applies"
