import Cookies from 'universal-cookie';

import { AccessTokens } from '../../types/tokens';
import { User } from '../../types/user';
import { SuccessResponse } from '../../types/api';

export const getAccessTokens = (): AccessTokens => {
	const cookies = new Cookies();
	return cookies.get<AccessTokens>('auth');
};

export const getUserInfo = async (): Promise<User | null> => {
	const tokens = getAccessTokens();
	try {
		const resp = await fetch(`${import.meta.env.VITE_BASE_API_URL}/api/spotify/user`, {
			headers: {
				Authorization: `Bearer ${tokens.accessToken}`,
			},
		});
		const json: SuccessResponse<User> = await resp.json();
		return json.data;
	} catch (error) {
		// TODO: catch toast error
		console.error(error);
		return null;
	}
};
