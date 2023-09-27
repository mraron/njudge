import { useEffect } from "react";
import { Navigate, useNavigate, useParams } from "react-router-dom";
import { verify } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";

function Verify() {
    const { token } = useParams();

    useEffect(() => {
        verify(token).then((ok) => {
            if (ok) {
                window.flash("flash.successful_verification", "success");
            } else {
                window.flash("flash.unsuccessful_verification", "failure");
            }
        });
    }, []);
    return <Navigate to={routeMap.home} />;
}

export default Verify;
