# frozen_string_literal: true

class PlaylistCurator
  def create_playlist(mood, tracks)
    result = []
    case mood
    when 'happy'
      tracks.each { |t| result << t if t[:tempo] > 120 }
      result = result.sort_by { |t| -t[:tempo] }[0, 20]
    when 'sad'
      tracks.each { |t| result << t if t[:tempo] < 90 }
      result = result.sort_by { |t| t[:tempo] }[0, 15]
    when 'workout'
      tracks.each { |t| result << t if t[:tempo] > 130 && t[:energy] > 7 }
      result = result[0, 25]
    when 'focus'
      tracks.each { |t| result << t if t[:instrumental] == true }
      result = result[0, 30]
    when 'party'
      tracks.each { |t| result << t if t[:tempo] > 110 && t[:danceability] > 6 }
      result = result[0, 20]
    else
      tracks.each { |t| result << t }
      result = result[0, 10]
    end
    result
  end
end
