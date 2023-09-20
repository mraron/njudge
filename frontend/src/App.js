import React from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import UserProvider from "./contexts/user/UserProvider";
import Menubar from "./components/concrete/other/Menubar";
import FlashEvent from "./components/util/flash/FlashEvent";
import {AnimatePresence, motion} from "framer-motion";
import './index.css';

window.flash = (message, type = "success") => FlashEvent.emit('flash', ({message, type}));

function App() {
    return (
        <UserProvider>
            <div className="text-white h-full min-h-screen pb-4">
                <FlashContainer/>
                <Router>
                    <AnimatePresence>
                        <motion.div initial={{opacity: 0.6}} animate={{opacity: 1, transition: {duration: 0.25}}}
                                    exit={{opacity: 0.6, transition: {duration: 0.25}}}>
                            <div className="pb-20">
                                <Menubar/>
                            </div>
                        </motion.div>
                    </AnimatePresence>
                    <RoutingComponent/>
                </Router>
            </div>
        </UserProvider>
    );
}

export default App;