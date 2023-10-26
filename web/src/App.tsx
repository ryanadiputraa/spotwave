import { Suspense, lazy } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import { AppProvider } from './context';
import { Loading } from './components/Loading';
import { Toast } from './components/Toast';

const Home = lazy(() => import('./pages/Home'));
const Auth = lazy(() => import('./pages/Auth'));
const Dashboard = lazy(() => import('./pages/Dashboard'));
const NotFound = lazy(() => import('./pages/NotFound'));

const App = () => {
	return (
		<Suspense fallback={<Loading />}>
			<AppProvider>
				<Router>
					<Routes>
						<Route path="/" element={<Home />} />
						<Route path="/auth" element={<Auth />} />
						<Route path="/dashboard" element={<Dashboard />} />
						<Route path="*" element={<NotFound />} />
					</Routes>
				</Router>
				<Toast />
			</AppProvider>
		</Suspense>
	);
};

export default App;
