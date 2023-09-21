import SubmissionTable from "../components/concrete/table/SubmissionTable";
import SubmissionsTable from "../components/concrete/table/SubmissionsTable";
import Editor from "@monaco-editor/react";

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
                            <Editor className="border-1 border-default" height="60vh" theme="vs-dark"
                                    defaultLanguage="cpp"
                                    options={{domReadOnly: true, readOnly: true, fontFamily: 'JetBrains Mono'}}
                                    value={data.summary.code}/>
                        </div>}
                        <SubmissionTable status={data.status}/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Submission;