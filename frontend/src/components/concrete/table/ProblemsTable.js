import RoundedTable from '../../container/RoundedTable';
import {Link} from 'react-router-dom';
import {SVGAvatar} from '../../../svg/SVGs';
import {routeMap} from "../../../config/RouteConfig";
import {useTranslation} from "react-i18next";

function Problem(data) {
    const {problem, title, category, tags, solverCount} = data.problem
    const tagsContent = tags.map((item, index) =>
        <span className="tag" key={index}>{item}</span>
    );
    return (
        <tr className={"divide-x divide-default"}>
            <td className="padding-td-default">
                {problem}
            </td>
            <td className="padding-td-default">
                <Link className="link" to={routeMap.problem.replace(":problem", problem)}>{title}</Link>
            </td>
            <td className="padding-td-default">
                <Link className="link" to={category.href}>{category.text}</Link>
            </td>
            <td className="padding-td-default">
                <div className="flex flex-wrap">
                    {tagsContent}
                </div>
            </td>
            <td className="padding-td-default">
                <Link className="link flex items-center justify-center"
                      to={`${routeMap.problemSubmissions.replace(":problem", problem)}?ac=1`}>
                    <SVGAvatar cls="w-[1.1rem] h-[1.1rem] mr-1"/>
                    <span>{solverCount}</span>
                </Link>
            </td>
        </tr>
    );
}

function ProblemsTable({problems}) {
    const {t} = useTranslation()
    const problemsContent = problems.map((item, index) =>
        <Problem problem={item} key={index}/>
    );
    return (
        <RoundedTable>
            <thead className="bg-grey-800">
            <tr className="divide-x divide-default">
                <th className="padding-td-default">{t("problems_table.id")}</th>
                <th className="padding-td-default">{t("problems_table.title")}</th>
                <th className="padding-td-default">{t("problems_table.category")}</th>
                <th className="padding-td-default">{t("problems_table.tags")}</th>
                <th className="padding-td-default">{t("problems_table.solved")}</th>
            </tr>
            </thead>
            <tbody className="divide-y divide-default">
                {problemsContent}
            </tbody>
        </RoundedTable>
    );
}

export default ProblemsTable;