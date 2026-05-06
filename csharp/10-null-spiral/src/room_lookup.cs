public record Student(int Id);
public record Enrollment(int CourseId);
public record Course(int DefaultSectionId);
public record Section(int RoomId);
public record Room(string Name);

public interface IRoomDb
{
    Student? GetStudent(int id);
    Enrollment? GetEnrollment(int id);
    Course? GetCourse(int id);
    Section? GetSection(int id);
    Room? GetRoom(int id);
}

public class RoomLookup
{
    public string? get_room_for_student(int id, IRoomDb? db)
    {
        if (db != null)
        {
            var student = db.GetStudent(id);
            if (student != null)
            {
                var enrollment = db.GetEnrollment(student.Id);
                if (enrollment != null)
                {
                    var course = db.GetCourse(enrollment.CourseId);
                    if (course != null)
                    {
                        var section = db.GetSection(course.DefaultSectionId);
                        if (section != null)
                        {
                            var room = db.GetRoom(section.RoomId);
                            if (room != null)
                            {
                                return room.Name;
                            }
                            else
                            {
                                return null;
                            }
                        }
                        else
                        {
                            return null;
                        }
                    }
                    else
                    {
                        return null;
                    }
                }
                else
                {
                    return null;
                }
            }
            else
            {
                return null;
            }
        }
        else
        {
            return null;
        }
    }
}
