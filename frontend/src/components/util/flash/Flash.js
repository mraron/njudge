import {SVGClose, SVGCorrect, SVGInformation, SVGWrong} from "../../../svg/SVGs";
import {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {AnimatePresence, motion} from "framer-motion";
import FlashEvent from "./FlashEvent";

function FlashMessage({message, type, onClose}) {
    const {t} = useTranslation()
    return (
        <div className="absolute bottom-0 left-0 right-0 bg-grey-850 border-1 rounded-md flex border-default w-full">
            <div className="w-full p-6 flex justify-between">
                <div className="flex items-center">
                    {type === "success" && <SVGCorrect cls="w-7 h-7 text-green-500 mr-3"/>}
                    {type === "failure" && <SVGWrong cls="w-7 h-7 text-red-500 mr-3"/>}
                    {type === "info" && <SVGInformation cls="w-7 h-7 text-indigo-500 mr-3"/>}
                    <span>{t(message)}</span>
                </div>
                <button className="rounded-full p-3 hover:bg-grey-800 transition duration-200" onClick={onClose}>
                    <SVGClose cls="w-4 h-4 text-white"/>
                </button>
            </div>
        </div>
    )
}

function FlashContainer() {
    const [messages, setMessages] = useState([])

    const pushMessage = (message, type) => {
        setMessages(prevMessages => [...prevMessages,
            <motion.div key={prevMessages.length} initial={{scaleX: 0.36, opacity: 0.01}}
                        animate={{scaleX: 1, opacity: 1, transition: {duration: 0.16, ease: "linear"}}}
                        exit={{scaleX: 0.36, opacity: 0.01, transition: {duration: 0.16, ease: "linear"}}}>
                <FlashMessage message={message} type={type} onClose={popMessage}/>
            </motion.div>
        ])
    }
    const popMessage = () => {
        setMessages(prevMessages => prevMessages.slice(0, -1))
    }
    useEffect(() => {
        FlashEvent.addListener('flash', ({message, type}) => {
            pushMessage(message, type)
        })
        return () => {
            FlashEvent.removeAllListeners()
        }
    }, [])
    return (
        <>
            <div className="z-10 fixed bottom-2 left-2 right-2 flex justify-center">
                <div className={`relative w-full max-w-7xl px-2`}>
                    <AnimatePresence>
                        {messages}
                    </AnimatePresence>
                </div>
            </div>
        </>
    )
}

export default FlashContainer