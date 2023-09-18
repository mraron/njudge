import UserContext from "./UserContext";
import {useState} from "react";

function UserProvider({children}) {
    const [userData, setUserData] = useState(null)
    const [isLoggedIn, setLoggedIn] = useState(null)

    return (
        <UserContext.Provider value={{userData, isLoggedIn, setUserData, setLoggedIn}}>
            {children}
        </UserContext.Provider>
    )
}

export default UserProvider