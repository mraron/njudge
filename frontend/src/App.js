import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import UserProvider from "./contexts/user/UserProvider";
import Menubar from "./components/concrete/other/Menubar";
import FadeIn from "./components/util/FadeIn";
import LanguageProvider from "./contexts/language/LanguageProvider";
import FlashEvent from "./components/util/flash/FlashEvent";
import './index.css';

window.flash = (message, type = "success") => FlashEvent.emit('flash', ({message, type}));

function App() {
    return (
        <UserProvider>
            <LanguageProvider>
                <div className="text-white h-full min-h-screen pb-4">
                    <FlashContainer/>
                    <Router>
                        <FadeIn>
                            <div className="pb-20">
                                <Menubar/>
                            </div>
                        </FadeIn>
                        <RoutingComponent/>
                    </Router>
                </div>
            </LanguageProvider>
        </UserProvider>
    );
}

export default App;