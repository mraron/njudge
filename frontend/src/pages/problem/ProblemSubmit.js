import { useContext, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useTranslation } from "react-i18next";
import MonacoEditor from "@monaco-editor/react";
import RoundedFrame from "../../components/container/RoundedFrame";
import DropdownMenu from "../../components/input/DropdownMenu";
import { routeMap } from "../../config/RouteConfig";
import submitSolution from "../../util/submitSolution";
import JudgeDataContext from "../../contexts/judgeData/JudgeDataContext";
import ThemeContext from "../../contexts/theme/ThemeContext";
import Button from "../../components/util/Button";

function SubmitControlsFrame({ onLanguageChanged, onSubmit }) {
    const { t } = useTranslation();
    const { judgeData } = useContext(JudgeDataContext);

    return (
        <RoundedFrame>
            <div className="px-4 py-3 sm:px-6 sm:py-5 flex items-stretch">
                <div className="w-full min-w-0 mr-3">
                    <DropdownMenu
                        itemNames={judgeData.languages.map(
                            (item) => item.label,
                        )}
                        onChange={onLanguageChanged}
                    />
                </div>
                <Button color="indigo" minWidth="8rem" onClick={onSubmit}>
                    {t("problem_submit.submit")}
                </Button>
            </div>
        </RoundedFrame>
    );
}

function ProblemSubmit() {
    const { judgeData } = useContext(JudgeDataContext);
    const { theme } = useContext(ThemeContext);
    const { problem, problemset } = useParams();
    const [langIndex, setLangIndex] = useState(0);
    const [submissionCode, setSubmissionCode] = useState("");
    const navigate = useNavigate();
    const handleLanguageChanged = (index) => {
        setLangIndex(index);
    };
    const handleSubmit = () => {
        submitSolution({
            problemset: problemset,
            problem: problem,
            language: judgeData.languages[langIndex].id,
            submissionCode: submissionCode,
        }).then((ok) => {
            if (ok) {
                window.flash("flash.successful_submission", "success");
                navigate(
                    routeMap.problemSubmissions
                        .replace(":problemset", problemset)
                        .replace(":problem", problem),
                );
            } else {
                window.flash("flash.unsuccessful_submission", "failure");
            }
        });
    };
    return (
        <div className="flex flex-col">
            <div className="mb-2">
                <SubmitControlsFrame
                    onSubmit={handleSubmit}
                    onLanguageChanged={handleLanguageChanged}
                />
            </div>
            <MonacoEditor
                className="editor"
                height="60vh"
                theme={`${theme === "light" ? "vs" : "vs-dark"}`}
                language={
                    judgeData.highlightCodes[judgeData.languages[langIndex].id]
                }
                options={{ fontFamily: "JetBrains Mono", fontSize: 13 }}
                onChange={setSubmissionCode}
            />
        </div>
    );
}

export default ProblemSubmit;
