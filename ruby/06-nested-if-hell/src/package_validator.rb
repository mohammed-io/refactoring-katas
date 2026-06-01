# frozen_string_literal: true

class PackageValidator
  def can_deliver(pkg)
    return { allowed: false, warning: 'No package' } unless pkg
    return { allowed: false, warning: 'No weight specified' } unless pkg[:weight]
    return { allowed: false, warning: 'Weight exceeded' } unless pkg[:weight] <= 50
    return { allowed: false, warning: 'Hazardous material' } unless pkg[:hazardous] == false
    return weekend_and_remote(pkg) unless pkg[:temperature_required]
    return weekend_and_remote(pkg) if pkg[:temperature_required] >= -20 && pkg[:temperature_required] <= 40

    { allowed: false, warning: 'Temperature out of range' }
  end

  def weekend_and_remote(pkg)
    if pkg[:weekend] == false
      if pkg[:remote_area]
        return pkg[:weight] <= 20 ? { allowed: true,
                                      warning: 'Remote surcharge applies' } : { allowed: false,
                                                                                warning: 'Too heavy for remote' }
      end

      { allowed: true, warning: nil }
    else
      { allowed: false, warning: 'No weekend delivery' }
    end
  end
end
