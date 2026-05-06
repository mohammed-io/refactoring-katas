package kata

type Student struct{ ID int }
type Enrollment struct{ CourseID int }
type Course struct{ DefaultSectionID int }
type Section struct{ RoomID int }
type Room struct{ Name string }
type DB interface {
	GetStudent(int) *Student
	GetEnrollment(int) *Enrollment
	GetCourse(int) *Course
	GetSection(int) *Section
	GetRoom(int) *Room
}

type RoomLookup struct{}

func NewRoomLookup() *RoomLookup {
	return &RoomLookup{}
}

func (rl *RoomLookup) get_room_for_student(studentID int, db DB) *string {
	if db != nil {
		student := db.GetStudent(studentID)
		if student != nil {
			enrollment := db.GetEnrollment(student.ID)
			if enrollment != nil {
				course := db.GetCourse(enrollment.CourseID)
				if course != nil {
					section := db.GetSection(course.DefaultSectionID)
					if section != nil {
						room := db.GetRoom(section.RoomID)
						if room != nil {
							return &room.Name
						} else {
							return nil
						}
					} else {
						return nil
					}
				} else {
					return nil
				}
			} else {
				return nil
			}
		} else {
			return nil
		}
	} else {
		return nil
	}
}
