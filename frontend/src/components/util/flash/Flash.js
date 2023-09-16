import {SVGClose, SVGCorrect, SVGInformation, SVGWrong} from "../../../svg/SVGs";
import {useEffect, useState} from "react";
import FlashEvent from "./FlashEvent";
import {AnimatePresence, motion} from "framer-motion";

function FlashMessage({ message, type, onClose }) {
    return (
        <div className={`absolute bottom-0 left-0 right-0 bg-grey-850 border-1 rounded-md flex border-default w-full`}>
            <div className="w-full p-6 flex justify-between">
                <div className="flex items-center">
                    {type === "success" && <SVGCorrect cls="w-7 h-7 text-green-500 mr-3" />}
                    {type === "failure" && <SVGWrong cls="w-7 h-7 text-red-500 mr-3" />}
                    {type === "info" && <SVGInformation cls="w-7 h-7 text-indigo-500 mr-3" />}
                    <span>{message}</span>
                </div>
                <button className="rounded-full p-3 hover:bg-grey-800 transition duration-200" onClick={onClose}>
                    <SVGClose cls="w-4 h-4 text-white" />
                </button>
            </div>
        </div>
    )
}

function FlashContainer() {
    const [messages, setMessages] = useState([])

    const pushMessage = (message, type) => {
        setMessages(prevMessages => [...prevMessages,
            <motion.div key={prevMessages.length} initial={{scaleX: 0.2, opacity: 0.2}} animate={{scaleX: 1, opacity: 1, transition: {duration: 0.2}}}
                        exit={{scaleX: 0.2, opacity: 0.2, transition: {duration: 0.2}}}>
                <FlashMessage message={message} type={type} onClose={popMessage} />
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
    })
    return (
        <>
            <div className="fixed bottom-2 left-0 right-0 flex justify-center">
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