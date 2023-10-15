import { User } from '../../types/user';

export const mainReducer = (state: MainState, action: MainAction) => {
	switch (action.type) {
		case 'SET_USER':
			return {
				...state,
				user: action.payload,
			};

		default:
			return { ...state };
	}
};

export interface MainState {
	user: User | null;
}

export type MainAction = { type: 'SET_USER'; payload: User | null };
