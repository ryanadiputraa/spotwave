import { useContext, useEffect } from 'react';

import { AppBar } from './components/AppBar';
import { AppContext } from '../../context';
import { useMainAction } from '../../context/actions/main';

const Dashboard = () => {
	const { main } = useContext(AppContext);
	const { getUserInfo } = useMainAction();

	useEffect(() => {
		if (main.user) return;
		getUserInfo();
	}, [main.user]);

	return (
		<div>
			<AppBar />
		</div>
	);
};

export default Dashboard;
