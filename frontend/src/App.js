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
import ProblemResults from './pages/Problem/ProblemResults';
import Register from './pages/Auth/Register';
import './index.css';

function App() {
	return (
        <div className="text-white h-full min-h-screen pb-4">
			<Router>
				<div className="pb-20">
					<Menubar />
				</div>
				<Routes>
					<Route path="/" element={<Main />} />
					<Route path="/contests/" element={<Contests />} />
					<Route path="/info/" element={<Info />} />
					<Route path="/archive/" element={<Archive />} />
					<Route path="/problemset/status/" element={<Submissions />} />
					<Route path="/problemset/main/" element={<Problems />} />
					<Route path="/user/profile/" element={<Profile />} >
						<Route path="/user/profile/" element={<ProfileMain />} />
						<Route path="/user/profile/submissions/" element={<ProfileSubmissions />} />
						<Route path="/user/profile/settings/" element={<ProfileSettings />} />
					</Route>
					<Route path="/user/profile/submissions/" element={<Profile />} />
					<Route path="/user/profile/settings/" element={<Profile />} />
					<Route path="*" element={<NotFound />} />
					<Route path="/problemset/main/:task/" element={<Problem />} >
						<Route path="/problemset/main/:task/" element={<ProblemStatement />} />
						<Route path="/problemset/main/:task/submit/" element={<ProblemSubmit />} />
						<Route path="/problemset/main/:task/status/" element={<ProblemSubmissions />} />
						<Route path="/problemset/main/:task/ranklist/" element={<ProblemResults />} />
					</Route>
					<Route path="/user/login" element={<Login />} />
					<Route path="/user/register" element={<Register />} />
				</Routes>
			</Router>
        </div>
	);
}
  
export default App;