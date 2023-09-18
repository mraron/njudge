import {useNavigate} from "react-router-dom";
import {logout} from "../../util/Auth";
import {useEffect} from "react";
import {routeMap} from "../../config/RouteConfig";

function Logout() {
    const navigate = useNavigate()

    useEffect(() => {
        if (logout()) {
            window.flash("Sikeres kilépés!", "success")
        } else {
            window.flash("Nem vagy belépve.", "failure")
        }
        navigate(routeMap.main)
    }, [])
    return (
        <></>
    )
}

export default Logout