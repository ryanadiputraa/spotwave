import { Dispatch } from 'react';

import { PlaylistItem } from '../../../types/spotify';

interface Props {
	playlist: PlaylistItem;
	setSelectedPlaylist: Dispatch<PlaylistItem>;
}

export const PlaylistCard = ({ playlist, setSelectedPlaylist }: Props) => {
	const image = playlist.images.length ? playlist.images[0].url : '';

	return (
		<div
			className="flex items-center gap-4 w-full shadow-lg bg-black text-white rounded-md overflow-hidden cursor-pointer btn"
			onClick={() => setSelectedPlaylist(playlist)}
		>
			<img className="w-20" src={image} alt={playlist.name ?? ''} />
			<h4>{playlist.name}</h4>
		</div>
	);
};
