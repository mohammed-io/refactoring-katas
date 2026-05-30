import os, json
from src.config_loader import ConfigLoader

def test_returns_defaults_when_file_is_missing(tmp_path, monkeypatch):
    monkeypatch.setenv("APP_CONFIG_PATH", str(tmp_path / "missing.json"))
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["retries"] == 3
    assert result["theme"] == "standard"
    assert result["discount"] == 0

def test_reads_local_config_file_when_present(tmp_path, monkeypatch):
    path = tmp_path / "config.json"
    path.write_text(json.dumps({"name": "test", "retries": 5}))
    monkeypatch.setenv("APP_CONFIG_PATH", str(path))
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["name"] == "test"
    assert result["retries"] == 5

def test_environment_overrides_local_config(tmp_path, monkeypatch):
    path = tmp_path / "config.json"
    path.write_text(json.dumps({"theme": "local", "retries": 5}))
    monkeypatch.setenv("APP_CONFIG_PATH", str(path))
    monkeypatch.setenv("APP_THEME", "env-theme")
    monkeypatch.setenv("APP_RETRIES", "9")
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["theme"] == "env-theme"
    assert result["retries"] == 9

def test_handles_malformed_json_gracefully(tmp_path, monkeypatch):
    path = tmp_path / "config.json"
    path.write_text("not json")
    monkeypatch.setenv("APP_CONFIG_PATH", str(path))
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["theme"] == "standard"

def test_winter_seasonal_config_has_highest_precedence(tmp_path, monkeypatch):
    path = tmp_path / "config.json"
    path.write_text(json.dumps({"theme": "local", "discount": 0.25}))
    monkeypatch.setenv("APP_CONFIG_PATH", str(path))
    monkeypatch.setenv("APP_THEME", "env-theme")
    monkeypatch.setenv("APP_CURRENT_MONTH", "12")
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["theme"] == "winter"
    assert result["discount"] == 0.1

def test_summer_seasonal_config_is_deterministic(tmp_path, monkeypatch):
    monkeypatch.setenv("APP_CONFIG_PATH", str(tmp_path / "missing.json"))
    monkeypatch.setenv("APP_CURRENT_MONTH", "7")
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["theme"] == "summer"
    assert result["discount"] == 0.05
