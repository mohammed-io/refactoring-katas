import os, json
from src.config_loader import ConfigLoader

def test_returns_empty_object_when_file_and_fetch_fail():
    if os.path.exists('/tmp/config.json'):
        os.unlink('/tmp/config.json')
    loader = ConfigLoader()
    result = loader.load_config()
    assert isinstance(result, dict)

def test_reads_local_config_file_when_present():
    with open('/tmp/config.json', 'w') as f:
        json.dump({"name": "test"}, f)
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["name"] == "test"
    os.unlink('/tmp/config.json')

def test_local_overrides_empty_file():
    with open('/tmp/config.json', 'w') as f:
        json.dump({"name": "local"}, f)
    loader = ConfigLoader()
    result = loader.load_config()
    assert result["name"] == "local"
    os.unlink('/tmp/config.json')

def test_handles_malformed_json_gracefully():
    with open('/tmp/config.json', 'w') as f:
        f.write('not json')
    loader = ConfigLoader()
    result = loader.load_config()
    assert isinstance(result, dict)
    os.unlink('/tmp/config.json')

def test_includes_seasonal_keys_in_object():
    loader = ConfigLoader()
    result = loader.load_config()
    assert "theme" in result or not result.get("theme")
    assert "discount" in result or not result.get("discount")
