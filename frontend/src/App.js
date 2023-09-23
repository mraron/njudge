import React, {useContext, useEffect} from 'react';
import {BrowserRouter as Router} from 'react-router-dom';
import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import Menubar from "./components/concrete/other/Menubar";
import {getCategories, getHighlightCodes, getLanguages, getTags} from "./util/getJudgeData";
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
                if (resp.success) {
                    setJudgeData(prevJudgeData => {
                        return {...prevJudgeData, languages: resp.languages}
                    })
                }
            })
            await getCategories().then(resp => {
                if (resp.success) {
                    setJudgeData(prevJudgeData => {
                        return {...prevJudgeData, categories: resp.categories}
                    })
                }
            })
            await getTags().then(resp => {
                if (resp.success) {
                    setJudgeData(prevJudgeData => {
                        return {...prevJudgeData, tags: resp.tags}
                    })
                }
            })
        }
        fetchJudgeData().then(
            setJudgeData(prevJudgeData => {
                return {...prevJudgeData, highlightCodes: getHighlightCodes()}
            })
        )
    }, []);

    return (
        <AnimatePresence>
            <motion.div initial={{opacity: 0.6}} animate={{opacity: 1, transition: {duration: 0.25}}}
                        exit={{opacity: 0.6, transition: {duration: 0.25}}}>
                {!!judgeData && !!judgeData.languages && !!judgeData.languages && !!judgeData.tags &&
                    <div className="text-white h-full min-h-screen pb-4">
                        <FlashContainer/>
                        <Router>
                            <div className="pb-20">
                                <Menubar/>
                            </div>
                            <RoutingComponent/>
                        </Router>
                    </div>}
            </motion.div>
        </AnimatePresence>
    );
}

export default App;