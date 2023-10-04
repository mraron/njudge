import React, { useContext, useState } from "react";
import { useTranslation } from "react-i18next";
import { Link, useNavigate, useParams } from "react-router-dom";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import MapDataFrame from "../../components/container/MapDataFrame";
import DropdownMenu from "../../components/input/DropdownMenu";
import RoundedFrame from "../../components/container/RoundedFrame";
import SVGTitleComponent from "../../components/svg/SVGTitleComponent";
import { SVGSpinner, SVGView } from "../../components/svg/SVGs";
import RoundedTable from "../../components/container/RoundedTable";
import TagModal from "../../components/container/modal/TagModal";
import JudgeDataContext from "../../contexts/judgeData/JudgeDataContext";
import ThemeContext from "../../contexts/theme/ThemeContext";
import submitSolution from "../../util/submitSolution";
import { apiRoute, routeMap } from "../../config/RouteConfig";
import Tag from "../../components/util/Tag";
import Button from "../../components/util/Button";

function ProblemInfo({ info }) {
    const { t } = useTranslation();
    const [isModalOpen, setModalOpen] = useState(false);
    const tagsContent = (
        <div className="flex-col">
            <div className="flex flex-wrap -m-1">
                {info.tags.map((tagName, index) => (
                    <Tag key={index}>{t(tagName)}</Tag>
                ))}
                {info.isSolved && (
                    <button
                        key={info.tags.length}
                        onClick={() => setModalOpen(true)}>
                        <Tag cls="items-center">
                            <FontAwesomeIcon
                                icon="fa-regular fa-edit"
                                className="w-5 h-3.5"
                            />
                        </Tag>
                    </button>
                )}
            </div>
        </div>
    );
    const titleComponent = (
        <SVGTitleComponent
            svg={<FontAwesomeIcon icon="fa-info" className="w-4 h-4 mr-3" />}
            title={t("problem_statement.information")}
        />
    );
    return (
        <div>
            {info.isSolved && (
                <TagModal
                    isOpen={isModalOpen}
                    onClose={() => setModalOpen(false)}
                />
            )}
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
        </div>
    );
}

function ProblemSubmit() {
    const { t } = useTranslation();
    const { judgeData } = useContext(JudgeDataContext);
    const { problem, problemset } = useParams();
    const [file, setFile] = useState(null);
    const [langIndex, setLangIndex] = useState(0);
    const navigate = useNavigate();
    const titleComponent = (
        <SVGTitleComponent
            svg={
                <FontAwesomeIcon
                    icon="fa-regular fa-paper-plane"
                    className="w-4 h-4 mr-3"
                />
            }
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
            problemset: problemset,
            problem: problem,
            language: judgeData.languages[langIndex].id,
            file: file,
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
                    <div className="flex justify-center space-x-2">
                        <Button
                            color="gray"
                            onClick={() =>
                                document.getElementById("uploadFile").click()
                            }
                            fullWidth={true}>
                            <input
                                id="uploadFile"
                                type="file"
                                className="hidden"
                                onChange={handleFileUploaded}
                            />
                            {t("problem_statement.choose")}
                        </Button>
                        <Button
                            color="indigo"
                            onClick={handleSubmit}
                            fullWidth={true}>
                            {t("problem_statement.submit")}
                        </Button>
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
            svg={
                <FontAwesomeIcon
                    icon="fa-regular fa-clock"
                    className="w-4 h-4 mr-3"
                />
            }
            title={t("problem_statement.last_submissions")}
        />
    );
    const rows = submissions.map((item, index) => (
        <tr className="divide-x divide-dividecol" key={index}>
            <td className=" w-0">
                <Link
                    className="link"
                    to={routeMap.submission.replace(":id", item.id)}>
                    {item.id}
                </Link>
            </td>
            <td style={{ maxWidth: 100 }}>
                <div className="flex items-center">
                    {item.verdictType === 0 && (
                        <SVGSpinner cls="w-4 h-4 mr-3" />
                    )}
                    {item.verdictType === 1 && (
                        <FontAwesomeIcon
                            icon="fa-xmark"
                            className="w-4 h-4 highlight-red mr-3"
                        />
                    )}
                    {item.verdictType === 2 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-yellow mr-3"
                        />
                    )}
                    {item.verdictType === 3 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-4 h-4 highlight-green mr-3"
                        />
                    )}
                    <span className="truncate">{item.verdictName}</span>
                </div>
            </td>
            <td className=" w-0 text-center">
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
                className="link no-underline flex items-center my-0.5"
                href={apiRoute(href)}
                download="statement.pdf"
                target="_blank"
                rel="noreferrer">
                {type === "file" && (
                    <FontAwesomeIcon
                        icon="fa-regular fa-file"
                        className="w-4 h-4 mr-3"
                    />
                )}
                {type === "statement" && (
                    <FontAwesomeIcon
                        icon="fa-regular fa-file-lines"
                        className="w-4 h-4 mr-3"
                    />
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
            svg={
                <FontAwesomeIcon icon="fa-paperclip" className="w-4 h-4 mr-3" />
            }
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
    return (
        <div className="flex flex-col lg:flex-row space-y-3 lg:space-y-0 lg:space-x-3">
            <div className="w-full flex flex-col min-w-0 space-y-2">
                <div className="w-full">
                    <RoundedFrame>
                        <div className="w-full px-4 py-3 sm:px-6 sm:py-5 flex">
                            <div className="w-full mr-3 min-w-0">
                                <DropdownMenu
                                    itemNames={data.attachments.statements.map(
                                        (item) => item.name,
                                    )}
                                    onChange={setStatementIndex}
                                />
                            </div>
                            <a
                                href={apiRoute(statementSrc)}
                                target="_blank"
                                rel="noreferrer">
                                <Button color="gray">
                                    <SVGView cls="w-[1.4rem] h-[1.4rem]" />
                                </Button>
                            </a>
                        </div>
                    </RoundedFrame>
                </div>
                <div className="w-full">
                    {statementType === "pdf" && (
                        <object
                            color-scheme={theme}
                            data={apiRoute(statementSrc)}
                            aria-label="Problem statement"
                            type="application/pdf"
                            width="100%"
                            className="h-[36rem] lg:h-[52rem] border border-bordefcol"></object>
                    )}
                    {statementType === "html" && (
                        <iframe
                            src={apiRoute(statementSrc)}
                            width="100%"
                            title="Problem statement"
                            className="h-[36rem] lg:h-[52rem] border border-bordefcol"></iframe>
                    )}
                </div>
            </div>
            <div className="w-full lg:w-96 shrink-0 space-y-3">
                <ProblemInfo info={data.info} />
                <ProblemSubmit />
                {data.lastSubmissions && data.lastSubmissions.length > 0 && (
                    <ProblemLastSubmissions
                        submissions={data.lastSubmissions}
                        maxScore={data.info.maxScore}
                    />
                )}
                <ProblemAttachments attachments={data.attachments} />
            </div>
        </div>
    );
}

export default ProblemStatement;
