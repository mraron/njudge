import SubmissionTable from "../components/SubmissionTable";
import SubmissionsTable from "../components/SubmissionsTable";
import Editor from "@monaco-editor/react";
import checkData from "../util/CheckData";
import {useOutletContext} from "react-router-dom";

function Submission({ data }) {
    if (!checkData(data)) {
        return
    }
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <div className="mb-3">
                            <SubmissionsTable submissions={[data.summary]}/>
                        </div>
                        <div className="mb-3">
                            <Editor className="border-1 border-default" height="60vh" theme="vs-dark"
                                    defaultLanguage="cpp"
                                    options={{domReadOnly: true, readOnly: true, fontFamily: 'JetBrains Mono'}}
                                    value={data.summary.code}/>
                        </div>
                        <SubmissionTable status={data.status}/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Submission;