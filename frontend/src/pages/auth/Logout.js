import { useNavigate } from "react-router-dom";
import { logout } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";
import { useEffect } from "react";

function Logout() {
    const navigate = useNavigate()
    useEffect(() => {
        logout().then((result) => {
            if (result) {
                window.flash("flash.successful_logout", "success")
            } else {
                window.flash("flash.not_logged_in", "failure")
            }
            navigate(routeMap.home)
        })
    }, [])
}

export default Logout
