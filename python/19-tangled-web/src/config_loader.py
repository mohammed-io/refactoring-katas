import json, os, datetime

class ConfigLoader:
    def __init__(self):
        pass

    def load_config(self):
        try:
            with open('/tmp/config.json') as f:
                local = json.load(f)
        except Exception:
            local = {}
        remote = {}
        seasonal = {}
        month = datetime.datetime.now().month - 1
        if month == 11 or month == 0 or month == 1:
            seasonal = {"theme":"winter","discount":0.1}
        elif month >= 5 and month <= 7:
            seasonal = {"theme":"summer","discount":0.05}
        result = {}
        result.update(remote); result.update(local); result.update(seasonal)
        return result
