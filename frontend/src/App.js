import React, {useEffect} from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import FlashEvent from "./components/util/flash/FlashEvent";
import './index.css';

window.flash = (message, type = "success") => FlashEvent.emit('flash', ({message, type}));

function App() {
    return (
        <div className="text-white h-full min-h-screen pb-4">
            <Router>
                <RoutingComponent/>
            </Router>
            <FlashContainer />
        </div>
    );
}

export default App;