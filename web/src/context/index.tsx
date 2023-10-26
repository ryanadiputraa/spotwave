import { createContext, Dispatch, ReactNode, useReducer } from 'react';

import { MainState, MainAction, mainReducer } from './reducers/main';
import { SpotifyAction, spotifyReducer, SpotifyState } from './reducers/spotify';

interface IInitialState {
	main: MainState;
	spotify: SpotifyState;
}
const initialState: IInitialState = {
	main: {
		user: null,
		toast: {
			isOpen: false,
			type: 'success',
			message: '',
		},
	},
	spotify: {
		playlists: null,
		tracks: {},
	},
};

const AppContext = createContext<{
	main: MainState;
	mainDispatch: Dispatch<MainAction>;
	spotify: SpotifyState;
	spotifyDispatch: Dispatch<SpotifyAction>;
}>({
	main: initialState.main,
	mainDispatch: () => null,
	spotify: initialState.spotify,
	spotifyDispatch: () => null,
});

const AppProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
	const [mainState, mainDispatch] = useReducer(
		(main: MainState, action: MainAction) => mainReducer(main, action),
		initialState.main
	);
	const [spotifyState, spotifyDispatch] = useReducer(
		(spotify: SpotifyState, action: SpotifyAction) => spotifyReducer(spotify, action),
		initialState.spotify
	);

	return (
		<AppContext.Provider
			value={{
				main: mainState,
				mainDispatch: mainDispatch,
				spotify: spotifyState,
				spotifyDispatch: spotifyDispatch,
			}}
		>
			{children}
		</AppContext.Provider>
	);
};

export { AppContext, AppProvider };
