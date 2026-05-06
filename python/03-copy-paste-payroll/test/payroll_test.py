from src.payroll_calculator import PayrollCalculator

def test_fulltime_payslip():
    calc = PayrollCalculator()
    result = calc.generate_payslips([{"id": 1, "name": "Alice", "type": "fulltime", "salary": 60000}])
    assert result[0]["gross"] == 5000
    assert result[0]["deductions"] == 1250
    assert result[0]["net"] == 3750

def test_fulltime_with_bonus():
    calc = PayrollCalculator()
    result = calc.generate_payslips([{"id": 1, "name": "Alice", "type": "fulltime", "salary": 60000, "bonus": 12000}])
    assert result[0]["gross"] == 6000
    assert result[0]["net"] == 4750

def test_parttime_payslip():
    calc = PayrollCalculator()
    result = calc.generate_payslips([{"id": 2, "name": "Bob", "type": "parttime", "hours": 80, "rate": 25}])
    assert result[0]["gross"] == 2000
    assert result[0]["deductions"] == 300
    assert result[0]["net"] == 1700

def test_contract_payslip():
    calc = PayrollCalculator()
    result = calc.generate_payslips([{"id": 3, "name": "Carol", "type": "contract", "flatFee": 5000}])
    assert result[0]["gross"] == 5000
    assert result[0]["deductions"] == 500
    assert result[0]["net"] == 4500

def test_multiple_employees():
    calc = PayrollCalculator()
    result = calc.generate_payslips([
        {"id": 1, "name": "Alice", "type": "fulltime", "salary": 60000},
        {"id": 2, "name": "Bob", "type": "parttime", "hours": 80, "rate": 25}
    ])
    assert len(result) == 2
    assert result[0]["name"] == "Alice"
    assert result[1]["name"] == "Bob"

def test_empty_input():
    calc = PayrollCalculator()
    result = calc.generate_payslips([])
    assert result == []
