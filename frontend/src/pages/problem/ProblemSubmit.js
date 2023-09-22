import MonacoEditor from '@monaco-editor/react';
import RoundedFrame from "../../components/container/RoundedFrame";
import DropdownMenu from "../../components/input/DropdownMenu";
import {useTranslation} from "react-i18next";
import {useState} from "react";
import submitSolution from "../../util/submitSolution";
import {routeMap} from "../../config/RouteConfig";
import {useNavigate, useParams} from "react-router-dom";

const languages = ["cpp11", "cpp14", "cpp17", "go", "java", "python3"]
const langCodes = ["cpp", "cpp", "cpp", "go", "java", "python"]

function SubmitControlsFrame({onLanguageChanged, onSubmit}) {
    const {t} = useTranslation()
    return (
        <RoundedFrame>
            <div className="px-4 py-3 sm:px-6 sm:py-5 flex">
                <DropdownMenu itemNames={["C++ 11", "C++ 14", "C++ 17", "Go", "Java", "Python 3"]}
                              onChange={onLanguageChanged}/>
                <button className="ml-3 btn-indigo padding-btn-default" onClick={onSubmit}>
                    {t("problem_submit.submit")}
                </button>
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmit({data}) {
    const {problem} = useParams()
    const [langIndex, setLangIndex] = useState(0)
    const [submissionCode, setSubmissionCode] = useState("")
    const navigate = useNavigate()
    const handleLanguageChanged = (index) => {
        setLangIndex(index)
    }
    const handleSubmit = () => {
        submitSolution({problem: problem, language: languages[langIndex], submissionCode: submissionCode}).then(ok => {
            if (ok) {
                window.flash("flash.successful_submission", "success")
                navigate(routeMap.problemSubmissions.replace(":problem", problem))
            } else {
                window.flash("flash.unsuccessful_submission", "failure")
            }
        })
    }
    return (
        <div className="flex flex-col">
            <div className="mb-2">
                <SubmitControlsFrame onSubmit={handleSubmit} onLanguageChanged={handleLanguageChanged}/>
            </div>
            <MonacoEditor className="border-1 border-default" height="60vh" theme="vs-dark"
                          language={langCodes[langIndex]}
                          options={{fontFamily: 'JetBrains Mono'}} onChange={setSubmissionCode}/>
        </div>
    )
}

export default ProblemSubmit;