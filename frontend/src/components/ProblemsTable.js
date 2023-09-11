import RoundedTable from './RoundedTable';
import { useLocation } from 'react-router-dom';
import { useEffect } from 'react';
import { useState } from 'react';
import { SVGAvatar } from '../svg/SVGs';

function Problem({ id, title, category, tagTitles, solved }) {
    const tags = tagTitles.map((tagName, index) =>
        <span className="tag" key={index}>{tagName}</span>
    );
    return (
        <tr className={"divide-x divide-default"}>
            <td className="padding-td-default">
                {id}
            </td>
            <td className="padding-td-default">
                <span className="link">{title}</span>
            </td>
            <td className="padding-td-default">
                <span className="link">{category}</span>
            </td>
            <td className="padding-td-default">
                <div className="flex flex-wrap">
                    {tags}
                </div>
            </td>
            <td className="padding-td-default">
                <div className="text-indigo-200 hover:text-indigo-100 underline flex items-center cursor-pointer">
                    <span className="mr-1">
                        <SVGAvatar cls="w-4 h-4" />
                    </span>
                    <span>{solved}</span>
                </div>
            </td>
        </tr>
    );
}

function ProblemsTable() {
    const location = useLocation();
    const [problems, setProblems] = useState([
        ["NT21_Atvagas", "Átvágás", "Nemes Tihamér", ["mohó", "matematika"], "7"],
        ["NT23_Vilagnaptar", "Világnaptár (45 pont)", "Nemes Tihamér", ["mohó", "adatszerkezetek"], "7"],
        ["NT21_Atvagas", "Átvágás", "Nemes Tihamér", ["dinamikus programozás", "matematika"], "5"],
        ["NT23_Vilagnaptar", "Világnaptár (45 pont)", "Nemes Tihamér", ["mohó", "adatszerkezetek"], "4"],
    ]);
    useEffect(() =>
        setProblems(prevProblems => prevProblems.concat([
            [location.search, "Világnaptár (45 pont)", "Nemes Tihamér", ["mohó", "adatszerkezetek"], "4"]
        ])), [location.search]);

    const problemElems = problems.map((item, index) => 
        <Problem id={item[0]} title={item[1]} category={item[2]} tagTitles={item[3]} solved={item[4]} key={index} />
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
                {problemElems}
            </tbody>
        </RoundedTable>
    );
}

export default ProblemsTable;