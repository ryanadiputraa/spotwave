import { useContext } from 'react';

import { AppContext } from '..';
import { SuccessResponse } from '../../types/api';
import { BASE_API_URL, getAccessTokens } from '../../utils';
import { Playlists } from '../../types/spotify';

export const useSpotifyAction = () => {
	const { spotifyDispatch } = useContext(AppContext);

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

	return {
		getUserPlaylists,
	};
};
