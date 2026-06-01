public class Package
{
    public int? Weight;
    public bool? Hazardous;
    public bool? Weekend;
    public int? TemperatureRequired;
    public bool RemoteArea;
}

public class DeliveryResult
{
    public bool Allowed;
    public string? Warning;
}

public class PackageValidator
{
    public DeliveryResult can_deliver(Package? pkg)
    {
        if (pkg == null)
        {
            return new DeliveryResult { Allowed = false, Warning = "No package" };
        }

        if (!pkg.Weight.HasValue)
        {
            return new DeliveryResult { Allowed = false, Warning = "No weight specified" };
        }

        if (pkg.Weight > 50)
        {
            return new DeliveryResult { Allowed = false, Warning = "Weight exceeded" };
        }

        if (pkg.Hazardous != false)
        {
            return new DeliveryResult { Allowed = false, Warning = "Hazardous material" };
        }

        if (pkg.TemperatureRequired.HasValue)
        {
            if (pkg.TemperatureRequired < -20 || pkg.TemperatureRequired > 40)
            {
                return new DeliveryResult { Allowed = false, Warning = "Temperature out of range" };
            }
        }

        if (pkg.Weekend != false)
        {
            return new DeliveryResult { Allowed = false, Warning = "No weekend delivery" };
        }

        if (pkg.RemoteArea)
        {
            if (pkg.Weight <= 20)
            {
                return new DeliveryResult { Allowed = true, Warning = "Remote surcharge applies" };
            }
            return new DeliveryResult { Allowed = false, Warning = "Too heavy for remote" };
        }

        return new DeliveryResult { Allowed = true, Warning = null };
    }
}
