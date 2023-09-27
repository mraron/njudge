import { Navigate, useNavigate } from "react-router-dom";
import { logout } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";

function Logout() {
    const navigate = useNavigate();
    logout().then((result) => {
        if (result) {
            window.flash("flash.successful_logout", "success");
        } else {
            window.flash("flash.not_logged_in", "failure");
        }
    });
    return <Navigate to={routeMap.home} />;
}

export default Logout;
