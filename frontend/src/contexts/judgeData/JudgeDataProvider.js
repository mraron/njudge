import { useState } from "react";

import UserContext from "./JudgeDataContext";

function JudgeDataProvider({ children }) {
    const [judgeData, setJudgeData] = useState(null);

    return (
        <UserContext.Provider value={{ judgeData, setJudgeData }}>
            {children}
        </UserContext.Provider>
    );
}

export default JudgeDataProvider;
