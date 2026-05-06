class PayrollCalculator {
  constructor() {}

  generate_payslips(employees) {
    let payslips = [];

    for (let i = 0; i < employees.length; i++) {
      let emp = employees[i];
      let gross = 0;
      let deductions = 0;
      let net = 0;

      if (emp.type === 'fulltime') {
        gross = emp.salary / 12;
        deductions = gross * 0.25;
        if (emp.bonus) {
          gross += emp.bonus / 12;
        }
        net = gross - deductions;
      } else if (emp.type === 'parttime') {
        gross = emp.hours * emp.rate;
        deductions = gross * 0.15;
        net = gross - deductions;
      } else if (emp.type === 'contract') {
        gross = emp.flatFee;
        deductions = gross * 0.1;
        net = gross - deductions;
      }

      payslips.push({
        id: emp.id,
        name: emp.name,
        type: emp.type,
        gross: Math.round(gross * 100) / 100,
        deductions: Math.round(deductions * 100) / 100,
        net: Math.round(net * 100) / 100
      });
    }

    return payslips;
  }
}

export { PayrollCalculator };
