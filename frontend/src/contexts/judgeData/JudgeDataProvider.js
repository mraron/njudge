import { useState } from "react";

import JudgeDataContext from "./JudgeDataContext";

function JudgeDataProvider({ children }) {
    const [judgeData, setJudgeData] = useState(null)

    const allLoaded = () => {
        return judgeData && judgeData.languages && judgeData.tags && judgeData.categories
    }
    return (
        <JudgeDataContext.Provider value={{ judgeData, setJudgeData, allLoaded }}>{children}</JudgeDataContext.Provider>
    )
}

export default JudgeDataProvider
