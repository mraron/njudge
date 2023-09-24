import { useState } from "react";

import JudgeDataContext from "./JudgeDataContext";

function JudgeDataProvider({ children }) {
    const [judgeData, setJudgeData] = useState(null);

    return (
        <JudgeDataContext.Provider value={{ judgeData, setJudgeData }}>
            {children}
        </JudgeDataContext.Provider>
    );
}

export default JudgeDataProvider;
