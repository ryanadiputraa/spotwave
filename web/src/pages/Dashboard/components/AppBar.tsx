import { useContext } from 'react';

import { AppContext } from '../../../context';
import { Link } from 'react-router-dom';

export const AppBar = () => {
	const { main } = useContext(AppContext);
	const profileImage = main.user?.images.length ? main.user.images[0].url : '';
	const displayName = main.user?.display_name ?? '';
	const email = main.user?.email ?? '';

	return (
		<header className="flex items-center justify-between py-2 px-[2%] sm:px-6">
			<Link className="flex items-center justify-center gap-2" to={'/'}>
				<img className="w-10" src="/logo.svg" alt="spotwave-logo" />
				<h1 className="font-bold">Spotwave</h1>
			</Link>
			<div className="flex items-center justify-center gap-2">
				<img className="w-10 rounded-full" src={profileImage} alt="spotwave-logo" />
				<div className="hidden sm:flex flex-col items-start justify-center">
					<span className="font-bold">{displayName}</span>
					<span className="font-light text-sm">{email}</span>
				</div>
			</div>
		</header>
	);
};
