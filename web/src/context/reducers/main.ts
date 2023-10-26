import { User } from '../../types/user';

export const mainReducer = (state: MainState, action: MainAction) => {
	switch (action.type) {
		case 'SET_USER':
			return {
				...state,
				user: action.payload,
			};

		case 'TOGGLE_TOAST':
			return {
				...state,
				toast: {
					isOpen: action.isOpen ?? false,
					type: action.toastType ?? 'success',
					message: action.msg ?? '',
				},
			};

		default:
			return { ...state };
	}
};

export interface MainState {
	user: User | null;
	toast: ToastState;
}

interface ToastState {
	isOpen: boolean;
	type: 'success' | 'warning' | 'error';
	message: string;
}

export type MainAction =
	| { type: 'SET_USER'; payload: User | null }
	| { type: 'TOGGLE_TOAST'; isOpen?: boolean; toastType?: 'success' | 'warning' | 'error'; msg?: string };
