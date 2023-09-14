import RoundedTable from './RoundedTable';
import {Link, useLocation} from 'react-router-dom';
import { useEffect } from 'react';
import { useState } from 'react';
import { SVGAvatar } from '../svg/SVGs';
import queryString from 'query-string';
import {routeMap} from "../config/RouteConfig";

function Problem({ problem }) {
    const {id, title, category, tags, numSolved} = problem
    const tagsContent = tags.map((item, index) =>
        <span className="tag" key={index}>{item}</span>
    );
    return (
        <tr className={"divide-x divide-default"}>
            <td className="padding-td-default">
                {id}
            </td>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.problem.replace(":problem", id)}>{title}</Link>
            </td>
            <td className="padding-td-default">
                <Link className="link" to={category.link}>{category.label}</Link>
            </td>
            <td className="padding-td-default">
                <div className="flex flex-wrap">
                    {tagsContent}
                </div>
            </td>
            <td className="padding-td-default">
                <Link className="link flex items-center" to={`${routeMap.problemSubmissions.replace(":problem", id)}?ac=1`}>
                    <SVGAvatar cls="w-4 h-4 mr-1" />
                    <span>{numSolved}</span>
                </Link>
            </td>
        </tr>
    );
}

function ProblemsTable({ problems }) {
    const problemsContent = problems.map((item, index) =>
        <Problem problem={item} key={index} />
    );
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
                <tr className="divide-x divide-default">
                    <th className="padding-td-default">Azonosító</th>
                    <th className="padding-td-default">Feladatcím</th>
                    <th className="padding-td-default">Kategória</th>
                    <th className="padding-td-default">Címkék</th>
                    <th className="padding-td-default">Megoldók</th>
                </tr>
            </thead>
            <tbody className="divide-y divide-default">
                {problemsContent}
            </tbody>
        </RoundedTable>
    );
}

export default ProblemsTable;