class LoanApprover:
    def __init__(self):
        pass

    def can_deliver(self, pkg):
        if pkg:
            if pkg.get("weight"):
                if pkg["weight"] <= 50:
                    if pkg.get("hazardous") is False:
                        if pkg.get("temperatureRequired"):
                            if pkg["temperatureRequired"] >= -20 and pkg["temperatureRequired"] <= 40:
                                if pkg.get("weekend") is False:
                                    if pkg.get("remoteArea"):
                                        if pkg["weight"] <= 20: return {"allowed": True, "warning": "Remote surcharge applies"}
                                        else: return {"allowed": False, "warning": "Too heavy for remote"}
                                    return {"allowed": True, "warning": None}
                                return {"allowed": False, "warning": "No weekend delivery"}
                            return {"allowed": False, "warning": "Temperature out of range"}
                        else:
                            if pkg.get("weekend") is False:
                                if pkg.get("remoteArea"):
                                    if pkg["weight"] <= 20: return {"allowed": True, "warning": "Remote surcharge applies"}
                                    else: return {"allowed": False, "warning": "Too heavy for remote"}
                                return {"allowed": True, "warning": None}
                            return {"allowed": False, "warning": "No weekend delivery"}
                    return {"allowed": False, "warning": "Hazardous material"}
                return {"allowed": False, "warning": "Weight exceeded"}
            return {"allowed": False, "warning": "No weight specified"}
        return {"allowed": False, "warning": "No package"}
