import { useContext } from 'react';

import { AppContext } from '..';
import { SuccessResponse } from '../../types/api';
import { BASE_API_URL, getAccessTokens } from '../../utils';
import { PlaylistTracks, Playlists, Track } from '../../types/spotify';

export const useSpotifyAction = () => {
	const { spotify, spotifyDispatch } = useContext(AppContext);

	const getUserPlaylists = async (): Promise<void> => {
		const tokens = getAccessTokens();
		if (!tokens) {
			window.location.href = `${BASE_API_URL}/oauth/login`;
			return;
		}
		try {
			const resp = await fetch(`${BASE_API_URL}/api/spotify/playlists`, {
				headers: {
					Authorization: `Bearer ${tokens.accessToken}`,
				},
			});

			if (!resp.ok) {
				console.error(resp);
				return;
			}

			const json: SuccessResponse<Playlists> = await resp.json();
			spotifyDispatch({ type: 'SET_PLAYLISTS', payload: json.data });
		} catch (error) {
			console.error(error);
		}
	};

	const getPlaylistTracks = async (playlistId: string): Promise<void> => {
		if (spotify.tracks[playlistId]) return;

		const tokens = getAccessTokens();
		if (!tokens) {
			window.location.href = `${BASE_API_URL}/oauth/login`;
			return;
		}
		try {
			if (spotify.tracks[playlistId]) return;
			const resp = await fetch(`${BASE_API_URL}/api/spotify/playlists/tracks?playlist_id=${playlistId}`, {
				headers: {
					Authorization: `Bearer ${tokens.accessToken}`,
				},
			});

			if (!resp.ok) {
				console.error(resp);
				return;
			}

			const json: SuccessResponse<PlaylistTracks> = await resp.json();
			let tracks: Track[] = [];
			json.data.items.map((item) => tracks.push(item.track));
			spotifyDispatch({ type: 'SET_PLAYLIST_TRACKS', playlistId: playlistId, tracks: tracks });
		} catch (error) {
			console.error(error);
		}
	};

	return {
		getUserPlaylists,
		getPlaylistTracks,
	};
};
