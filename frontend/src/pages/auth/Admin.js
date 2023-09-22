import {useEffect} from "react";

function Admin() {
    useEffect(() => {
        window.location.href = "/admin"
    }, []);

    return (
        <></>
    )
}

export default Admin