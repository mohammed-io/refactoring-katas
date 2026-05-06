# frozen_string_literal: true

class RoomLookup
  def get_room_for_student(student_id, db)
    return nil if db.nil?

    student = db.get_student(student_id)
    return nil if student.nil?

    enrollment = db.get_enrollment(student[:id])
    return nil if enrollment.nil?

    course = db.get_course(enrollment[:course_id])
    return nil if course.nil?

    section = db.get_section(course[:default_section_id])
    return nil if section.nil?

    room = db.get_room(section[:room_id])
    return room[:name] unless room.nil?

    nil
  end
end
