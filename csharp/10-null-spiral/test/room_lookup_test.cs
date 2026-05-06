using Xunit;

class FakeDb : IRoomDb
{
    private string? _missing;

    public FakeDb(string? missing = null)
    {
        _missing = missing;
    }

    public Student? GetStudent(int id) => _missing == "student" ? null : new Student(1);
    public Enrollment? GetEnrollment(int id) => _missing == "enrollment" ? null : new Enrollment(2);
    public Course? GetCourse(int id) => _missing == "course" ? null : new Course(3);
    public Section? GetSection(int id) => _missing == "section" ? null : new Section(4);
    public Room? GetRoom(int id) => _missing == "room" ? null : new Room("Room A");
}

public class RoomLookupTest
{
    [Fact]
    public void ReturnsRoomNameForValidChain()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb());
        Assert.Equal("Room A", result);
    }

    [Fact]
    public void ReturnsNullForNullDb()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, null);
        Assert.Null(result);
    }

    [Fact]
    public void ReturnsNullForNullStudent()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb("student"));
        Assert.Null(result);
    }

    [Fact]
    public void ReturnsNullForNullEnrollment()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb("enrollment"));
        Assert.Null(result);
    }

    [Fact]
    public void ReturnsNullForNullCourse()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb("course"));
        Assert.Null(result);
    }

    [Fact]
    public void ReturnsNullForNullSection()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb("section"));
        Assert.Null(result);
    }

    [Fact]
    public void ReturnsNullForNullRoom()
    {
        var lookup = new RoomLookup();
        var result = lookup.get_room_for_student(1, new FakeDb("room"));
        Assert.Null(result);
    }
}
