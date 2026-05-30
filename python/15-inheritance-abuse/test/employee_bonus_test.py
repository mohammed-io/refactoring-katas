from src.employee import Employee, Manager, SeniorManager, Director

def test_employee_gets_2_percent_bonus():
    emp = Employee("Alice", 50000)
    assert emp.calculate_bonus() == 1000

def test_manager_gets_5_percent_bonus():
    mgr = Manager("Bob", 80000)
    assert mgr.calculate_bonus() == 4000

def test_senior_manager_capped_at_10000():
    sm = SeniorManager("Carol", 300000)
    assert sm.calculate_bonus() == 10000

def test_senior_manager_under_cap():
    sm = SeniorManager("Carol", 100000)
    assert sm.calculate_bonus() == 5000

def test_director_capped_at_20000():
    dir = Director("Dave", 600000)
    assert dir.calculate_bonus() == 20000

def test_director_under_cap():
    dir = Director("Dave", 200000)
    assert dir.calculate_bonus() == 10000

def test_manager_respects_cap_from_senior_manager():
    mgr = Manager("Eve", 300000)
    assert mgr.calculate_bonus() == 15000

def test_employee_total_reward_adds_high_performance_and_tenure():
    emp = Employee("Alice", 50000)
    assert emp.calculate_total_reward("high", 5) == 2000

def test_senior_manager_total_reward_uses_capped_bonus_and_tenure_rule():
    sm = SeniorManager("Carol", 300000)
    assert sm.calculate_total_reward("normal", 7) == 11500

def test_director_total_reward_uses_director_performance_rule():
    director = Director("Dana", 200000)
    assert director.calculate_total_reward("high", 3) == 16000
