import { useNavigate } from 'react-router-dom';

import { Button } from '../../components/Button';
import { BASE_API_URL, getAccessTokens } from '../../utils';

const Home = () => {
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
		<div className="min-h-screen bg-gradient-to-br from-cyan-100 to-accent">
			<header className="flex items-center justify-between py-2 px-[2%] sm:px-6">
				<a className="flex items-center justify-center gap-2" href="#main">
					<img className="w-10" src="/logo.svg" alt="spotwave-logo" />
					<h1 className="font-bold">Spotwave</h1>
				</a>
				<Button variant="primary" classNames="btn btn-primary" onClick={onLogin}>
					{tokens && tokens.accessToken.length ? 'Dashboard' : 'Login'}
				</Button>
			</header>
			<main className="mt-56 px-[2%]">
				<section id="main" className="flex justify-center items-center flex-col">
					<h1 className="font-bold text-3xl sm:text-6xl sm:w-[60%] text-center">
						Download your Spotify playlist in just few clicks!
					</h1>
					<p className="mt-4 text-center">
						Sign in to your spotify account, and chose songs in your playlist to download!
					</p>
					<Button variant="primary" classNames="px-8 py-3 mt-8" onClick={onLogin}>
						Download Now!
					</Button>
				</section>
			</main>
		</div>
	);
};

export default Home;
