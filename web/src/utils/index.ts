import Cookies from 'universal-cookie';

import { AccessTokens } from '../types/tokens';

export const BASE_API_URL = import.meta.env.VITE_BASE_API_URL;

export const getAccessTokens = (): AccessTokens | null => {
	const cookies = new Cookies();
	const tokens = cookies.get<AccessTokens | null>('auth');
	if (!tokens || !tokens.accessToken?.length) return null;
	return tokens;
};

export const setAccessTokens = (tokens: AccessTokens) => {
	const cookies = new Cookies();
	cookies.set('auth', tokens);
};
