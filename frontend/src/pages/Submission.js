import SubmissionTable from "../components/concrete/table/SubmissionTable";
import SubmissionsTable from "../components/concrete/table/SubmissionsTable";
import {useTranslation} from "react-i18next";
import SVGTitleComponent from "../svg/SVGTitleComponent";
import CopyableCode from "../components/util/copy/CopyableCode";
import {SVGWrongSimple} from "../svg/SVGs";
import Editor from "@monaco-editor/react";
import {useContext} from "react";
import UserContext from "../contexts/user/UserContext";
import DropdownFrame from "../components/container/DropdownFrame";


function CompileErrorFrame({message}) {
    const {t} = useTranslation()
    const titleComponent =
        <SVGTitleComponent title={t("submission.compilation_error")}
                           svg={<SVGWrongSimple cls="w-6 h-6 mr-2 text-red-500"/>}/>

    return (
        <CopyableCode text={message} titleComponent={titleComponent} maxHeight="16rem"/>
    )
}

function Submission({data}) {
    const {userData} = useContext(UserContext)
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        {userData && userData.isAdmin &&
                            <div className="mb-3">
                                <DropdownFrame title="Kezelés">
                                    <div className="px-4 py-3 sm:px-6 sm:py-5 flex items-center justify-center">
                                        <button className="w-32 btn-indigo padding-btn-default mr-2">Újrafordít</button>
                                        <button className="w-32 btn-gray padding-btn-default">Újraértékel</button>
                                    </div>
                                </DropdownFrame>
                            </div>}
                        <div className="mb-3">
                            <SubmissionsTable submissions={[data.summary]}/>
                        </div>
                        {data.language !== "zip" &&
                            <div className="mb-3">
                                <Editor className="border-1 border-default" height="40vh" theme="vs-dark"
                                        defaultLanguage="cpp"
                                        options={{domReadOnly: true, readOnly: true, fontFamily: 'JetBrains Mono'}}
                                        value={data.summary.code}/>
                            </div>}
                        {data.summary.compileError && <CompileErrorFrame message={data.summary.compileErrorMessage}/>}
                        {!data.summary.compileError && <SubmissionTable status={data.status}/>}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Submission;