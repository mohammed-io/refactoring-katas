public class Employee
{
    public int Id;
    public string Name = "";
    public string Type = "";
    public decimal Salary;
    public decimal Bonus;
    public decimal Hours;
    public decimal Rate;
    public decimal FlatFee;
}

public class Payslip
{
    public int Id;
    public string Name = "";
    public string Type = "";
    public decimal Gross;
    public decimal Deductions;
    public decimal Net;
}

public class PayrollCalculator
{
    public List<Payslip> generate_payslips(List<Employee> employees)
    {
        var slips = new List<Payslip>();
        foreach (var emp in employees)
        {
            decimal gross = 0, deductions = 0, net = 0;
            if (emp.Type == "fulltime")
            {
                var baseGross = emp.Salary / 12;
                gross = baseGross + (emp.Bonus != 0 ? emp.Bonus / 12 : 0);
                deductions = baseGross * 0.25m;
                net = gross - deductions;
            }
            else if (emp.Type == "parttime")
            {
                gross = emp.Hours * emp.Rate;
                deductions = gross * 0.15m;
                net = gross - deductions;
            }
            else if (emp.Type == "contract")
            {
                gross = emp.FlatFee;
                deductions = gross * 0.1m;
                net = gross - deductions;
            }

            slips.Add(new Payslip
            {
                Id = emp.Id,
                Name = emp.Name,
                Type = emp.Type,
                Gross = Math.Round(gross, 2),
                Deductions = Math.Round(deductions, 2),
                Net = Math.Round(net, 2)
            });
        }
        return slips;
    }
}
