import MonacoEditor from '@monaco-editor/react';
import RoundedFrame from "../../components/container/RoundedFrame";
import DropdownMenu from "../../components/input/DropdownMenu";
import {useTranslation} from "react-i18next";
import {useContext, useState} from "react";
import submitSolution from "../../util/submitSolution";
import {routeMap} from "../../config/RouteConfig";
import {useNavigate, useParams} from "react-router-dom";
import JudgeDataContext from "../../contexts/judgeData/JudgeDataContext";

function SubmitControlsFrame({onLanguageChanged, onSubmit}) {
    const {t} = useTranslation()
    const {judgeData} = useContext(JudgeDataContext)

    return (
        <RoundedFrame>
            <div className="px-4 py-3 sm:px-6 sm:py-5 flex">
                <DropdownMenu itemNames={judgeData.languages.map(item => item.label)} onChange={onLanguageChanged}/>
                <button className="ml-3 btn-indigo padding-btn-default w-32" onClick={onSubmit}>
                    {t("problem_submit.submit")}
                </button>
            </div>
        </RoundedFrame>
    )
}

function ProblemSubmit() {
    const {judgeData} = useContext(JudgeDataContext)
    const {problem} = useParams()
    const [langIndex, setLangIndex] = useState(0)
    const [submissionCode, setSubmissionCode] = useState("")
    const navigate = useNavigate()
    const handleLanguageChanged = (index) => {
        setLangIndex(index)
    }
    const handleSubmit = () => {
        submitSolution({
            problem: problem,
            language: judgeData.languages[langIndex].id,
            submissionCode: submissionCode
        }).then(ok => {
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
                          language={judgeData.highlightCodes[judgeData.languages[langIndex].id]}
                          options={{fontFamily: 'JetBrains Mono'}} onChange={setSubmissionCode}/>
        </div>
    )
}

export default ProblemSubmit;