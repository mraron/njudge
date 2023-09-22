import {Navigate, useNavigate} from "react-router-dom";
import {logout} from "../../util/auth";
import {useEffect} from "react";
import {routeMap} from "../../config/RouteConfig";
import {useTranslation} from "react-i18next";

function Logout() {
    const navigate = useNavigate()
    logout().then((result) => {
        if (result) {
            window.flash("flash.successful_logout", "success")
            navigate(routeMap.home)
        } else {
            window.flash("flash.not_logged_in", "failure")
        }
    })
}

export default Logout