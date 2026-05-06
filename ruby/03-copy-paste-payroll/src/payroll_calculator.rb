# frozen_string_literal: true

class PayrollCalculator
  def generate_payslips(employees)
    payslips = []
    employees.each do |emp|
      gross = deductions = net = 0
      case emp[:type]
      when 'fulltime'
        gross = emp[:salary] / 12.0
        deductions = gross * 0.25
        gross += emp[:bonus] / 12.0 if emp[:bonus]
        net = gross - deductions
      when 'parttime'
        gross = emp[:hours] * emp[:rate]
        deductions = gross * 0.15
        net = gross - deductions
      when 'contract'
        gross = emp[:flat_fee]
        deductions = gross * 0.1
        net = gross - deductions
      end
      payslips << { id: emp[:id], name: emp[:name], type: emp[:type], gross: gross.round(2),
                    deductions: deductions.round(2), net: net.round(2) }
    end
    payslips
  end
end
