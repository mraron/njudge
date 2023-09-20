import {Navigate, useNavigate, useParams} from "react-router-dom";
import {useEffect} from "react";
import {routeMap} from "../../config/RouteConfig";
import {verify} from "../../util/auth";

function Verify() {
    const {token} = useParams()
    const navigate = useNavigate()

    useEffect(() => {
        verify(token).then(ok => {
            if (ok) {
                window.flash("flash.successful_verification", "success")
            } else {
                window.flash("flash.unsuccessful_verification", "failure")
            }
            navigate(routeMap.home)
        })
    }, [])
    return (
        <Navigate to={routeMap.home} />
    )
}

export default Verify