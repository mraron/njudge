import { useContext, useEffect } from "react"
import { BrowserRouter as Router } from "react-router-dom"
import { useMonaco } from "@monaco-editor/react"
import { AnimatePresence, motion } from "framer-motion"

import RoutingComponent from "./RoutingComponent"
import FlashContainer from "./components/util/flash/Flash"
import Menubar from "./components/concrete/other/Menubar"
import FlashEvent from "./components/util/flash/FlashEvent"

import { getCategories, getHighlightCodes, getLanguages, getTags } from "./util/getJudgeData"

import JudgeDataContext from "./contexts/judgeData/JudgeDataContext"
import ThemeContext from "./contexts/theme/ThemeContext"

window.flash = (message, type = "success") => FlashEvent.emit("flash", { message, type })

function App() {
    const { setJudgeData, allLoaded } = useContext(JudgeDataContext)
    const { theme } = useContext(ThemeContext)
    const monaco = useMonaco()

    useEffect(() => {
        if (theme === "light") {
            monaco?.editor.defineTheme("standard-theme", {
                base: "vs",
                inherit: true,
                rules: [],
                colors: {
                    "editor.background": "#FDF9FF",
                },
            })
        } else {
            monaco?.editor.defineTheme("standard-theme", {
                base: "vs-dark",
                inherit: true,
                rules: [],
                colors: {
                    "editor.background": "#0c080f",
                },
            })
        }
    }, [monaco, theme])

    useEffect(() => {
        const fetchWithCredentialsJudgeData = async () => {
            await getLanguages().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return { ...prevJudgeData, languages: resp.languages }
                    })
                }
            })
            await getCategories().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return {
                            ...prevJudgeData,
                            categories: resp.categories,
                        }
                    })
                }
            })
            await getTags().then((resp) => {
                if (resp.success) {
                    setJudgeData((prevJudgeData) => {
                        return { ...prevJudgeData, tags: resp.tags }
                    })
                }
            })
        }
        fetchWithCredentialsJudgeData().then(
            setJudgeData((prevJudgeData) => {
                return {
                    ...prevJudgeData,
                    highlightCodes: getHighlightCodes(),
                }
            }),
        )
    }, [])

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
    )
}

export default App
