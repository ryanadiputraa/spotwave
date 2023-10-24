import { Button } from '../../../components/Button';
import { Track } from '../../../types/spotify';

interface Props {
	track: Track;
	num: number;
}

export const TrackCard = ({ track, num }: Props) => {
	const seconds = Math.floor((track.duration_ms / 1000) % 60);
	const minutes = Math.floor((track.duration_ms / (1000 * 60)) % 60);

	return (
		<div className="flex items-center gap-1 w-full shadow-lg bg-black text-white rounded-md overflow-hidden cursor-pointer btn py-4 px-2 sm:px-8">
			<span className="text-xs sm:text-sm w-[10%]"># {num}</span>
			<div className="flex flex-col gap-1 w-[60%]">
				<h4 className="font-bold">{track.name}</h4>
				<span className="text-sm">
					{track.artists.map(
						(artist, i) => artist.name + (track.artists.length > 1 && track.artists.length !== i + 1 ? ', ' : '')
					)}
				</span>
			</div>
			<span className="hidden sm:inline-block w-[20%]">
				{minutes}:{seconds}
			</span>
			<Button variant="secondary" classNames="font-bold">
				Download
			</Button>
		</div>
	);
};
