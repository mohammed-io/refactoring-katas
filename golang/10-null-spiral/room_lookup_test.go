package kata

import "testing"

type fakeDB struct{ missing string }

func (d fakeDB) GetStudent(int) *Student {
	if d.missing == "student" {
		return nil
	}
	return &Student{1}
}
func (d fakeDB) GetEnrollment(int) *Enrollment {
	if d.missing == "enrollment" {
		return nil
	}
	return &Enrollment{2}
}
func (d fakeDB) GetCourse(int) *Course {
	if d.missing == "course" {
		return nil
	}
	return &Course{3}
}
func (d fakeDB) GetSection(int) *Section {
	if d.missing == "section" {
		return nil
	}
	return &Section{4}
}
func (d fakeDB) GetRoom(int) *Room {
	if d.missing == "room" {
		return nil
	}
	return &Room{"Room A"}
}

func TestGetRoomForStudentReturnsRoomNameForValidChain(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{})
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if *result != "Room A" {
		t.Errorf("expected 'Room A', got %q", *result)
	}
}

func TestGetRoomForStudentReturnsNullForNullDb(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, nil)
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetRoomForStudentReturnsNullForNullStudent(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{missing: "student"})
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetRoomForStudentReturnsNullForNullEnrollment(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{missing: "enrollment"})
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetRoomForStudentReturnsNullForNullCourse(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{missing: "course"})
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetRoomForStudentReturnsNullForNullSection(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{missing: "section"})
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}

func TestGetRoomForStudentReturnsNullForNullRoom(t *testing.T) {
	rl := NewRoomLookup()
	result := rl.get_room_for_student(1, fakeDB{missing: "room"})
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
}
