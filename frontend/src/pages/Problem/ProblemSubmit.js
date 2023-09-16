import Editor from '@monaco-editor/react';
import RoundedFrame from "../../components/RoundedFrame";
import DropdownMenu from "../../components/DropdownMenu";

function SubmitControlsFrame() {
    return (
        <RoundedFrame>
            <div className="px-4 py-3 sm:px-6 sm:py-5 flex">
                <DropdownMenu itemNames={["C++ 11", "C++ 14", "C++ 17", "Go", "Java", "Python 3"]}/>
                <button className="ml-3 btn-indigo">Beküldés</button>
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmit() {
    return (
        <div className="flex flex-col">
            <div className="mb-2">
                <SubmitControlsFrame/>
            </div>
            <Editor className="border-1 border-default" height="60vh" theme="vs-dark" defaultLanguage="cpp"
                    options={{fontFamily: 'JetBrains Mono'}}/>
        </div>
    )
}

export default ProblemSubmit;