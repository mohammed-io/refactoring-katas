public class Employee
{
    protected decimal Salary;

    public Employee(string name, decimal salary)
    {
        Salary = salary;
    }

    public virtual decimal calculate_bonus() => Salary * 0.02m;

    public virtual decimal calculate_total_reward(string performance, int years)
    {
        var x = calculate_bonus();
        if (performance == "high") x += Salary * 0.01m;
        if (years >= 5) x += 500;
        return x;
    }
}

public class Manager : Employee
{
    public Manager(string n, decimal s) : base(n, s)
    {
    }

    public override decimal calculate_bonus() => Salary * 0.05m;

    public override decimal calculate_total_reward(string performance, int years)
    {
        var x = calculate_bonus();
        if (performance == "high") x += Salary * 0.02m;
        if (years >= 5) x += 1000;
        return x;
    }
}

public class SeniorManager : Manager
{
    public SeniorManager(string n, decimal s) : base(n, s)
    {
    }

    public override decimal calculate_bonus()
    {
        var b = Salary * 0.05m;
        return b > 10000 ? 10000 : b;
    }

    public override decimal calculate_total_reward(string performance, int years)
    {
        var x = calculate_bonus();
        if (performance == "high") x += Salary * 0.02m;
        if (years >= 5) x += 1500;
        return x;
    }
}

public class Director : SeniorManager
{
    public Director(string n, decimal s) : base(n, s)
    {
    }

    public override decimal calculate_bonus()
    {
        var b = Salary * 0.05m;
        return b > 20000 ? 20000 : b;
    }

    public override decimal calculate_total_reward(string performance, int years)
    {
        var x = calculate_bonus();
        if (performance == "high") x += Salary * 0.03m;
        if (years >= 5) x += 2500;
        return x;
    }
}
