import time
from src.fraud_detector import FraudDetector

def test_low_risk_small_transaction():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    result = detector.detect({"amount": 10, "timestamp": now, "history": [], "merchant": "grocery", "country": "US", "cardCountry": "US"})
    assert result["rating"] == "low"
    assert result["level"] == 1

def test_medium_risk_for_large_amount():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    result = detector.detect({"amount": 1100, "timestamp": now, "history": [], "merchant": "grocery", "country": "US", "cardCountry": "US"})
    assert result["rating"] == "medium"
    assert result["level"] == 2

def test_elevated_risk_for_gambling():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    result = detector.detect({"amount": 100, "timestamp": now, "history": [], "merchant": "gambling", "country": "US", "cardCountry": "US"})
    assert result["rating"] == "medium"
    assert result["level"] == 2

def test_high_risk_for_cross_border():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    result = detector.detect({"amount": 1000, "timestamp": now, "history": [], "merchant": "grocery", "country": "FR", "cardCountry": "US"})
    assert result["rating"] == "low"
    assert result["level"] == 1

def test_critical_risk_for_late_night_crypto():
    detector = FraudDetector()
    ts = int(time.mktime(time.strptime("2024-01-01T02:00:00", "%Y-%m-%dT%H:%M:%S")) * 1000)
    result = detector.detect({"amount": 600, "timestamp": ts, "history": [], "merchant": "crypto", "country": "CN", "cardCountry": "US"})
    assert result["rating"] == "high"
    assert result["level"] == 4

def test_velocity_increases_risk():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    hist = [{"amount": 10, "timestamp": now - i * 1000} for i in range(1, 5)]
    result = detector.detect({"amount": 50, "timestamp": now, "history": hist, "merchant": "grocery", "country": "US", "cardCountry": "US"})
    assert result["level"] == 2

def test_volume_spikes_increase_risk():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    hist = [{"amount": 200, "timestamp": now - i * 10000} for i in range(1, 4)]
    result = detector.detect({"amount": 50, "timestamp": now, "history": hist, "merchant": "grocery", "country": "US", "cardCountry": "US"})
    assert result["level"] == 1

def test_includes_score():
    detector = FraudDetector()
    now = int(time.time() * 1000)
    result = detector.detect({"amount": 10, "timestamp": now, "history": [], "merchant": "grocery", "country": "US", "cardCountry": "US"})
    assert isinstance(result["score"], int)
