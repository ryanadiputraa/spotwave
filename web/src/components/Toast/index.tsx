import { useContext, useEffect } from 'react';

import { AppContext } from '../../context';
import { Icon } from './Icon';

const ERR_CODE: { [key: string]: string } = {
	bad_request: 'Invalid request',
	unauthorized: "You're not authorized",
	forbidden: 'Forbidden',
	not_found: 'Data not found',
	server_error: 'Please try again later',
	bad_gateway: 'Connection Timeout',
};

export const Toast = () => {
	const { main, mainDispatch } = useContext(AppContext);

	const onClose = () => mainDispatch({ type: 'TOGGLE_TOAST' });

	useEffect(() => {
		if (!main.toast.isOpen) return;
		const t = setTimeout(() => {
			onClose();
		}, 5000);

		return () => clearTimeout(t);
	}, [main.toast.isOpen]);

	return (
		<div
			id="toast-success"
			className={`${
				main.toast.isOpen ? 'flex' : 'hidden'
			} fixed bottom-1 right-4 items-center w-full max-w-xs p-4 mb-4 text-white bg-black rounded-lg shadow`}
			role="alert"
		>
			<Icon />
			<div className="ml-3 text-sm font-normal">{ERR_CODE[main.toast.message] ?? main.toast.message}</div>
			<button
				type="button"
				className="ml-auto -mx-1.5 -my-1.5 text-white rounded-lg p-1.5 inline-flex items-center justify-center h-8 w-8"
				data-dismiss-target="#toast-success"
				aria-label="Close"
				onClick={onClose}
			>
				<span className="sr-only">Close</span>
				<svg className="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
					<path
						stroke="currentColor"
						strokeLinecap="round"
						strokeLinejoin="round"
						strokeWidth="2"
						d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
					/>
				</svg>
			</button>
		</div>
	);
};
