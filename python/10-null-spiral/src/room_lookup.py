class RoomLookup:
    def __init__(self):
        pass

    def get_room_for_student(self, student_id, db):
        if db is not None:
            student = db.get_student(student_id)
            if student is not None:
                enrollment = db.get_enrollment(student["id"])
                if enrollment is not None:
                    course = db.get_course(enrollment["courseId"])
                    if course is not None:
                        section = db.get_section(course["defaultSectionId"])
                        if section is not None:
                            room = db.get_room(section["roomId"])
                            if room is not None:
                                return room["name"]
                            else: return None
                        else: return None
                    else: return None
                else: return None
            else: return None
        else: return None
