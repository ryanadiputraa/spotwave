import { useContext, useEffect } from 'react';

import { AppBar } from './components/AppBar';
import { AppContext } from '../../context';
import { useMainAction } from '../../context/actions/main';
import { useSpotifyAction } from '../../context/actions/spotify';
import { PlaylistCard } from './components/PlaylistCard';

const Dashboard = () => {
	const { main, spotify } = useContext(AppContext);
	const { getUserInfo } = useMainAction();
	const { getUserPlaylists } = useSpotifyAction();

	useEffect(() => {
		if (!main.user) {
			getUserInfo();
			getUserPlaylists();
			return;
		}
	}, [main.user]);

	return (
		<div className="min-h-screen bg-gradient-to-br from-white to-accent">
			<AppBar />
			<main className="py-2 mt-8 px-[2%] sm:px-6">
				<h2 className="font-bold text-center text-2xl">User Playlists</h2>
				<div className="mt-4 flex flex-col items-center mx-auto w-full max-w-xl">
					{spotify.playlists?.items.map((playlist) => (
						<PlaylistCard key={playlist.id} playlist={playlist} />
					))}
				</div>
			</main>
		</div>
	);
};

export default Dashboard;
