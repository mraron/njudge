import { useContext } from "react"
import { useTranslation } from "react-i18next"
import SubmissionTable from "../components/concrete/table/SubmissionTable"
import SubmissionsTable from "../components/concrete/table/SubmissionsTable"
import SVGTitleComponent from "../components/svg/SVGTitleComponent"
import CopyableCode from "../components/util/copy/CopyableCode"
import DropdownFrame from "../components/container/DropdownFrame"

import JudgeDataContext from "../contexts/judgeData/JudgeDataContext"
import UserContext from "../contexts/user/UserContext"
import Editor from "@monaco-editor/react"
import ThemeContext from "../contexts/theme/ThemeContext"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import Button from "../components/util/Button"

function CompileErrorFrame({ message }) {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent
            title={t("submission.compilation_error")}
            svg={
                <FontAwesomeIcon
                    icon="fa-xmark"
                    className="w-4 h-4 mr-3 text-red-600"
                />
            }
        />
    )
    return (
        <CopyableCode
            text={message}
            titleComponent={titleComponent}
            maxHeight="16rem"
        />
    )
}

function Submission({ data }) {
    const { userData } = useContext(UserContext)
    const { judgeData } = useContext(JudgeDataContext)
    const { theme } = useContext(ThemeContext)

    return (
        <div className="w-full flex justify-center">
            <div className="flex justify-center w-full max-w-7xl">
                <div className="w-full px-4 space-y-3">
                    {userData && userData.isAdmin && (
                        <DropdownFrame title="Kezelés">
                            <div className="px-4 py-3 sm:px-6 sm:py-5 flex items-center justify-center space-x-2">
                                <Button color="indigo" minWidth="8rem">
                                    Újfafordít
                                </Button>
                                <Button color="gray" minWidth="8rem">
                                    Újraértékel
                                </Button>
                            </div>
                        </DropdownFrame>
                    )}
                    <SubmissionsTable submissions={[data.summary]} />
                    {data.language !== "zip" && (
                        <Editor
                            className="editor"
                            height="40vh"
                            theme={`${theme === "light" ? "vs" : "vs-dark"}`}
                            options={{
                                domReadOnly: true,
                                readOnly: true,
                                fontFamily: "JetBrains Mono",
                                fontSize: 13,
                            }}
                            value={data.summary.code}
                            language={
                                judgeData.highlightCodes[data.summary.language]
                            }
                        />
                    )}
                    {data.summary.compileError && (
                        <CompileErrorFrame
                            message={data.summary.compileErrorMessage}
                        />
                    )}
                    {!data.summary.compileError && (
                        <SubmissionTable status={data.status} />
                    )}
                </div>
            </div>
        </div>
    )
}

export default Submission
