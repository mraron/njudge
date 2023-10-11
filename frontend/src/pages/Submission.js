import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { SVGTitleComponent } from "../components/container/RoundedFrame";
import SubmissionTable from "../components/concrete/table/SubmissionTable";
import SubmissionsTable from "../components/concrete/table/SubmissionsTable";
import CopyableCode from "../components/util/copy/CopyableCode";
import DropdownFrame from "../components/container/DropdownFrame";
import Button from "../components/basic/Button";
import CodeEditor from "../components/input/CodeEditor";
import JudgeDataContext from "../contexts/judgeData/JudgeDataContext";
import UserContext from "../contexts/user/UserContext";
import WidePage from "./wrappers/WidePage";

function CompileErrorFrame({ message }) {
    const { t } = useTranslation()
    const titleComponent = (
        <SVGTitleComponent
            title={t("submission.compilation_error")}
            icon={<FontAwesomeIcon icon="fa-xmark" className="w-4 h-4 mr-3 text-red-600" />}
        />
    )
    return <CopyableCode text={message} titleComponent={titleComponent} maxHeight="16rem" />
}

function Submission({ data }) {
    const { userData } = useContext(UserContext)
    const { judgeData } = useContext(JudgeDataContext)

    return (
        <WidePage>
            <div className="w-full space-y-2">
                {userData && userData.isAdmin && (
                    <DropdownFrame title="Kezelés">
                        <div className="px-4 py-3 sm:px-6 sm:py-5 flex items-center justify-center space-x-2">
                            <Button color="indigo" minWidth="8rem">
                                Újrafordít
                            </Button>
                            <Button color="gray" minWidth="8rem">
                                Újraértékel
                            </Button>
                        </div>
                    </DropdownFrame>
                )}
                <SubmissionsTable submissions={[data.summary]} />
                {data.language !== "zip" && (
                    <CodeEditor
                        value={data.summary.code}
                        language={judgeData.highlightCodes[data.summary.language]}
                    />
                )}
                {data.summary.compileError && <CompileErrorFrame message={data.summary.compileErrorMessage} />}
                {!data.summary.compileError && <SubmissionTable status={data.status} />}
            </div>
        </WidePage>
    )
}

export default Submission
