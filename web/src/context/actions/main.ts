import { useContext } from 'react';

import { AppContext } from '..';
import { User } from '../../types/user';
import { SuccessResponse } from '../../types/api';
import { BASE_API_URL, getAccessTokens, setAccessTokens } from '../../utils';
import { RefreshAccessTokens } from '../../types/tokens';

export const useMainAction = () => {
	const { mainDispatch } = useContext(AppContext);

	const refreshAccessToken = async (): Promise<void> => {
		const tokens = getAccessTokens();

		if (!tokens) {
			window.location.href = `${BASE_API_URL}/oauth/login`;
			return;
		}
		try {
			const resp = await fetch(`${BASE_API_URL}/oauth/refresh_token?token=${tokens.refreshToken}`);
			if (!resp.ok) {
				window.location.href = `${BASE_API_URL}/oauth/login`;
			}

			const json: SuccessResponse<RefreshAccessTokens> = await resp.json();
			setAccessTokens({
				accessToken: json.data.access_token,
				expiresIn: json.data.expires_in,
				refreshToken: tokens.refreshToken,
			});
			window.location.reload();
		} catch (error) {
			if (error instanceof Error) {
				mainDispatch({ type: 'TOGGLE_TOAST', isOpen: true, toastType: 'error', msg: error.message });
			}
		}
	};

	const getUserInfo = async (): Promise<void> => {
		const tokens = getAccessTokens();
		if (!tokens) {
			window.location.href = `${BASE_API_URL}/oauth/login`;
			return;
		}
		try {
			const resp = await fetch(`${BASE_API_URL}/api/spotify/users`, {
				headers: {
					Authorization: `Bearer ${tokens.accessToken}`,
				},
			});

			if (!resp.ok) {
				await refreshAccessToken();
				return;
			}

			const json: SuccessResponse<User> = await resp.json();
			mainDispatch({ type: 'SET_USER', payload: json.data });
		} catch (error) {
			if (error instanceof Error) {
				mainDispatch({ type: 'TOGGLE_TOAST', isOpen: true, toastType: 'error', msg: error.message });
			}
		}
	};

	return {
		getUserInfo,
	};
};
