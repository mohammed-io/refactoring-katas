using Xunit;
using System.Linq;

public class PlaylistCuratorTest
{
    [Fact]
    public void FiltersHappyTracks()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("happy", new List<Track>
        {
            new Track { Title = "A", Tempo = 130 },
            new Track { Title = "B", Tempo = 100 }
        });
        Assert.Single(result);
        Assert.Equal("A", result[0].Title);
    }

    [Fact]
    public void FiltersSadTracks()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("sad", new List<Track>
        {
            new Track { Title = "A", Tempo = 80 },
            new Track { Title = "B", Tempo = 100 }
        });
        Assert.Single(result);
        Assert.Equal("A", result[0].Title);
    }

    [Fact]
    public void FiltersWorkoutTracks()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("workout", new List<Track>
        {
            new Track { Title = "A", Tempo = 140, Energy = 8 },
            new Track { Title = "B", Tempo = 140, Energy = 5 }
        });
        Assert.Single(result);
        Assert.Equal("A", result[0].Title);
    }

    [Fact]
    public void FiltersFocusTracks()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("focus", new List<Track>
        {
            new Track { Title = "A", Instrumental = true },
            new Track { Title = "B", Instrumental = false }
        });
        Assert.Single(result);
        Assert.Equal("A", result[0].Title);
    }

    [Fact]
    public void FiltersPartyTracks()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("party", new List<Track>
        {
            new Track { Title = "A", Tempo = 120, Danceability = 7 },
            new Track { Title = "B", Tempo = 100, Danceability = 5 }
        });
        Assert.Single(result);
        Assert.Equal("A", result[0].Title);
    }

    [Fact]
    public void CapsHappyPlaylistAt20()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 25)
                              .Select(i => new Track { Title = i.ToString(), Tempo = 130 })
                              .ToList();
        var result = curator.create_playlist("happy", tracks);
        Assert.Equal(20, result.Count);
    }

    [Fact]
    public void CapsSadPlaylistAt15()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 20)
                              .Select(i => new Track { Title = i.ToString(), Tempo = 80 })
                              .ToList();
        var result = curator.create_playlist("sad", tracks);
        Assert.Equal(15, result.Count);
    }

    [Fact]
    public void DefaultsToFirst10ForUnknownMood()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 20)
                              .Select(i => new Track { Title = i.ToString(), Tempo = 100 })
                              .ToList();
        var result = curator.create_playlist("mysterious", tracks);
        Assert.Equal(10, result.Count);
    }

    [Fact]
    public void CapsWorkoutPlaylistAt25()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 30)
                              .Select(i => new Track { Title = i.ToString(), Tempo = 140, Energy = 8 })
                              .ToList();
        var result = curator.create_playlist("workout", tracks);
        Assert.Equal(25, result.Count);
    }

    [Fact]
    public void CapsFocusPlaylistAt30()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 35)
                              .Select(i => new Track { Title = i.ToString(), Instrumental = true })
                              .ToList();
        var result = curator.create_playlist("focus", tracks);
        Assert.Equal(30, result.Count);
    }

    [Fact]
    public void CapsPartyPlaylistAt20()
    {
        var curator = new PlaylistCurator();
        var tracks = Enumerable.Range(0, 25)
                              .Select(i => new Track { Title = i.ToString(), Tempo = 120, Danceability = 7 })
                              .ToList();
        var result = curator.create_playlist("party", tracks);
        Assert.Equal(20, result.Count);
    }

    [Fact]
    public void SortsHappyByDescendingTempo()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("happy", new List<Track>
        {
            new Track { Title = "A", Tempo = 130 },
            new Track { Title = "B", Tempo = 150 }
        });
        Assert.Equal("B", result[0].Title);
        Assert.Equal("A", result[1].Title);
    }

    [Fact]
    public void SortsSadByAscendingTempo()
    {
        var curator = new PlaylistCurator();
        var result = curator.create_playlist("sad", new List<Track>
        {
            new Track { Title = "A", Tempo = 80 },
            new Track { Title = "B", Tempo = 60 }
        });
        Assert.Equal("B", result[0].Title);
        Assert.Equal("A", result[1].Title);
    }
}
