import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import { AnimatePresence, motion } from "framer-motion";
import FlashEvent from "./FlashEvent";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function FlashMessage({ message, type, onClose }) {
    const { t } = useTranslation();
    return (
        <div className="absolute bottom-0 left-0 right-0 bg-grey-850 border rounded-md flex border-bordefcol w-full">
            <div className="w-full p-6 flex justify-between items-center rounded-md">
                <div className="flex items-center mr-2">
                    {type === "success" && (
                        <FontAwesomeIcon
                            icon="fa-regular fa-circle-check"
                            className="w-7 h-7 highlight-green mr-3"
                        />
                    )}
                    {type === "failure" && (
                        <FontAwesomeIcon
                            icon="fa-regular fa-circle-xmark"
                            className="w-7 h-7 highlight-red mr-3"
                        />
                    )}
                    <span>{t(message)}</span>
                </div>
                <button
                    className="flex rounded-full p-3 hover:bg-framebgcol"
                    onClick={onClose}>
                    <FontAwesomeIcon icon="fa-close" className="w-5 h-5" />
                </button>
            </div>
        </div>
    );
}

function FlashContainer() {
    const [messages, setMessages] = useState([]);

    const pushMessage = (message, type) => {
        setMessages((prevMessages) => [
            ...prevMessages,
            <motion.div
                key={prevMessages.length}
                initial={{ scaleX: 0.36, opacity: 0.01 }}
                animate={{
                    scaleX: 1,
                    opacity: 1,
                    transition: { duration: 0.16, ease: "easeOut" },
                }}
                exit={{
                    scaleX: 0.36,
                    opacity: 0.01,
                    transition: { duration: 0.16, ease: "easeIn" },
                }}>
                <FlashMessage
                    message={message}
                    type={type}
                    onClose={popMessage}
                />
            </motion.div>,
        ]);
    };
    const popMessage = () => {
        setMessages((prevMessages) => prevMessages.slice(0, -1));
    };
    useEffect(() => {
        FlashEvent.addListener("flash", ({ message, type }) => {
            pushMessage(message, type);
        });
        return () => {
            FlashEvent.removeAllListeners();
        };
    }, []);

    return (
        <div className="z-10 fixed bottom-2 left-2 right-2 flex justify-center">
            <div className={`relative w-full max-w-7xl`}>
                <AnimatePresence>{messages}</AnimatePresence>
            </div>
        </div>
    );
}

export default FlashContainer;
