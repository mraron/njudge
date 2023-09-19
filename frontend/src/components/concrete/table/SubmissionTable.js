import RoundedTable from "../../container/RoundedTable";
import {SVGCorrectSimple, SVGSpinner, SVGWrongSimple} from "../../../svg/SVGs";
import {useState} from "react";
import {useTranslation} from "react-i18next";

function TestCase({index, numCases, testCase, group, isLastGroup, isLastCase}) {
    const bottomBorderCase = isLastGroup && isLastCase ? "border-b-0" : ""
    const bottomBorderGroup = isLastGroup ? "border-b-0" : ""
    return (
        <tr>
            {index === 0 && <>
                <td className={`padding-td-default border border-t-0 border-divide-col text-center ${bottomBorderGroup}`}
                    rowSpan={numCases}>
                    <div className="flex flex-col justify-center">
                        <div className="flex items-center justify-center mb-2">
                            {group["failed"] && <SVGWrongSimple cls="w-7 h-7 text-red-500"/>}
                            {!group["failed"] && group["completed"] &&
                                <SVGCorrectSimple cls="w-7 h-7 text-indigo-500"/>}
                            {!group["failed"] && !group["completed"] && <SVGSpinner cls="w-7 h-7"/>}
                        </div>
                        {group["name"]}
                    </div>
                </td>
                <td className={`padding-td-default border border-t-0 border-divide-col text-center ${bottomBorderGroup}`}
                    rowSpan={numCases}>
                    {`${group["score"]} / ${group["maxScore"]}`}
                </td>
            </>}
            <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>{testCase["index"]}</td>
            {group["scoring"] !== 1 &&
                <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}
                    colSpan={2}>
                    <div className="flex">
                        <SVGSpinner cls="mr-2 w-5 h-5"/>
                        <span className="whitespace-nowrap">{testCase["verdictName"]}</span>
                    </div>
                </td>
            }
            {group["scoring"] === 1 && <>
                <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>
                    <div className="flex items-center">
                        <SVGWrongSimple cls="mr-2 w-5 h-5 text-red-500"/>
                        <span className="whitespace-nowrap">{testCase["verdictName"]}</span>
                    </div>
                </td>
                <td className={`padding-td-default border border-t-0 border-divide-col whitespace-nowrap ${bottomBorderCase}`}>{testCase["score"]} / {testCase["maxScore"]}</td>
            </>}
            <td className={`padding-td-default border border-t-0 border-divide-col ${bottomBorderCase}`}>{testCase["timeSpent"]}</td>
            <td className={`padding-td-default border border-t-0 border-r-0 border-divide-col ${bottomBorderCase}`}>{testCase["memoryUsed"]}</td>
        </tr>
    )
}

function TestGroup({group, isLast}) {
    const testCases = group["testCases"]
    const testCaseContents = testCases.map((testCase, index) =>
        <TestCase index={index} numCases={testCases.length} testCase={testCase} group={group} key={index}
                  isLastGroup={isLast} isLastCase={index === testCases.length - 1}/>
    )
    return (
        <>
            {testCaseContents}
        </>
    )
}

function SubmissionTable({status}) {
    const {t} = useTranslation()
    const testSet = status["testSets"][0]
    const groups = testSet["groups"]
    const groupContents = groups.map((group, index) =>
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
            {groupContents}
            </tbody>
        </RoundedTable>
    )
}

export default SubmissionTable