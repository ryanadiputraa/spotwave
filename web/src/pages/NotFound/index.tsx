import { Link, useNavigate } from 'react-router-dom';

import { Button } from '../../components/Button';
import { BASE_API_URL, getAccessTokens } from '../../utils';

const NotFound = () => {
	const navigate = useNavigate();
	const tokens = getAccessTokens();

	const onLogin = () => {
		if (tokens && tokens.accessToken.length) {
			navigate('/dashboard');
			return;
		}
		window.location.href = `${BASE_API_URL}/oauth/login`;
	};

	return (
		<div className="min-h-screen bg-gradient-to-br from-white to-accent">
			<header className="flex items-center justify-between py-2 px-[2%] sm:px-6">
				<Link to={'/'} className="flex items-center justify-center gap-2">
					<img className="w-10" src="/logo.svg" alt="spotwave-logo" />
					<h1 className="font-bold">Spotwave</h1>
				</Link>
				<Button variant="primary" classNames="btn btn-primary" onClick={onLogin}>
					{tokens && tokens.accessToken.length ? 'Dashboard' : 'Login'}
				</Button>
			</header>
			<div className="flex flex-col items-center mt-40">
				<img className="w-3/4 sm:w-80" src="/not-found.svg" alt="not found" />
				<p className="font-bold text-xl text-center">This is not the web page you are looking for.</p>
			</div>
		</div>
	);
};

export default NotFound;
