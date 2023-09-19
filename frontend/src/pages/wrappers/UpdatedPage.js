import React, {useContext, useEffect, useState} from "react";
import {useLocation} from 'react-router-dom';
import {updateData} from "../../util/updateData";
import PageLoadingAnimation from "../../components/util/PageLoadingAnimation";
import FadeIn from "../../components/util/FadeIn";
import UserContext from "../../contexts/user/UserContext";

function UpdatePage({ page: Page }) {
    const {setUserData, setLoggedIn} = useContext(UserContext)
    const location = useLocation()
    const [data, setData] = useState(null)
    const [isLoading, setLoading] = useState(true)
    const abortController = new AbortController();

    useEffect(() => {
        let isMounted = true
        updateData(location, abortController, setData, setUserData, setLoggedIn, () => isMounted).then(() =>
            setLoading(false)
        )
        return () => {
            isMounted = false
            abortController.abort()
        }
    }, []);

    let passedData = null
    if (!isLoading) {
        passedData = data
    }
    return (
        <div className="relative w-full">
            <PageLoadingAnimation isVisible={isLoading} />
            {!isLoading &&
                <FadeIn><Page isLoading={isLoading} data={passedData} /></FadeIn>}
        </div>
    );
}

export default UpdatePage;