from datetime import datetime

class FraudDetector:
    def __init__(self):
        pass

    def detect(self, tx):
        s = 0; v = 0; m = 0
        hour = datetime.fromtimestamp(tx["timestamp"] / 1000).hour
        for h in tx["history"]:
            if h["amount"] > tx["amount"] * 2: v += 1
            if abs(tx["timestamp"] - h["timestamp"]) < 3600000: s += 1
        if tx["amount"] > 500 and hour >= 0 and hour < 6: m += 30
        if tx["amount"] > 1000: m += 20
        if tx["merchant"] == "gambling" or tx["merchant"] == "crypto": m += 25
        if tx["country"] != tx["cardCountry"]: m += 15
        if s > 3: m += 20
        if v > 2: m += 15
        if m < 20: level = 1
        elif m < 40: level = 2
        elif m < 60: level = 3
        elif m < 80: level = 4
        else: level = 5
        rating = {1:"low",2:"medium",3:"elevated",4:"high",5:"critical"}[level]
        return {"score":m,"level":level,"rating":rating}
