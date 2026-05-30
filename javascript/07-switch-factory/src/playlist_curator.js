class PlaylistCurator {
  constructor() {}

  create_playlist(mood, tracks) {
    let result = [];

    switch (mood) {
      case 'happy':
        for (let i = 0; i < tracks.length; i++) {
          if (tracks[i].tempo > 120) {
            result.push(tracks[i]);
          }
        }
        result.sort((a, b) => b.tempo - a.tempo);
        if (result.length > 20) {
          result = result.slice(0, 20);
        }
        break;
      case 'sad':
        for (let i = 0; i < tracks.length; i++) {
          if (tracks[i].tempo < 90) {
            result.push(tracks[i]);
          }
        }
        result.sort((a, b) => a.tempo - b.tempo);
        if (result.length > 15) {
          result = result.slice(0, 15);
        }
        break;
      case 'workout':
        for (let i = 0; i < tracks.length; i++) {
          if (tracks[i].tempo > 130 && tracks[i].energy > 7) {
            result.push(tracks[i]);
          }
        }
        if (result.length > 25) {
          result = result.slice(0, 25);
        }
        break;
      case 'focus':
        for (let i = 0; i < tracks.length; i++) {
          if (tracks[i].instrumental === true) {
            result.push(tracks[i]);
          }
        }
        if (result.length > 30) {
          result = result.slice(0, 30);
        }
        break;
      case 'party':
        for (let i = 0; i < tracks.length; i++) {
          if (tracks[i].tempo > 110 && tracks[i].danceability > 6) {
            result.push(tracks[i]);
          }
        }
        if (result.length > 20) {
          result = result.slice(0, 20);
        }
        break;
      default:
        for (let i = 0; i < tracks.length; i++) {
          result.push(tracks[i]);
        }
        if (result.length > 10) {
          result = result.slice(0, 10);
        }
    }

    return result;
  }
}

export { PlaylistCurator };
