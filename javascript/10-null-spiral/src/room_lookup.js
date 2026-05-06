class RoomLookup {
  constructor() {}

  get_room_for_student(studentId, db) {
    if (db !== null) {
      let student = db.getStudent(studentId);
      if (student !== null) {
        let enrollment = db.getEnrollment(student.id);
        if (enrollment !== null) {
          let course = db.getCourse(enrollment.courseId);
          if (course !== null) {
            let section = db.getSection(course.defaultSectionId);
            if (section !== null) {
              let room = db.getRoom(section.roomId);
              if (room !== null) {
                return room.name;
              } else {
                return null;
              }
            } else {
              return null;
            }
          } else {
            return null;
          }
        } else {
          return null;
        }
      } else {
        return null;
      }
    } else {
      return null;
    }
  }
}

export { RoomLookup };
