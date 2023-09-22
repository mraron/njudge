import RoundedTable from "../../container/RoundedTable";
import {SVGClock, SVGCopy, SVGCorrectSimple, SVGHDD, SVGSpinner, SVGWrong, SVGWrongSimple} from "../../../svg/SVGs";
import React, {useState} from "react";
import {useTranslation} from "react-i18next";
import RoundedFrame from "../../container/RoundedFrame";
import CopyableCode from "../../util/copy/CopyableCode";

function TestCase_1_3({index, numCases, testCase, group, isLastGroup, isLastCase}) {
    const bottomBorderCase = isLastGroup && isLastCase ? "border-b-0" : ""
    const bottomBorderGroup = isLastGroup ? "border-b-0" : ""
    return (
        <tr>
            {index === 0 && <>
                <td className={`padding-td-default border border-t-0 border-divide-col text-center ${bottomBorderGroup}`}
                    rowSpan={numCases}>
                    <div className="flex flex-col justify-center">
                        <div className="flex items-center justify-center mb-2">
                            {group.failed && <SVGWrongSimple cls="w-7 h-7 text-red-500"/>}
                            {!group.failed && group.completed &&
                                <SVGCorrectSimple cls="w-7 h-7 text-indigo-500"/>}
                            {!group.failed && !group.completed && <SVGSpinner cls="w-7 h-7"/>}
                        </div>
                        {group.name}
                    </div>
                </td>
                <td className={`padding-td-default border border-t-0 border-divide-col text-center ${bottomBorderGroup}`}
                    rowSpan={numCases}>
                    {`${group.score} / ${group.maxScore}`}
                </td>
            </>}
            <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>{testCase.index}</td>
            {group.scoring !== 1 &&
                <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}
                    colSpan={2}>
                    <div className="flex">
                        {testCase.verdictType === 0 && <SVGSpinner cls="w-5 h-5 mr-2 shrink-0"/>}
                        {testCase.verdictType === 1 && <SVGWrongSimple cls="w-5 h-5 text-red-500 mr-2 shrink-0"/>}
                        {testCase.verdictType === 2 && <SVGCorrectSimple cls="w-5 h-5 text-indigo-500 mr-2 shrink-0"/>}
                        {testCase.verdictType === 3 && <SVGCorrectSimple cls="w-5 h-5 text-green-500 mr-2 shrink-0"/>}
                        <span className="whitespace-nowrap">{testCase.verdictName}</span>
                    </div>
                </td>
            }
            {group.scoring === 1 && <>
                <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>
                    <div className="flex items-center">
                        <SVGWrongSimple cls="mr-2 w-5 h-5 text-red-500"/>
                        <span className="whitespace-nowrap">{testCase.verdictName}</span>
                    </div>
                </td>
                <td className={`padding-td-default border border-t-0 border-divide-col whitespace-nowrap ${bottomBorderCase}`}>{testCase.score} / {testCase.maxScore}</td>
            </>}
            <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>{testCase.timeSpent}</td>
            <td className={`padding-td-default border border-t-0 border-r-0 border-divide-col ${bottomBorderCase}`}>{testCase.memoryUsed}</td>
        </tr>
    )
}

function TestGroup({group, isLast}) {
    const testCases = group.testCases
    const testCasesContent = testCases.map((testCase, index) =>
        <TestCase_1_3 index={index} numCases={testCases.length} testCase={testCase} group={group} key={index}
                  isLastGroup={isLast} isLastCase={index === testCases.length - 1}/>
    )
    return (
        <>
            {testCasesContent}
        </>
    )
}

function TestCase_0({testCase, index}) {
    const {t} = useTranslation()
    const titleComponent =
        <div className="py-3 px-5 border-b border-default flex items-center text-table">
            {testCase.verdictType === 0 && <SVGSpinner cls="mr-3 w-6 h-6 shrink-0"/>}
            {testCase.verdictType === 1 && <SVGWrongSimple cls="mr-3 w-6 h-6 text-red-500 shrink-0"/>}
            {testCase.verdictType === 2 && <SVGCorrectSimple cls="mr-3 w-6 h-6 text-indigo-500 shrink-0"/>}
            {testCase.verdictType === 3 && <SVGCorrectSimple cls="mr-3 w-6 h-6 text-green-500 shrink-0"/>}
            <span>{index + 1}</span>
            <span className="mx-2">–</span>
            <span>{testCase.verdictName}</span>
        </div>

    const outputRows = [
        ["submission_table.output", testCase.output],
        ["submission_table.expected_output", testCase.expectedOutput],
        ["submission_table.checker_output", testCase.checkerOutput],
    ].map(item =>
        <tr className="divide-x divide-default">
            <td className="padding-td-default whitespace-nowrap w-48">{t(item[0])}</td>
            <td style={{maxWidth: 0}}>
                <div className="flex py-2 px-3">
                    <CopyableCode text={item[1]} isMultiline={true} maxHeight={"6rem"}/>
                </div>
            </td>
        </tr>
    )
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="py-3 px-5 flex justify-between border-b border-default">
                <div className="mr-2 whitespace-nowrap truncate flex items-center">
                    <SVGClock cls="w-5 h-5 mr-2"/>
                    <span className="hidden sm:inline">{t("submission_table.time")}:&nbsp;</span>
                    <span>{testCase.timeSpent}</span>
                </div>
                <div className="ml-2 whitespace-nowrap truncate flex items-center">
                    <SVGHDD cls="w-5 h-5 mr-2"/>
                    <span className="hidden sm:inline">{t("submission_table.memory")}:&nbsp;</span>
                    <span>{testCase.memoryUsed}</span>
                </div>
            </div>
            <div className="overflow-x-auto">
                <table className="w-full text-table bg-grey-850 rounded-b-md">
                    <tbody className="divide-y divide-default">
                        {outputRows}
                    </tbody>
                </table>
            </div>
        </RoundedFrame>
    )
}

function SubmissionTable_0({status}) {
    const testCases = status.groups[0].testCases
    const testCasesContent = testCases.map((testCase, index) =>
        <div className="mb-3">
            <TestCase_0 testCase={testCase} index={index} key={index}/>
        </div>
    )
    return (
        <>
            {testCasesContent}
        </>
    )
}

function SubmissionTable_1_3({status}) {
    const {t} = useTranslation()
    const groups = status.groups
    const groupsContent = groups.map((group, index) =>
        <TestGroup group={group} isLast={index === groups.length - 1} key={index}/>
    )
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
            <tr className="divide-x divide-default">
                <th className="padding-td-default">{t("submission_table.subtask")}</th>
                <th className="padding-td-default">{t("submission_table.total")}</th>
                <th className="padding-td-default">{t("submission_table.test")}</th>
                <th className="padding-td-default" colSpan={2}>{t("submission_table.verdict")}</th>
                <th className="padding-td-default">{t("submission_table.time")}</th>
                <th className="padding-td-default">{t("submission_table.memory")}</th>
            </tr>
            </thead>
            <tbody>
                {groupsContent}
            </tbody>
        </RoundedTable>
    )
}

function SubmissionTable({status}) {
    if (status.feedbackType === 0) {
        return (
            <SubmissionTable_0 status={status} />
        )
    }
    if ([1, 3].includes(status.feedbackType)) {
        return (
            <SubmissionTable_1_3 status={status} />
        )
    }
}

export default SubmissionTable