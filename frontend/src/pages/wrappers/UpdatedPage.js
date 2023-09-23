import { useContext, useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { AnimatePresence, motion } from "framer-motion";
import { updateData } from "../../util/updateData";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import UserContext from "../../contexts/user/UserContext";

function UpdatePage({ page: Page }) {
    const { setUserData, setLoggedIn } = useContext(UserContext);
    const location = useLocation();
    const [data, setData] = useState(null);
    const [isLoading, setLoading] = useState(true);
    const abortController = new AbortController();

    useEffect(() => {
        let isMounted = true;
        updateData(
            location,
            abortController,
            setData,
            setUserData,
            setLoggedIn,
            () => isMounted,
        ).then(() => setLoading(false));
        return () => {
            isMounted = false;
            abortController.abort();
        };
    }, []);

    let passedData = null;
    if (!isLoading) {
        passedData = data;
    }
    return (
        <div className="relative w-full">
            <PageLoadingAnimation isVisible={isLoading} />
            <AnimatePresence>
                {!isLoading && (
                    <motion.div
                        initial={{ opacity: 0.6 }}
                        animate={{ opacity: 1, transition: { duration: 0.25 } }}
                        exit={{ opacity: 0.6, transition: { duration: 0.25 } }}>
                        <Page isLoading={isLoading} data={passedData} />
                    </motion.div>
                )}
            </AnimatePresence>
        </div>
    );
}

export default UpdatePage;
