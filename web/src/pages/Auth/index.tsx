import { useContext, useEffect } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import Cookies from 'universal-cookie';

import { AppContext } from '../../context';
import { AccessTokens } from '../../types/tokens';

const Auth = () => {
	const { mainDispatch } = useContext(AppContext);
	const [params] = useSearchParams();
	const navigate = useNavigate();

	const catchError = (err: string) => {
		mainDispatch({ type: 'TOGGLE_TOAST', isOpen: true, toastType: 'error', msg: err });
		navigate('/');
	};

	const setAuthCookies = (tokens: AccessTokens) => {
		const cookies = new Cookies(null, { path: '/' });
		cookies.set('auth', tokens);
		navigate('/dashboard');
	};

	useEffect(() => {
		const err = params.get('err');
		if (err) {
			catchError(err);
			return;
		}

		const accessToken = params.get('access_token') ?? '';
		const expiresIn = Number(params.get('expires_in')) ?? 0;
		const refreshToken = params.get('refresh_token') ?? '';
		setAuthCookies({ accessToken, expiresIn, refreshToken });
	}, []);

	return <></>;
};

export default Auth;
