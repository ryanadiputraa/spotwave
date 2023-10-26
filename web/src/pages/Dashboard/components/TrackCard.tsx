import { useContext, useState } from 'react';

import { AppContext } from '../../../context';
import { Button } from '../../../components/Button';
import { DownloadLink, Track } from '../../../types/spotify';
import { BASE_API_URL } from '../../../utils';
import { SuccessResponse } from '../../../types/api';

interface Props {
	track: Track;
	num: number;
}

export const TrackCard = ({ track, num }: Props) => {
	const { mainDispatch } = useContext(AppContext);
	const [donwloadLink, setDownloadLink] = useState<string | null>(null);
	const [isLoading, setIsLoading] = useState<boolean>(false);

	const seconds = Math.floor((track.duration_ms / 1000) % 60);
	const minutes = Math.floor((track.duration_ms / (1000 * 60)) % 60);

	const onConvert = async () => {
		setIsLoading(true);
		const artists = track.artists.map((artist) => `${artist.name}`).join(' ');
		try {
			const resp = await fetch(
				`${BASE_API_URL}/api/spotify/tracks/download?artists=${encodeURI(artists)}&title=${encodeURI(track.name)}`
			);
			if (!resp.ok) {
				mainDispatch({ type: 'TOGGLE_TOAST', isOpen: true, toastType: 'error', msg: 'Please try again later' });
				return;
			}

			const json: SuccessResponse<DownloadLink> = await resp.json();
			setDownloadLink(json.data.link);
		} catch (error) {
			if (error instanceof Error) {
				mainDispatch({ type: 'TOGGLE_TOAST', isOpen: true, toastType: 'error', msg: error.message });
			}
		} finally {
			setIsLoading(false);
		}
	};

	const onDownload = () => {
		if (donwloadLink) window.open(donwloadLink);
	};

	return (
		<div className="flex items-center gap-2 w-full shadow-lg bg-black text-white rounded-md overflow-hidden cursor-pointer btn py-4 px-2 sm:px-8">
			<span className="text-xs sm:text-sm w-[10%] sm:w-[5%] text-center sm:text-left">{num}</span>
			<div className="flex flex-col gap-1 w-[50%] sm:w-[65%]">
				<h4 className="font-bold text-sm sm:text-base">{track.name}</h4>
				<span className="text-sm">
					{track.artists.map(
						(artist, i) => artist.name + (track.artists.length > 1 && track.artists.length !== i + 1 ? ', ' : '')
					)}
				</span>
			</div>
			<span className="hidden sm:inline-block w-[15%] text-center">
				{minutes}:{seconds}
			</span>
			<div className="w-[40%] sm:w-[15%] text-center">
				{donwloadLink ? (
					<Button variant="secondary" classNames="font-bold text-sm sm:text-base px-3" onClick={onDownload}>
						Download
					</Button>
				) : (
					<Button variant="secondary" classNames="font-bold text-sm sm:text-base px-3 text-center" onClick={onConvert}>
						{isLoading ? (
							<div role="status">
								<svg
									aria-hidden="true"
									className="w-8 h-8 text-accent animate-spin fill-black"
									viewBox="0 0 100 101"
									fill="none"
									xmlns="http://www.w3.org/2000/svg"
								>
									<path
										d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
										fill="currentColor"
									/>
									<path
										d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
										fill="currentFill"
									/>
								</svg>
								<span className="sr-only">Loading...</span>
							</div>
						) : (
							'Get Link'
						)}
					</Button>
				)}
			</div>
		</div>
	);
};
