import {useNavigate} from "react-router-dom";
import {logout} from "../../util/auth";
import {useEffect} from "react";
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

function Logout() {
    const {t} = useTranslation()
    const navigate = useNavigate()

    useEffect(() => {
        if (logout()) {
            window.flash(t("flash.successful_logout"), "success")
        } else {
            window.flash(t("flash.not_logged_in"), "failure")
        }
        navigate(routeMap.main)
    }, [])
    return (
        <></>
    )
}

export default Logout