import { useEffect } from "react";

function Admin() {
    useEffect(() => {
        window.location.href = "/admin"
    }, [])
}

export default Admin
