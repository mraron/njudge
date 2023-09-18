import {useNavigate} from "react-router-dom";
import {logout} from "../../util/User";
import {useEffect} from "react";

function Logout() {
    const navigate = useNavigate()

    useEffect(() => {
        if (logout()) {
            window.flash("Sikeres kilépés!", "success")
        } else {
            window.flash("Nem vagy belépve.", "failure")
        }
        navigate("/")
    }, [])
    return (
        <></>
    )
}

export default Logout