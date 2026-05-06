class PayrollCalculator:
    def __init__(self):
        pass

    def generate_payslips(self, employees):
        payslips = []
        for emp in employees:
            gross = deductions = net = 0
            if emp["type"] == "fulltime":
                gross = emp["salary"] / 12
                deductions = gross * 0.25
                if emp.get("bonus"):
                    gross += emp["bonus"] / 12
                net = gross - deductions
            elif emp["type"] == "parttime":
                gross = emp["hours"] * emp["rate"]
                deductions = gross * 0.15
                net = gross - deductions
            elif emp["type"] == "contract":
                gross = emp["flatFee"]
                deductions = gross * 0.1
                net = gross - deductions
            payslips.append({"id": emp["id"], "name": emp["name"], "type": emp["type"], "gross": round(gross, 2), "deductions": round(deductions, 2), "net": round(net, 2)})
        return payslips
