import { useContext, useEffect, useState } from 'react';

import { AppBar } from './components/AppBar';
import { AppContext } from '../../context';
import { useMainAction } from '../../context/actions/main';
import { useSpotifyAction } from '../../context/actions/spotify';
import { PlaylistCard } from './components/PlaylistCard';
import { PlaylistItem } from '../../types/spotify';
import { Button } from '../../components/Button';
import { TrackCard } from './components/TrackCard';

const Dashboard = () => {
	const [selectedPlaylist, setSelectedPlaylist] = useState<PlaylistItem | null>(null);
	const { main, spotify } = useContext(AppContext);
	const { getUserInfo } = useMainAction();
	const { getUserPlaylists, getPlaylistTracks } = useSpotifyAction();

	useEffect(() => {
		if (!main.user) {
			getUserInfo();
			getUserPlaylists();
			return;
		}
	}, [main.user]);

	useEffect(() => {
		if (!selectedPlaylist) return;
		getPlaylistTracks(selectedPlaylist.id);
	}, [selectedPlaylist]);

	return (
		<div className="min-h-screen bg-gradient-to-br from-white to-accent">
			<AppBar />
			<main className="py-2 mt-8 px-[2%] sm:px-6">
				<h2 className="font-bold text-center text-2xl">
					{selectedPlaylist ? selectedPlaylist.name : 'Your Spotify Playlists'}
				</h2>
				{selectedPlaylist ? (
					<div className="mx-auto flex flex-col items-center justify-start max-w-5xl">
						<Button variant="primary" classNames="self-start" onClick={() => setSelectedPlaylist(null)}>
							Back
						</Button>
						<div className="mt-8 w-full flex flex-col gap-2">
							{spotify.tracks[selectedPlaylist.id]?.map((track, i) => (
								<TrackCard key={track.id} num={i + 1} track={track} />
							))}
						</div>
					</div>
				) : (
					<div className="mt-4 flex flex-col items-center mx-auto w-full max-w-xl gap-4">
						{spotify.playlists?.items.map((playlist) => (
							<PlaylistCard key={playlist.id} playlist={playlist} setSelectedPlaylist={setSelectedPlaylist} />
						))}
					</div>
				)}
			</main>
		</div>
	);
};

export default Dashboard;
