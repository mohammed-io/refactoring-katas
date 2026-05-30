package kata

import "testing"

func TestCreatePlaylistHappyFilter(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 130, 0, false, 0}, {"B", 100, 0, false, 0}}
	result := pc.create_playlist("happy", tracks)
	if len(result) != 1 || result[0].Title != "A" {
		t.Fatal("happy filter")
	}
}

func TestCreatePlaylistSadFilter(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 80, 0, false, 0}, {"B", 100, 0, false, 0}}
	result := pc.create_playlist("sad", tracks)
	if len(result) != 1 || result[0].Title != "A" {
		t.Fatal("sad filter")
	}
}

func TestCreatePlaylistWorkoutFilter(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 140, 8, false, 0}, {"B", 140, 5, false, 0}}
	result := pc.create_playlist("workout", tracks)
	if len(result) != 1 || result[0].Title != "A" {
		t.Fatal("workout filter")
	}
}

func TestCreatePlaylistFocusFilter(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 0, 0, true, 0}, {"B", 0, 0, false, 0}}
	result := pc.create_playlist("focus", tracks)
	if len(result) != 1 || result[0].Title != "A" {
		t.Fatal("focus filter")
	}
}

func TestCreatePlaylistPartyFilter(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 120, 0, false, 7}, {"B", 100, 0, false, 5}}
	result := pc.create_playlist("party", tracks)
	if len(result) != 1 || result[0].Title != "A" {
		t.Fatal("party filter")
	}
}

func TestCreatePlaylistCapsHappyAt20(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 25)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i)), Tempo: 130, Energy: 0, Instrumental: false, Danceability: 0}
	}
	result := pc.create_playlist("happy", tracks)
	if len(result) != 20 {
		t.Fatalf("happy cap: expected 20, got %d", len(result))
	}
}

func TestCreatePlaylistCapsSadAt15(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 20)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i)), Tempo: 80, Energy: 0, Instrumental: false, Danceability: 0}
	}
	result := pc.create_playlist("sad", tracks)
	if len(result) != 15 {
		t.Fatalf("sad cap: expected 15, got %d", len(result))
	}
}

func TestCreatePlaylistDefaultsTo10(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 20)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i)), Tempo: 100, Energy: 0, Instrumental: false, Danceability: 0}
	}
	result := pc.create_playlist("mysterious", tracks)
	if len(result) != 10 {
		t.Fatalf("default cap: expected 10, got %d", len(result))
	}
}

func TestCreatePlaylistCapsWorkoutAt25(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 30)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i%26)), Tempo: 140, Energy: 8, Instrumental: false, Danceability: 0}
	}
	result := pc.create_playlist("workout", tracks)
	if len(result) != 25 {
		t.Fatalf("workout cap: expected 25, got %d", len(result))
	}
}

func TestCreatePlaylistCapsFocusAt30(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 35)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i%26)), Tempo: 0, Energy: 0, Instrumental: true, Danceability: 0}
	}
	result := pc.create_playlist("focus", tracks)
	if len(result) != 30 {
		t.Fatalf("focus cap: expected 30, got %d", len(result))
	}
}

func TestCreatePlaylistCapsPartyAt20(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := make([]Track, 25)
	for i := range tracks {
		tracks[i] = Track{Title: string(rune('A' + i%26)), Tempo: 120, Energy: 0, Instrumental: false, Danceability: 7}
	}
	result := pc.create_playlist("party", tracks)
	if len(result) != 20 {
		t.Fatalf("party cap: expected 20, got %d", len(result))
	}
}

func TestCreatePlaylistSortsHappyDescending(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 130, 0, false, 0}, {"B", 150, 0, false, 0}}
	result := pc.create_playlist("happy", tracks)
	if result[0].Title != "B" || result[1].Title != "A" {
		t.Fatal("happy sort")
	}
}

func TestCreatePlaylistSortsSadAscending(t *testing.T) {
	pc := NewPlaylistCurator()
	tracks := []Track{{"A", 80, 0, false, 0}, {"B", 60, 0, false, 0}}
	result := pc.create_playlist("sad", tracks)
	if result[0].Title != "B" || result[1].Title != "A" {
		t.Fatal("sad sort")
	}
}
