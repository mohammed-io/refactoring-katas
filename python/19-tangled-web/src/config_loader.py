import json, os, datetime

class ConfigLoader:
    def __init__(self):
        pass

    def load_config(self):
        defaults = {"retries": 3, "theme": "standard", "discount": 0}
        try:
            path = os.environ.get("APP_CONFIG_PATH", "/tmp/config.json")
            with open(path) as f:
                local = json.load(f)
        except Exception:
            local = {}

        env_config = {}
        if os.environ.get("APP_THEME"):
            env_config["theme"] = os.environ["APP_THEME"]
        if os.environ.get("APP_RETRIES"):
            env_config["retries"] = int(os.environ["APP_RETRIES"])

        seasonal = {}
        month = int(os.environ.get("APP_CURRENT_MONTH", datetime.datetime.now().month))
        if month == 12 or month == 1 or month == 2:
            seasonal = {"theme":"winter","discount":0.1}
        elif month >= 6 and month <= 8:
            seasonal = {"theme":"summer","discount":0.05}
        result = {}
        result.update(defaults); result.update(local); result.update(env_config); result.update(seasonal)
        return result
