public class Employee
{
    protected decimal Salary;

    public Employee(string name, decimal salary)
    {
        Salary = salary;
    }

    public virtual decimal calculate_bonus() => Salary * 0.02m;
}

public class Manager : Employee
{
    public Manager(string n, decimal s) : base(n, s)
    {
    }

    public override decimal calculate_bonus() => Salary * 0.05m;
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
}
