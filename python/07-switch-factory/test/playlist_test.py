from src.playlist_curator import PlaylistCurator


def test_filters_happy_tracks():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 130}, {"title": "B", "tempo": 100}]
    result = curator.create_playlist("happy", tracks)
    assert len(result) == 1
    assert result[0]["title"] == "A"


def test_filters_sad_tracks():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 80}, {"title": "B", "tempo": 100}]
    result = curator.create_playlist("sad", tracks)
    assert len(result) == 1
    assert result[0]["title"] == "A"


def test_filters_workout_tracks():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 140, "energy": 8}, {"title": "B", "tempo": 140, "energy": 5}]
    result = curator.create_playlist("workout", tracks)
    assert len(result) == 1
    assert result[0]["title"] == "A"


def test_filters_focus_tracks():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "instrumental": True}, {"title": "B", "instrumental": False}]
    result = curator.create_playlist("focus", tracks)
    assert len(result) == 1
    assert result[0]["title"] == "A"


def test_filters_party_tracks():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 120, "danceability": 7}, {"title": "B", "tempo": 100, "danceability": 5}]
    result = curator.create_playlist("party", tracks)
    assert len(result) == 1
    assert result[0]["title"] == "A"


def test_caps_happy_playlist_at_20():
    curator = PlaylistCurator()
    tracks = [{"title": str(i), "tempo": 130} for i in range(25)]
    assert len(curator.create_playlist("happy", tracks)) == 20


def test_caps_sad_playlist_at_15():
    curator = PlaylistCurator()
    tracks = [{"title": str(i), "tempo": 80} for i in range(20)]
    assert len(curator.create_playlist("sad", tracks)) == 15


def test_defaults_to_first_10_for_unknown_mood():
    curator = PlaylistCurator()
    tracks = [{"title": str(i), "tempo": 100} for i in range(20)]
    assert len(curator.create_playlist("mysterious", tracks)) == 10


def test_sorts_happy_by_descending_tempo():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 130}, {"title": "B", "tempo": 150}]
    result = curator.create_playlist("happy", tracks)
    assert [track["title"] for track in result] == ["B", "A"]


def test_sorts_sad_by_ascending_tempo():
    curator = PlaylistCurator()
    tracks = [{"title": "A", "tempo": 80}, {"title": "B", "tempo": 60}]
    result = curator.create_playlist("sad", tracks)
    assert [track["title"] for track in result] == ["B", "A"]
