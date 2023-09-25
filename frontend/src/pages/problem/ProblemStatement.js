import { useContext, useState } from "react";
import { useTranslation } from "react-i18next";
import { Link, useNavigate, useParams } from "react-router-dom";
import MapDataFrame from "../../components/container/MapDataFrame";
import DropdownMenu from "../../components/input/DropdownMenu";
import RoundedFrame from "../../components/container/RoundedFrame";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import {
    SVGAttachment,
    SVGAttachmentDescription,
    SVGAttachmentFile,
    SVGCorrectSimple,
    SVGInformation,
    SVGPartiallyCorrect,
    SVGRecent,
    SVGSpinner,
    SVGSubmit,
    SVGView,
    SVGWrongSimple,
} from "../../components/svg/SVGs";
import RoundedTable from "../../components/container/RoundedTable";
import JudgeDataContext from "../../contexts/judgeData/JudgeDataContext";
import submitSolution from "../../util/submitSolution";
import { routeMap } from "../../config/RouteConfig";
import themeContext from "../../contexts/theme/ThemeContext";
import UserContext from "../../contexts/user/UserContext";
import ThemeContext from "../../contexts/theme/ThemeContext";

function ProblemInfo({ info }) {
    const { t } = useTranslation();
    const tagsContent = (
        <div className="flex flex-wrap">
            {info.tags.map((tagName, index) => (
                <span className="tag" key={index}>
                    {t(tagName)}
                </span>
            ))}
        </div>
    );

    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGInformation cls="w-6 h-6 mr-2" />}
            title={t("problem_statement.information")}
        />
    );
    return (
        <MapDataFrame
            titleComponent={titleComponent}
            data={[
                [t("problem_statement.id"), info.id],
                [t("problem_statement.title"), info.title],
                [t("problem_statement.time_limit"), `${info.timeLimit} ms`],
                [
                    t("problem_statement.memory_limit"),
                    `${info.memoryLimit} MiB`,
                ],
                [t("problem_statement.tags"), tagsContent],
                [t("problem_statement.type"), info.type],
            ]}
            labelColWidth="9rem"
        />
    );
}

function ProblemSubmit() {
    const { t } = useTranslation();
    const { judgeData } = useContext(JudgeDataContext);
    const { problem } = useParams();
    const [file, setFile] = useState(null);
    const [langIndex, setLangIndex] = useState(0);
    const navigate = useNavigate();
    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGSubmit />}
            title={t("problem_statement.submit_solution")}
        />
    );
    const handleFileUploaded = (event) => {
        setFile(event.target.files[0]);
    };
    const handleSubmit = () => {
        if (!file) {
            window.flash("flash.must_choose_file", "failure");
            return;
        }
        submitSolution({
            problem: problem,
            language: judgeData.languages[langIndex].id,
            file: file,
        }).then((ok) => {
            if (ok) {
                window.flash("flash.successful_submission", "success");
                navigate(
                    routeMap.problemSubmissions.replace(":problem", problem),
                );
            } else {
                window.flash("flash.unsuccessful_submission", "failure");
            }
        });
    };
    const handleLanguageChanged = (index) => {
        setLangIndex(index);
    };
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-6 py-5">
                <div className="flex flex-col">
                    <div className="mb-4">
                        <DropdownMenu
                            itemNames={judgeData.languages.map(
                                (item) => item.label,
                            )}
                            onChange={handleLanguageChanged}
                        />
                    </div>
                    <span className="mb-2 mx-1 text-label break-words">
                        {file
                            ? file.name
                            : t("problem_statement.no_file_selected")}
                    </span>
                    <div className="flex justify-center">
                        <button
                            className="btn-gray padding-btn-default w-1/2"
                            onClick={() =>
                                document.getElementById("uploadFile").click()
                            }>
                            <span>{t("problem_statement.choose")}</span>
                            <input
                                id="uploadFile"
                                className="hidden"
                                type="file"
                                onChange={handleFileUploaded}
                            />
                        </button>
                        <button
                            className="ml-2 btn-indigo padding-btn-default w-1/2"
                            onClick={handleSubmit}>
                            {t("problem_statement.submit")}
                        </button>
                    </div>
                </div>
            </div>
        </RoundedFrame>
    );
}

function ProblemLastSubmissions({ submissions, maxScore }) {
    const { t } = useTranslation();
    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGRecent cls="w-6 h-6 mr-2 fill-current" />}
            title={t("problem_statement.last_submissions")}
        />
    );
    const rows = submissions.map((item, index) => (
        <tr className="divide-x divide-dividecol" key={index}>
            <td className="padding-td-default w-0">
                <Link
                    className="link"
                    to={routeMap.submission.replace(":id", item.id)}>
                    {item.id}
                </Link>
            </td>
            <td className="padding-td-default" style={{ maxWidth: 100 }}>
                <div className="flex items-center">
                    {item.verdictType === 0 && (
                        <SVGSpinner cls="w-5 h-5 mr-2 shrink-0" />
                    )}
                    {item.verdictType === 1 && (
                        <SVGWrongSimple cls="w-5 h-5 text-red-600 mr-2 shrink-0" />
                    )}
                    {item.verdictType === 2 && (
                        <SVGPartiallyCorrect cls="w-5 h-5 text-yellow-600 mr-2 shrink-0" />
                    )}
                    {item.verdictType === 3 && (
                        <SVGCorrectSimple cls="w-5 h-5 text-green-600 mr-2 shrink-0" />
                    )}
                    <span className="truncate">{item.verdictName}</span>
                </div>
            </td>
            <td className="padding-td-default w-0 text-center">
                <span className="whitespace-nowrap">
                    {item.score} / {maxScore}
                </span>
            </td>
        </tr>
    ));
    return (
        <RoundedTable titleComponent={titleComponent}>
            <tbody className="divide-y divide-dividecol">{rows}</tbody>
        </RoundedTable>
    );
}

function ProblemAttachment({ type, name, href }) {
    const { t } = useTranslation();
    return (
        <li>
            <a
                className="link no-underline flex items-start my-0.5"
                href={`http://localhost:5555${href}`} download>
                {type === "file" && (
                    <SVGAttachmentFile cls="w-5 h-5 mr-2 shrink-0" />
                )}
                {type === "statement" && (
                    <SVGAttachmentDescription cls="w-5 h-5 mr-2 shrink-0" />
                )}
                <span className="underline truncate">
                    {type === "statement"
                        ? t("problem_statement.statement")
                        : t("problem_statement.file")}
                    &nbsp;({name})
                </span>
            </a>
        </li>
    );
}

function ProblemAttachments({ attachments }) {
    const { t } = useTranslation();
    console.log(attachments.statements);
    const attachmentsContent = attachments.statements
        .map((item, index) => (
            <ProblemAttachment
                key={index}
                type="statement"
                name={item.name}
                href={item.href}
            />
        ))
        .concat(
            attachments.files.map((item, index) => (
                <ProblemAttachment
                    key={attachments.statements.length + index}
                    type="file"
                    name={item.name}
                    href={item.href}
                />
            )),
        );
    const titleComponent = (
        <SVGTitleComponent
            svg={<SVGAttachment />}
            title={t("problem_statement.attachments")}
        />
    );
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-6 py-5">
                <ul>{attachmentsContent}</ul>
            </div>
        </RoundedFrame>
    );
}

function ProblemStatement({ data }) {
    const { theme } = useContext(ThemeContext);
    const [statementIndex, setStatementIndex] = useState(0);
    const statementSrc = data.attachments.statements[statementIndex].href;
    const statementType = data.attachments.statements[statementIndex].type;
    console.log(statementSrc + " -- " + statementType);
    return (
        <div className="flex flex-col lg:flex-row">
            <div className="w-full flex flex-col">
                <div className="w-full mb-2">
                    <RoundedFrame>
                        <div className="w-full px-4 py-3 sm:px-6 sm:py-5 flex">
                            <div className="w-full mr-3 min-w-0">
                                <DropdownMenu
                                    itemNames={[
                                        "Hungarian (Tom és Jerry)",
                                        "English (Tom and Jerry)",
                                    ]}
                                    onChange={setStatementIndex}
                                />
                            </div>
                            <a
                                className="btn-gray py-2 px-4 flex justify-center items-center"
                                href={statementSrc}
                                target="_blank"
                                rel="noreferrer">
                                <SVGView cls="w-[1.4rem] h-[1.4rem]" />
                            </a>
                        </div>
                    </RoundedFrame>
                </div>
                <div className="w-full mb-3">
                    {statementType === "pdf" && (
                        <object
                            color-scheme={theme}
                            data={statementSrc}
                            aria-label="Problem statement"
                            type="application/pdf"
                            width="100%"
                            className="h-[36rem] lg:h-[52rem] border border-grey-600"></object>
                    )}
                    {statementType === "html" && (
                        <iframe
                            src={statementSrc}
                            width="100%"
                            title="Problem statement"
                            className="h-[36rem] lg:h-[52rem] border border-grey-600"></iframe>
                    )}
                </div>
            </div>
            <div className="w-full lg:w-96 mb-3 lg:ml-3 shrink-0">
                <div className="mb-3">
                    <ProblemInfo info={data.info} />
                </div>
                <div className="mb-3">
                    <ProblemSubmit />
                </div>
                {data.lastSubmissions && data.lastSubmissions.length > 0 && (
                    <div className="mb-3">
                        <ProblemLastSubmissions
                            submissions={data.lastSubmissions}
                            maxScore={data.info.maxScore}
                        />
                    </div>
                )}
                <div className="mb-3">
                    <ProblemAttachments attachments={data.attachments} />
                </div>
            </div>
        </div>
    );
}

export default ProblemStatement;
