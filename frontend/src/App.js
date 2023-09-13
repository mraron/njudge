import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Contests from './pages/Contests';
import Main from './pages/Main';
import Info from './pages/Info';
import Submissions from './pages/Submissions';
import Problems from './pages/Problems';
import Archive from './pages/Archive';
import Profile from './pages/Profile/Profile';
import ProfileMain from './pages/Profile/ProfileMain';
import ProfileSubmissions from './pages/Profile/ProfileSubmissions';
import ProfileSettings from './pages/Profile/ProfileSettings';
import Menubar from './components/Menubar';
import NotFound from './pages/Error/NotFound';
import Problem from './pages/Problem/Problem';
import ProblemStatement from './pages/Problem/ProblemStatement';
import ProblemSubmit from './pages/Problem/ProblemSubmit';
import ProblemSubmissions from './pages/Problem/ProblemSubmissions';
import Login from './pages/Auth/Login';
import Register from './pages/Auth/Register';
import ProblemRanklist from "./pages/Problem/ProblemRankings";
import Submission from "./pages/Submission";
import './index.css';
import {routeMap} from "./config/RouteConfig";

function App() {
	return (
        <div className="text-white h-full min-h-screen pb-4">
			<Router>
				<div className="pb-20">
					<Menubar />
				</div>
				<Routes>
					<Route path={routeMap.main} element={<Main />} />
					<Route path={routeMap.contests} element={<Contests />} />
					<Route path={routeMap.info} element={<Info />} />
					<Route path={routeMap.archive} element={<Archive />} />
					<Route path={routeMap.submissions} element={<Submissions />} />
					<Route path={routeMap.problems} element={<Problems />} />
					<Route path={routeMap.submission} element={<Submission />} />
					<Route path={routeMap.profile} element={<Profile />} >
						<Route path={routeMap.profile} element={<ProfileMain />} />
						<Route path={routeMap.profileSubmissions} element={<ProfileSubmissions />} />
						<Route path={routeMap.profileSettings} element={<ProfileSettings />} />
					</Route>
					<Route path={routeMap.problem} element={<Problem />} >
						<Route path={routeMap.problem} element={<ProblemStatement />} />
						<Route path={routeMap.problemSubmit} element={<ProblemSubmit />} />
						<Route path={routeMap.problemSubmissions} element={<ProblemSubmissions />} />
						<Route path={routeMap.problemRanklist} element={<ProblemRanklist />} />
					</Route>
					<Route path={routeMap.login} element={<Login />} />
					<Route path={routeMap.register} element={<Register />} />
					<Route path="*" element={<NotFound />} />
				</Routes>
			</Router>
        </div>
	);
}
  
export default App;