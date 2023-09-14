import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import './index.css';

function App() {
	return (
        <div className="text-white h-full min-h-screen pb-4">
			<Router>
				<RoutingComponent />
			</Router>
        </div>
	);
}
  
export default App;