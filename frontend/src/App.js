import React, {useContext, useEffect} from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import Menubar from "./components/concrete/other/Menubar";
import {getLanguages, getTags, getVerdicts} from "./util/fetchJudgeData";
import FlashEvent from "./components/util/flash/FlashEvent";
import {AnimatePresence, motion} from "framer-motion";
import JudgeDataContext from "./contexts/judgeData/JudgeDataContext";
import './index.css';

window.flash = (message, type = "success") => FlashEvent.emit('flash', ({message, type}));

function App() {
    const {judgeData, setJudgeData} = useContext(JudgeDataContext)

    useEffect(() => {
        const fetchJudgeData = async () => {
            await getLanguages().then(resp => {
                setJudgeData(prevJudgeData => { return {...prevJudgeData, languages: resp} })
            })
            await getVerdicts().then(resp => {
                setJudgeData(prevJudgeData => { return {...prevJudgeData, verdicts: resp} })
            })
            await getTags().then(resp => {
                setJudgeData(prevJudgeData => { return {...prevJudgeData, tags: resp} })
            })
        }
        fetchJudgeData()
    }, []);

    useEffect(() => {
        console.log(JSON.stringify(judgeData))
    }, [judgeData]);

    return (
        judgeData && judgeData.languages && judgeData.verdicts && judgeData.tags &&
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
    );
}

export default App;