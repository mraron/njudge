import {Navigate, useNavigate} from "react-router-dom";
import {logout} from "../../util/auth";
import {useEffect} from "react";
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

function Logout() {
    const navigate = useNavigate()

    useEffect(() => {
        if (logout()) {
            window.flash("flash.successful_logout", "success")
        } else {
            window.flash("flash.not_logged_in", "failure")
        }
    }, [])
    return (
        <Navigate to={routeMap.home} />
    )
}

export default Logout