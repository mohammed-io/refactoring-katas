import math
from src.calculator import Calculator

def test_normalizes_0_100():
    calc = Calculator()
    assert calc.normalize([10, 20, 30], 0, 100) == [0, 50, 100]

def test_normalizes_1_5():
    calc = Calculator()
    assert calc.normalize([10, 20, 30], 1, 5) == [1, 3, 5]

def test_single_value():
    calc = Calculator()
    result = calc.normalize([50], 0, 100)
    assert len(result) == 1
    assert math.isnan(result[0])

def test_negative_input_range():
    calc = Calculator()
    assert calc.normalize([-10, 0, 10], 0, 1) == [0, 0.5, 1]

def test_same_min_max():
    calc = Calculator()
    result = calc.normalize([5, 5, 5], 0, 100)
    assert math.isnan(result[0])
    assert math.isnan(result[1])
    assert math.isnan(result[2])

def test_reversed_output_range():
    calc = Calculator()
    assert calc.normalize([10, 20, 30], 100, 0) == [100, 50, 0]

def test_rounds_fractional_results():
    calc = Calculator()
    assert calc.normalize([2, 3, 5], 0, 100) == [0, 33.33, 100]

def test_empty_input_returns_empty_result():
    calc = Calculator()
    assert calc.normalize([], 0, 100) == []

def test_preserves_input_order_with_duplicates():
    calc = Calculator()
    assert calc.normalize([30, 10, 30, 20], 0, 100) == [100, 0, 100, 50]

def test_normalizes_decimal_output_range():
    calc = Calculator()
    assert calc.normalize([1.5, 2.5, 3.5], -1, 1) == [-1, 0, 1]

def test_rounds_negative_fractional_results():
    calc = Calculator()
    assert calc.normalize([2, 5, 8], -10, 10) == [-10, 0, 10]
