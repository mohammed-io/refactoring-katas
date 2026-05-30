package kata

import "sort"

type Track struct {
	Title        string
	Tempo        int
	Energy       int
	Instrumental bool
	Danceability int
}

type PlaylistCurator struct{}

func NewPlaylistCurator() *PlaylistCurator {
	return &PlaylistCurator{}
}

func (pc *PlaylistCurator) create_playlist(mood string, tracks []Track) []Track {
	result := []Track{}
	switch mood {
	case "happy":
		for _, t := range tracks {
			if t.Tempo > 120 {
				result = append(result, t)
			}
		}
		sort.Slice(result, func(i, j int) bool { return result[i].Tempo > result[j].Tempo })
		if len(result) > 20 {
			result = result[:20]
		}
	case "sad":
		for _, t := range tracks {
			if t.Tempo < 90 {
				result = append(result, t)
			}
		}
		sort.Slice(result, func(i, j int) bool { return result[i].Tempo < result[j].Tempo })
		if len(result) > 15 {
			result = result[:15]
		}
	case "workout":
		for _, t := range tracks {
			if t.Tempo > 130 && t.Energy > 7 {
				result = append(result, t)
			}
		}
		if len(result) > 25 {
			result = result[:25]
		}
	case "focus":
		for _, t := range tracks {
			if t.Instrumental {
				result = append(result, t)
			}
		}
		if len(result) > 30 {
			result = result[:30]
		}
	case "party":
		for _, t := range tracks {
			if t.Tempo > 110 && t.Danceability > 6 {
				result = append(result, t)
			}
		}
		if len(result) > 20 {
			result = result[:20]
		}
	default:
		result = append(result, tracks...)
		if len(result) > 10 {
			result = result[:10]
		}
	}
	return result
}
