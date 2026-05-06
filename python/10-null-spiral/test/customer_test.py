from src.room_lookup import RoomLookup

class Db:
    def __init__(self, missing=None):
        self.missing = missing

    def get_student(self, id):
        return None if self.missing == "student" else {"id": 1}

    def get_enrollment(self, id):
        return None if self.missing == "enrollment" else {"courseId": 2}

    def get_course(self, id):
        return None if self.missing == "course" else {"defaultSectionId": 3}

    def get_section(self, id):
        return None if self.missing == "section" else {"roomId": 4}

    def get_room(self, id):
        return None if self.missing == "room" else {"name": "Room A"}

def test_valid_chain():
    db = Db()
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) == "Room A"

def test_null_db():
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, None) is None

def test_null_student():
    db = Db(missing="student")
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) is None

def test_null_enrollment():
    db = Db(missing="enrollment")
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) is None

def test_null_course():
    db = Db(missing="course")
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) is None

def test_null_section():
    db = Db(missing="section")
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) is None

def test_null_room():
    db = Db(missing="room")
    lookup = RoomLookup()
    assert lookup.get_room_for_student(1, db) is None
