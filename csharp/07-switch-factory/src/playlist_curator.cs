public class Track
{
    public string Title = "";
    public int Tempo;
    public int Energy;
    public bool Instrumental;
    public int Danceability;
}

public class PlaylistCurator
{
    public List<Track> create_playlist(string mood, List<Track> tracks)
    {
        var result = new List<Track>();
        switch (mood)
        {
            case "happy":
                result = tracks.Where(t => t.Tempo > 120)
                              .OrderByDescending(t => t.Tempo)
                              .Take(20)
                              .ToList();
                break;

            case "sad":
                result = tracks.Where(t => t.Tempo < 90)
                              .OrderBy(t => t.Tempo)
                              .Take(15)
                              .ToList();
                break;

            case "workout":
                result = tracks.Where(t => t.Tempo > 130 && t.Energy > 7)
                              .Take(25)
                              .ToList();
                break;

            case "focus":
                result = tracks.Where(t => t.Instrumental)
                              .Take(20)
                              .ToList();
                break;

            case "party":
                result = tracks.Where(t => t.Tempo >= 120 && t.Danceability >= 6)
                              .Take(30)
                              .ToList();
                break;

            default:
                result = tracks.Take(10).ToList();
                break;
        }
        return result;
    }
}
