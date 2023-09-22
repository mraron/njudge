import SubmissionTable from "../components/concrete/table/SubmissionTable";
import SubmissionsTable from "../components/concrete/table/SubmissionsTable";
import {useTranslation} from "react-i18next";
import SVGTitleComponent from "../svg/SVGTitleComponent";
import CopyableCode from "../components/util/copy/CopyableCode";
import {SVGWrongSimple} from "../svg/SVGs";
import Editor from "@monaco-editor/react";


function CompileErrorFrame({message}) {
    const {t} = useTranslation()
    const titleComponent =
        <SVGTitleComponent title={t("submission.compilation_error")} svg={<SVGWrongSimple cls="w-6 h-6 mr-2 text-red-500" />} />

    return (
        <CopyableCode text={message} titleComponent={titleComponent} maxHeight="16rem"/>
    )
}

function Submission({data}) {
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <div className="mb-3">
                            <SubmissionsTable submissions={[data.summary]}/>
                        </div>
                        {data.language !== "zip" && <div className="mb-3">
                            <Editor className="border-1 border-default" height="40vh" theme="vs-dark"
                                    defaultLanguage="cpp"
                                    options={{domReadOnly: true, readOnly: true, fontFamily: 'JetBrains Mono'}}
                                    value={data.summary.code}/>
                        </div>}
                        {data.summary.compileError && <CompileErrorFrame message={data.summary.compileErrorMessage} />}
                        {!data.summary.compileError && <SubmissionTable status={data.status}/>}
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Submission;