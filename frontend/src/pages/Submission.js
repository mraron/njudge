import SubmissionTable from "../components/SubmissionTable";
import SubmissionsTable from "../components/SubmissionsTable";
import Editor from "@monaco-editor/react";

function Submission() {
    const submission = {
        "compiled": true,
        "feedbackType": 1,
        "testSets": [
            {
                "groups": [
                    {
                        "name": "subtask1",
                        "completed": true,
                        "failed": true,
                        "score": 0,
                        "maxScore": 9,
                        "scoring": 1,
                        "testCases": [
                            {
                                "index": 1,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 2,
                            },
                            {
                                "index": 2,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 3,
                            },
                            {
                                "index": 3,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 4,
                            },
                            {
                                "index": 4,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 2,
                            },
                            {
                                "index": 5,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 3,
                            },
                            {
                                "index": 6,
                                "verdictName": "Időlimit túllépés",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 4,
                            },
                        ]
                    },
                    {
                        "name": "subtask2",
                        "completed": false,
                        "failed": false,
                        "score": 0,
                        "maxScore": 9,
                        "scoring": 2,
                        "testCases": [
                            {
                                "index": 7,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 2,
                                "maxScore": 2,
                            },
                            {
                                "index": 8,
                                "verdictName": "Futtatás...",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 3,
                            },
                            {
                                "index": 9,
                                "verdictName": "Futtatás...",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 4,
                            },
                            {
                                "index": 10,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 2,
                                "maxScore": 2,
                            },
                            {
                                "index": 11,
                                "verdictName": "Futtatás...",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 3,
                            },
                            {
                                "index": 12,
                                "verdictName": "Futtatás...",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 0,
                                "maxScore": 4,
                            },
                        ]
                    },
                    {
                        "name": "subtask3",
                        "completed": true,
                        "failed": false,
                        "score": 9,
                        "maxScore": 9,
                        "scoring": 2,
                        "testCases": [
                            {
                                "index": 13,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 2,
                                "maxScore": 2,
                            },
                            {
                                "index": 14,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 3,
                                "maxScore": 3,
                            },
                            {
                                "index": 15,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 4,
                                "maxScore": 4,
                            },
                            {
                                "index": 16,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 2,
                                "maxScore": 2,
                            },
                            {
                                "index": 17,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 3,
                                "maxScore": 3,
                            },
                            {
                                "index": 18,
                                "verdictName": "Elfogadva",
                                "timeSpent": "150 ms",
                                "memoryUsed": "15127 KiB",
                                "score": 4,
                                "maxScore": 4,
                            },
                        ]
                    },
                ]
            }
        ]
    }
    return (
        <div className="text-white">
            <div className="w-full flex justify-center">
                <div className="flex justify-center w-full max-w-7xl">
                    <div className="w-full px-4">
                        <div className="mb-3">
                            <SubmissionsTable submissions={[]}/>
                        </div>
                        <div className="mb-3">
                            <Editor className="border-1 border-default" height="60vh" theme="vs-dark"
                                    defaultLanguage="cpp"
                                    options={{domReadOnly: true, readOnly: true, fontFamily: 'JetBrains Mono'}}
                                    value={`#include <iostream>
using namespace std;

int main() {
    cout << "Hello world" << endl;
}`}/>
                        </div>
                        <SubmissionTable submission={submission}/>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default Submission;