import { Playlists } from '../../types/spotify';

export const spotifyReducer = (state: SpotifyState, action: SpotifyAction) => {
	switch (action.type) {
		case 'SET_PLAYLISTS':
			return {
				...state,
				playlists: action.payload,
			};

		default:
			return { ...state };
	}
};

export interface SpotifyState {
	playlists: Playlists | null;
}

export type SpotifyAction = { type: 'SET_PLAYLISTS'; payload: Playlists | null };
