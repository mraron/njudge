import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import RoundedTable from "../../container/RoundedTable";
import { SVGSpinner } from "../../svg/SVGs";
import RoundedFrame from "../../container/RoundedFrame";
import CopyableCode from "../../util/copy/CopyableCode";
import React from "react";

function TestCase13({
    index,
    numCases,
    testCase,
    group,
    isLastGroup,
    isLastCase,
}) {
    const bottomBorderCase = isLastGroup && isLastCase ? "border-b-0" : "";
    const bottomBorderGroup = isLastGroup ? "border-b-0" : "";
    return (
        <tr>
            {index === 0 && (
                <>
                    <td
                        className={`padding-td-default border border-t-0 border-dividecol text-center ${bottomBorderGroup}`}
                        rowSpan={numCases}>
                        <div className="flex flex-col justify-center">
                            <div className="flex items-center justify-center mb-2">
                                {group.failed && (
                                    <FontAwesomeIcon
                                        icon="fa-xmark"
                                        className="w-6 h-6 highlight-red"
                                    />
                                )}
                                {!group.failed && group.completed && (
                                    <FontAwesomeIcon
                                        icon="fa-check"
                                        className="w-6 h-6 highlight-green"
                                    />
                                )}
                                {!group.failed && !group.completed && (
                                    <SVGSpinner cls="w-6 h-6" />
                                )}
                            </div>
                            {group.name}
                        </div>
                    </td>
                    <td
                        className={`padding-td-default border border-t-0 border-dividecol text-center ${bottomBorderGroup}`}
                        rowSpan={numCases}>
                        {`${group.score} / ${group.maxScore}`}
                    </td>
                </>
            )}
            <td
                className={`padding-td-default border border-t-0 border-dividecol ${bottomBorderCase}`}>
                {testCase.index}
            </td>
            {group.scoring !== 1 && (
                <td
                    className={`padding-td-default border border-t-0 border-dividecol ${bottomBorderCase}`}
                    colSpan={2}>
                    <div className="flex">
                        {testCase.verdictType === 0 && (
                            <SVGSpinner cls="w-4 h-4 mr-2" />
                        )}
                        {testCase.verdictType === 1 && (
                            <FontAwesomeIcon
                                icon="fa-xmark"
                                className="w-4 h-4 highlight-red mr-2"
                            />
                        )}
                        {testCase.verdictType === 2 && (
                            <FontAwesomeIcon
                                icon="fa-check"
                                className="w-4 h-4 highlight-yellow mr-2"
                            />
                        )}
                        {testCase.verdictType === 3 && (
                            <FontAwesomeIcon
                                icon="fa-check"
                                className="w-4 h-4 highlight-green mr-2"
                            />
                        )}
                        <span className="whitespace-nowrap">
                            {testCase.verdictName}
                        </span>
                    </div>
                </td>
            )}
            {group.scoring === 1 && (
                <>
                    <td
                        className={`padding-td-default border border-t-0 border-dividecol ${bottomBorderCase}`}>
                        <div className="flex items-center">
                            <FontAwesomeIcon
                                icon="fa-xmark"
                                className="w-4 h-4 highlight-red mr-2"
                            />
                            <span className="whitespace-nowrap">
                                {testCase.verdictName}
                            </span>
                        </div>
                    </td>
                    <td
                        className={`padding-td-default border border-t-0 border-dividecol whitespace-nowrap ${bottomBorderCase}`}>
                        {testCase.score} / {testCase.maxScore}
                    </td>
                </>
            )}
            <td
                className={`padding-td-default border border-t-0 border-dividecol ${bottomBorderCase}`}>
                {testCase.time} ms
            </td>
            <td
                className={`padding-td-default border border-t-0 border-r-0 border-dividecol ${bottomBorderCase}`}>
                {testCase.memory} KiB
            </td>
        </tr>
    );
}

function TestGroup({ group, isLast }) {
    const testCases = group.testCases;
    const testCasesContent = testCases.map((testCase, index) => (
        <TestCase13
            index={index}
            numCases={testCases.length}
            testCase={testCase}
            group={group}
            key={index}
            isLastGroup={isLast}
            isLastCase={index === testCases.length - 1}
        />
    ));
    return <>{testCasesContent}</>;
}

function TestCase0({ testCase, index }) {
    const { t } = useTranslation();
    const titleComponent = (
        <div className="py-3 px-5 border-b border-bordercol flex items-center text-table">
            {testCase.verdictType === 0 && <SVGSpinner cls="w-5 h-5 mr-3" />}
            {testCase.verdictType === 1 && (
                <FontAwesomeIcon
                    icon="fa-xmark"
                    className="w-5 h-5 highlight-red mr-3"
                />
            )}
            {testCase.verdictType === 2 && (
                <FontAwesomeIcon
                    icon="fa-check"
                    className="w-5 h-5 highlight-yellow mr-3"
                />
            )}
            {testCase.verdictType === 3 && (
                <FontAwesomeIcon
                    icon="fa-check"
                    className="w-5 h-5 highlight-green mr-3"
                />
            )}
            <span>{index + 1}</span>
            <span className="mx-2">–</span>
            <span className="truncate">{testCase.verdictName}</span>
        </div>
    );

    const outputRows = [
        ["submission_table.output", testCase.output],
        ["submission_table.expected_output", testCase.expectedOutput],
        ["submission_table.checker_output", testCase.checkerOutput],
    ].map((item, index) => (
        <tr className="divide-x divide-dividecol" key={index}>
            <td className="padding-td-default whitespace-nowrap w-48">
                {t(item[0])}
            </td>
            <td style={{ maxWidth: 0 }}>
                <CopyableCode
                    cls="border-0 rounded-none"
                    text={item[1]}
                    isMultiline={true}
                    maxHeight={"6rem"}
                />
            </td>
        </tr>
    ));
    return (
        <RoundedFrame cls="overflow-hidden" titleComponent={titleComponent}>
            <div className="py-3 px-5 flex justify-between border-b border-bordercol">
                <div className="mr-2 whitespace-nowrap truncate flex items-center">
                    <FontAwesomeIcon
                        icon="fa-regular fa-clock"
                        className="w-4 h-4 mr-2"
                    />
                    <span className="hidden sm:inline">
                        {t("submission_table.time")}:&nbsp;
                    </span>
                    <span className="truncate">{testCase.time} ms</span>
                </div>
                <div className="ml-2 whitespace-nowrap truncate flex items-center">
                    <FontAwesomeIcon
                        icon="fa-regular fa-hdd"
                        className="w-4 h-4 mr-2"
                    />
                    <span className="hidden sm:inline">
                        {t("submission_table.memory")}:&nbsp;
                    </span>
                    <span className="truncate">{testCase.memory} KiB</span>
                </div>
            </div>
            <div className="overflow-x-auto rounded-b-container">
                <table className="w-full text-table bg-grey-850 rounded-b-container">
                    <tbody className="divide-y divide-dividecol">
                        {outputRows}
                    </tbody>
                </table>
            </div>
        </RoundedFrame>
    );
}

function SubmissionTable0({ status }) {
    const testCases = status.groups?.[0].testCases;
    const testCasesContent = testCases?.map((testCase, index) => (
        <div className="mb-3" key={index}>
            <TestCase0 testCase={testCase} index={index} key={index} />
        </div>
    ));
    return <>{testCasesContent}</>;
}

function SubmissionTable13({ status }) {
    const { t } = useTranslation();
    const groups = status.groups;
    const groupsContent = groups.map((group, index) => (
        <TestGroup
            group={group}
            isLast={index === groups.length - 1}
            key={index}
        />
    ));
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
                <tr className="divide-x divide-dividecol">
                    <th className="padding-td-default">
                        {t("submission_table.subtask")}
                    </th>
                    <th className="padding-td-default">
                        {t("submission_table.total")}
                    </th>
                    <th className="padding-td-default">
                        {t("submission_table.test")}
                    </th>
                    <th className="padding-td-default" colSpan={2}>
                        {t("submission_table.verdict")}
                    </th>
                    <th className="padding-td-default">
                        {t("submission_table.time")}
                    </th>
                    <th className="padding-td-default">
                        {t("submission_table.memory")}
                    </th>
                </tr>
            </thead>
            <tbody>{groupsContent}</tbody>
        </RoundedTable>
    );
}

function SubmissionTable({ status }) {
    if (status.feedbackType === 0) {
        return <SubmissionTable0 status={status} />;
    }
    if ([1, 3].includes(status.feedbackType)) {
        return <SubmissionTable13 status={status} />;
    }
}

export default SubmissionTable;
