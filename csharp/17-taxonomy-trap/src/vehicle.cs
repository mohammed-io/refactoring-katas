public class Vehicle
{
    public string Brand;
    public string Model;
    public int Year;

    public Vehicle(string b, string m, int y)
    {
        Brand = b;
        Model = m;
        Year = y;
    }
}

public class Car : Vehicle
{
    public Car(string b, string m, int y) : base(b, m, y)
    {
    }

    public int daily_rate() => 40;

    public int insurance_cost(int d) => d * 12;
}

public class Truck : Vehicle
{
    public Truck(string b, string m, int y) : base(b, m, y)
    {
    }

    public int daily_rate() => 80;

    public int insurance_cost(int d) => d * 20;
}

public class ElectricCar : Car
{
    public ElectricCar(string b, string m, int y) : base(b, m, y)
    {
    }

    public int fuel_cost(int d) => 0;

    public int rental_total(int d, bool gps = false) => daily_rate() * d + fuel_cost(d) + insurance_cost(d) + (gps ? 8 * d : 0);
}

public class DieselCar : Car
{
    public DieselCar(string b, string m, int y) : base(b, m, y)
    {
    }

    public int fuel_cost(int d) => d * 5;

    public int rental_total(int d, bool gps = false) => daily_rate() * d + fuel_cost(d) + insurance_cost(d) + (gps ? 8 * d : 0);
}

public class ElectricTruck : Truck
{
    public ElectricTruck(string b, string m, int y) : base(b, m, y)
    {
    }

    public int fuel_cost(int d) => 0;

    public int rental_total(int d, bool gps = false) => daily_rate() * d + fuel_cost(d) + insurance_cost(d) + (gps ? 8 * d : 0);
}

public class DieselTruck : Truck
{
    public DieselTruck(string b, string m, int y) : base(b, m, y)
    {
    }

    public int fuel_cost(int d) => d * 15;

    public int rental_total(int d, bool gps = false) => daily_rate() * d + fuel_cost(d) + insurance_cost(d) + (gps ? 8 * d : 0);
}
