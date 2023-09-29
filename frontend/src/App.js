import { useContext, useEffect } from "react";
import { BrowserRouter as Router } from "react-router-dom";
import { AnimatePresence, motion } from "framer-motion";

import RoutingComponent from "./RoutingComponent";
import FlashContainer from "./components/util/flash/Flash";
import Menubar from "./components/concrete/other/Menubar";
import FlashEvent from "./components/util/flash/FlashEvent";

import {
    getCategories,
    getHighlightCodes,
    getLanguages,
    getTags,
} from "./util/getJudgeData";

import JudgeDataContext from "./contexts/judgeData/JudgeDataContext";

window.flash = (message, type = "success") =>
    FlashEvent.emit("flash", { message, type });

function App() {
    const { setJudgeData, allLoaded } = useContext(JudgeDataContext);

    useEffect(() => {
        const fetchWithCredentialsJudgeData = async () => {
            await getLanguages().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return { ...prevJudgeData, languages: resp.languages };
                    });
                }
            });
            await getCategories().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return {
                            ...prevJudgeData,
                            categories: resp.categories,
                        };
                    });
                }
            });
            await getTags().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return { ...prevJudgeData, tags: resp.tags };
                    });
                }
            });
        };
        fetchWithCredentialsJudgeData().then(
            setJudgeData((prevJudgeData) => {
                return {
                    ...prevJudgeData,
                    highlightCodes: getHighlightCodes(),
                };
            }),
        );
    }, []);

    return (
        <AnimatePresence>
            <motion.div
                initial={{ opacity: 0.6 }}
                animate={{ opacity: 1, transition: { duration: 0.25 } }}
                exit={{ opacity: 0.6, transition: { duration: 0.25 } }}>
                {allLoaded() && (
                    <div className="relative h-full min-h-screen pb-4">
                        <FlashContainer />
                        <Router>
                            <div className="pb-20">
                                <Menubar />
                            </div>
                            <div>
                                <RoutingComponent />
                            </div>
                        </Router>
                    </div>
                )}
            </motion.div>
        </AnimatePresence>
    );
}

export default App;
