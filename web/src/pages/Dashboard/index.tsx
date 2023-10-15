import { useContext, useEffect } from 'react';

import { AppBar } from './components/AppBar';
import { AppContext } from '../../context';
import { getUserInfo } from '../../context/actions/main';

const Dashboard = () => {
	const { main, mainDispatch } = useContext(AppContext);

	useEffect(() => {
		if (main.user) return;
		const fetchUser = async () => {
			const user = await getUserInfo();
			mainDispatch({ type: 'SET_USER', payload: user });
		};
		fetchUser();
	}, []);

	return (
		<div>
			<AppBar />
		</div>
	);
};

export default Dashboard;
