using Xunit;

public class VehicleTest
{
    [Fact]
    public void CarDailyRate()
    {
        var v = new Car("Toyota", "Camry", 2020);
        Assert.Equal(40, v.daily_rate());
    }

    [Fact]
    public void TruckDailyRate()
    {
        var v = new Truck("Ford", "F-150", 2020);
        Assert.Equal(80, v.daily_rate());
    }

    [Fact]
    public void ElectricCarFuelCost()
    {
        var v = new ElectricCar("Tesla", "Model 3", 2020);
        Assert.Equal(0, v.fuel_cost(5));
    }

    [Fact]
    public void DieselCarFuelCost()
    {
        var v = new DieselCar("VW", "Jetta", 2020);
        Assert.Equal(25, v.fuel_cost(5));
    }

    [Fact]
    public void ElectricTruckFuelCost()
    {
        var v = new ElectricTruck("Rivian", "R1T", 2020);
        Assert.Equal(0, v.fuel_cost(5));
    }

    [Fact]
    public void DieselTruckFuelCost()
    {
        var v = new DieselTruck("Ford", "F-250", 2020);
        Assert.Equal(75, v.fuel_cost(5));
    }

    [Fact]
    public void CarStoresBrand()
    {
        var v = new Car("Toyota", "Camry", 2020);
        Assert.Equal("Toyota", v.Brand);
    }

    [Fact]
    public void TruckStoresModel()
    {
        var v = new Truck("Ford", "F-150", 2020);
        Assert.Equal("F-150", v.Model);
    }

    [Fact]
    public void DieselCarRentalTotalCombinesRateFuelInsuranceAndGps()
    {
        var v = new DieselCar("VW", "Jetta", 2020);
        Assert.Equal(195, v.rental_total(3, true));
    }

    [Fact]
    public void ElectricTruckRentalTotalCombinesTruckRateAndInsurance()
    {
        var v = new ElectricTruck("Rivian", "R1T", 2020);
        Assert.Equal(200, v.rental_total(2, false));
    }

    [Fact]
    public void DieselTruckRentalTotalIncludesHigherFuelCost()
    {
        var v = new DieselTruck("Ford", "F-250", 2020);
        Assert.Equal(246, v.rental_total(2, true));
    }
}
