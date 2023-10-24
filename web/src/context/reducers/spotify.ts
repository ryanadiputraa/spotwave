import { Playlists, Track } from '../../types/spotify';

export const spotifyReducer = (state: SpotifyState, action: SpotifyAction) => {
	switch (action.type) {
		case 'SET_PLAYLISTS':
			return {
				...state,
				playlists: action.payload,
			};

		case 'SET_PLAYLIST_TRACKS':
			return {
				...state,
				tracks: {
					...state.tracks,
					[action.playlistId]: action.tracks,
				},
			};

		default:
			return { ...state };
	}
};

export interface SpotifyState {
	playlists: Playlists | null;
	tracks: {
		[playlistId: string]: Track[];
	};
}

export type SpotifyAction =
	| { type: 'SET_PLAYLISTS'; payload: Playlists | null }
	| { type: 'SET_PLAYLIST_TRACKS'; playlistId: string; tracks: Track[] };
