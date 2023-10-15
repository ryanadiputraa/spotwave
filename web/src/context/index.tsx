import { createContext, Dispatch, ReactNode, useReducer } from 'react';

import { MainState, MainAction, mainReducer } from './reducers/main';

interface IInitialState {
	main: MainState;
}
const initialState: IInitialState = {
	main: {
		user: null,
	},
};

const AppContext = createContext<{
	main: MainState;
	mainDispatch: Dispatch<MainAction>;
}>({
	main: initialState.main,
	mainDispatch: () => null,
});

const AppProvider: React.FC<{ children: ReactNode }> = ({ children }) => {
	const [mainState, mainDispatch] = useReducer(
		(main: MainState, action: MainAction) => mainReducer(main, action),
		initialState.main
	);

	return (
		<AppContext.Provider
			value={{
				main: mainState,
				mainDispatch: mainDispatch,
			}}
		>
			{children}
		</AppContext.Provider>
	);
};

export { AppContext, AppProvider };
