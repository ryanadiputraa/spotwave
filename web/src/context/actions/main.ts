import { useContext } from 'react';

import { AppContext } from '..';
import { User } from '../../types/user';
import { SuccessResponse } from '../../types/api';
import { BASE_API_URL, getAccessTokens } from '../../utils';

export const useMainAction = () => {
	const { mainDispatch } = useContext(AppContext);

	const getUserInfo = async (): Promise<void> => {
		const tokens = getAccessTokens();
		if (!tokens) {
			window.location.href = `${BASE_API_URL}/oauth/login`;
			return;
		}
		try {
			const resp = await fetch(`${BASE_API_URL}/api/spotify/user`, {
				headers: {
					Authorization: `Bearer ${tokens.accessToken}`,
				},
			});
			const json: SuccessResponse<User> = await resp.json();
			mainDispatch({ type: 'SET_USER', payload: json.data });
		} catch (error) {
			// TODO: catch toast error
			console.error(error);
		}
	};

	return {
		getUserInfo,
	};
};
