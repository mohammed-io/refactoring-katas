from src.loan_approver import LoanApprover

def test_normal_package():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 10, "hazardous": False, "weekend": False})
    assert result["allowed"] is True
    assert result["warning"] is None

def test_overweight():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 60, "hazardous": False, "weekend": False})
    assert result["allowed"] is False
    assert result["warning"] == "Weight exceeded"

def test_hazardous():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 10, "hazardous": True, "weekend": False})
    assert result["allowed"] is False
    assert result["warning"] == "Hazardous material"

def test_weekend():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 10, "hazardous": False, "weekend": True})
    assert result["allowed"] is False
    assert result["warning"] == "No weekend delivery"

def test_extreme_temperature():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": 50})
    assert result["allowed"] is False
    assert result["warning"] == "Temperature out of range"

def test_valid_temperature():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 10, "hazardous": False, "weekend": False, "temperatureRequired": 20})
    assert result["allowed"] is True

def test_remote_small():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 15, "hazardous": False, "weekend": False, "remoteArea": True})
    assert result["allowed"] is True
    assert result["warning"] == "Remote surcharge applies"

def test_remote_heavy():
    approver = LoanApprover()
    result = approver.can_deliver({"weight": 25, "hazardous": False, "weekend": False, "remoteArea": True})
    assert result["allowed"] is False
    assert result["warning"] == "Too heavy for remote"

def test_null_package():
    approver = LoanApprover()
    result = approver.can_deliver(None)
    assert result["allowed"] is False
    assert result["warning"] == "No package"

def test_missing_weight():
    approver = LoanApprover()
    result = approver.can_deliver({"hazardous": False})
    assert result["allowed"] is False
    assert result["warning"] == "No weight specified"
