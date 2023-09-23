import UserContext from "./JudgeDataContext";
import {useState} from "react";

function JudgeDataProvider({children}) {
    const [judgeData, setJudgeData] = useState(null)

    return (
        <UserContext.Provider value={{judgeData, setJudgeData}}>
            {children}
        </UserContext.Provider>
    )
}

export default JudgeDataProvider