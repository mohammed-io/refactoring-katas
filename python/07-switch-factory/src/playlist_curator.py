class PlaylistCurator:
    def __init__(self):
        pass

    def create_playlist(self, mood, tracks):
        result = []

        if mood == "happy":
            for track in tracks:
                if track.get("tempo", 0) > 120:
                    result.append(track)
            result.sort(key=lambda x: x.get("tempo", 0), reverse=True)
            result = result[:20]

        elif mood == "sad":
            for track in tracks:
                if track.get("tempo", 0) < 90:
                    result.append(track)
            result.sort(key=lambda x: x.get("tempo", 0))
            result = result[:15]

        elif mood == "workout":
            for track in tracks:
                if track.get("tempo", 0) > 130 and track.get("energy", 0) > 7:
                    result.append(track)
            result.sort(key=lambda x: x.get("energy", 0), reverse=True)
            result = result[:25]

        elif mood == "focus":
            for track in tracks:
                if track.get("instrumental") is True:
                    result.append(track)
            result.sort(key=lambda x: x.get("tempo", 0))
            result = result[:30]

        elif mood == "party":
            for track in tracks:
                if track.get("tempo", 0) > 110 and track.get("danceability", 0) > 6:
                    result.append(track)
            result.sort(key=lambda x: x.get("danceability", 0), reverse=True)
            result = result[:20]

        else:
            for track in tracks:
                result.append(track)
            result = result[:10]

        return result
