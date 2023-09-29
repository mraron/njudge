import { useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { verify } from "../../util/auth";
import { routeMap } from "../../config/RouteConfig";

function Verify() {
    const { token } = useParams();
    const navigate = useNavigate()

    useEffect(() => {
        verify(token).then((ok) => {
            if (ok) {
                window.flash("flash.successful_verification", "success");
            } else {
                window.flash("flash.unsuccessful_verification", "failure");
            }
            navigate(routeMap.home)
        });
    }, []);
}

export default Verify;
