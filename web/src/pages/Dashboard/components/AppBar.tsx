import { useContext } from 'react';

import { AppContext } from '../../../context';

export const AppBar = () => {
	const { main } = useContext(AppContext);
	console.log(main.user);

	return <header></header>;
};
