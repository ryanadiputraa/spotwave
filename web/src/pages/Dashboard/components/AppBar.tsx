import { useContext } from 'react';

import { AppContext } from '../../../context';

export const AppBar = () => {
	const { main } = useContext(AppContext);

	return <header></header>;
};
