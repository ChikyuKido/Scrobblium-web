package track

import (
	"github.com/gin-gonic/gin"
	"scrobblium-web/database/entity"
	"scrobblium-web/database/repo"
	"sort"
	"strconv"
	"strings"
)

type CombinedTrack struct {
	Artists     []string `json:"artists"`
	Album       string   `json:"album"`
	Track       string   `json:"track"`
	ListenedAt  int64    `json:"listenedAt"`
	ListenCount int64    `json:"listenCount"`
	ListenTime  int64    `json:"listenTime"`
}

func List() gin.HandlerFunc {
	return func(c *gin.Context) {
		combineBy := c.Query("combine")

		tracks, _ := repo.GetTracksWithAll()

		type key struct {
			Artist string
			Track  string
			Album  string
		}
		type aggregateTrack struct {
			Track entity.Track
			Count int64
		}

		resultMap := make(map[key]*aggregateTrack)

		var i = 0
		for _, t := range tracks {
			var artistNames []string
			for _, a := range t.Artists {
				artistNames = append(artistNames, a.Name)
			}

			sort.Strings(artistNames)
			artistKey := strings.Join(artistNames, ",")

			k := key{
				Artist: artistKey,
				Track:  t.Name,
				Album:  t.Album.Name,
			}

			var groupKey key
			switch combineBy {
			case "artist":
				groupKey = key{Artist: k.Artist}
			case "track":
				groupKey = key{Track: k.Track}
			case "album":
				groupKey = key{Album: k.Album}
			default:
				groupKey = key{Artist: strconv.Itoa(i)}
				i++
			}

			if existing, ok := resultMap[groupKey]; ok {
				existing.Track.Progress += t.Progress
				existing.Track.TimeListened += t.TimeListened
				existing.Track.MaxProgress += t.MaxProgress
				if t.ListenedAt < existing.Track.ListenedAt {
					existing.Track.ListenedAt = t.ListenedAt
				}
				existing.Count++
			} else {
				copyTrack := t
				resultMap[groupKey] = &aggregateTrack{
					Track: copyTrack,
					Count: 1,
				}
			}

		}

		combinedResult := make([]CombinedTrack, 0, len(resultMap))
		for _, ag := range resultMap {
			var artistNames []string
			for _, a := range ag.Track.Artists {
				artistNames = append(artistNames, a.Name)
			}
			sort.Strings(artistNames)

			combinedResult = append(combinedResult, CombinedTrack{
				Artists:     artistNames,
				Album:       ag.Track.Album.Name,
				Track:       ag.Track.Name,
				ListenedAt:  ag.Track.ListenedAt,
				ListenCount: ag.Count,
				ListenTime:  ag.Track.TimeListened,
			})
		}

		c.JSON(200, combinedResult)

	}
}
