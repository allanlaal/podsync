package feeds

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mxpv/podsync/pkg/api"
)

func TestParseYoutubeURL_Playlist(t *testing.T) {
	link, _ := url.ParseRequestURI("https://www.youtube.com/playlist?list=PLCB9F975ECF01953C")
	kind, id, err := parseYoutubeURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypePlaylist, kind)
	require.Equal(t, "PLCB9F975ECF01953C", id)

	link, _ = url.ParseRequestURI("https://www.youtube.com/watch?v=rbCbho7aLYw&list=PLMpEfaKcGjpWEgNtdnsvLX6LzQL0UC0EM")
	kind, id, err = parseYoutubeURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypePlaylist, kind)
	require.Equal(t, "PLMpEfaKcGjpWEgNtdnsvLX6LzQL0UC0EM", id)
}

func TestParseYoutubeURL_Channel(t *testing.T) {
	link, _ := url.ParseRequestURI("https://www.youtube.com/channel/UC5XPnUk8Vvv_pWslhwom6Og")
	kind, id, err := parseYoutubeURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeChannel, kind)
	require.Equal(t, "UC5XPnUk8Vvv_pWslhwom6Og", id)

	link, _ = url.ParseRequestURI("https://www.youtube.com/channel/UCrlakW-ewUT8sOod6Wmzyow/videos")
	kind, id, err = parseYoutubeURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeChannel, kind)
	require.Equal(t, "UCrlakW-ewUT8sOod6Wmzyow", id)
}

func TestParseYoutubeURL_User(t *testing.T) {
	link, _ := url.ParseRequestURI("https://youtube.com/user/fxigr1")
	kind, id, err := parseYoutubeURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeUser, kind)
	require.Equal(t, "fxigr1", id)
}

func TestParseYoutubeURL_InvalidLink(t *testing.T) {
	link, _ := url.ParseRequestURI("https://www.youtube.com/user///")
	_, _, err := parseYoutubeURL(link)
	require.Error(t, err)

	link, _ = url.ParseRequestURI("https://www.youtube.com/channel//videos")
	_, _, err = parseYoutubeURL(link)
	require.Error(t, err)
}

func TestParseVimeoURL_Group(t *testing.T) {
	link, _ := url.ParseRequestURI("https://vimeo.com/groups/109")
	kind, id, err := parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeGroup, kind)
	require.Equal(t, "109", id)

	link, _ = url.ParseRequestURI("http://vimeo.com/groups/109")
	kind, id, err = parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeGroup, kind)
	require.Equal(t, "109", id)

	link, _ = url.ParseRequestURI("http://www.vimeo.com/groups/109")
	kind, id, err = parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeGroup, kind)
	require.Equal(t, "109", id)

	link, _ = url.ParseRequestURI("https://vimeo.com/groups/109/videos/")
	kind, id, err = parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeGroup, kind)
	require.Equal(t, "109", id)
}

func TestParseVimeoURL_Channel(t *testing.T) {
	link, _ := url.ParseRequestURI("https://vimeo.com/channels/staffpicks")
	kind, id, err := parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeChannel, kind)
	require.Equal(t, "staffpicks", id)

	link, _ = url.ParseRequestURI("http://vimeo.com/channels/staffpicks/146224925")
	kind, id, err = parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeChannel, kind)
	require.Equal(t, "staffpicks", id)
}

func TestParseVimeoURL_User(t *testing.T) {
	link, _ := url.ParseRequestURI("https://vimeo.com/awhitelabelproduct")
	kind, id, err := parseVimeoURL(link)
	require.NoError(t, err)
	require.Equal(t, api.LinkTypeUser, kind)
	require.Equal(t, "awhitelabelproduct", id)
}

func TestParseVimeoURL_InvalidLink(t *testing.T) {
	link, _ := url.ParseRequestURI("http://www.apple.com")
	_, _, err := parseVimeoURL(link)
	require.Error(t, err)

	link, _ = url.ParseRequestURI("http://www.vimeo.com")
	_, _, err = parseVimeoURL(link)
	require.Error(t, err)
}
