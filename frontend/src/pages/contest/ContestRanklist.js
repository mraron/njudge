import RoundedTable from "../../components/container/RoundedTable";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Link, useLocation, useNavigate } from "react-router-dom";
import { routeMap } from "../../config/RouteConfig";
import Pagination from "../../components/util/Pagination";
import DropdownFrame from "../../components/container/DropdownFrame";
import TextBox from "../../components/input/TextBox";
import { useTranslation } from "react-i18next";
import Button from "../../components/util/Button";
import updateQueryString from "../../util/updateQueryString";
import { useState } from "react";
import queryString from "query-string";

function RanklistRow(data) {
    const { place, name, score, verdicts } = data.result;
    const verdictsContent = verdicts.map((item, index) => {
        return (
            <td className="py-2.5 px-0 w-0" key={index}>
                <span className="flex justify-center items-center">
                    {item.verdictType === 1 && (
                        <FontAwesomeIcon
                            icon="fa-xmark"
                            className="w-3 h-3 highlight-red"
                        />
                    )}
                    {item.verdictType === 2 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-3 h-3 highlight-yellow"
                        />
                    )}
                    {item.verdictType === 3 && (
                        <FontAwesomeIcon
                            icon="fa-check"
                            className="w-3 h-3 highlight-green"
                        />
                    )}
                </span>
            </td>
        );
    });
    return (
        <tr
            className={`divide-x divide-dividecol ${
                data.index % 2 === 0 ? "bg-grey-850" : "bg-grey-825"
            } hover:bg-grey-800 cursor-pointer`}>
            <td className="py-2.5 text-center">{place}</td>
            <td className="py-2.5">
                <Link
                    className="link"
                    to={routeMap.profile.replace(":user", "dbence")}>
                    {name}
                </Link>
            </td>
            <td className="py-2.5 w-0 text-center">{score}</td>
            {verdictsContent}
        </tr>
    );
}

function RanklistFilter() {
    const { t } = useTranslation();
    const location = useLocation();
    const navigate = useNavigate();

    const qData = queryString.parse(location.search);
    const [username, setUsername] = useState(qData["user"]);
    const handleSubmit = () => {
        updateQueryString({
            location: location,
            navigate: navigate,
            args: ["user"],
            values: [username],
            validArgs: ["user"],
        });
    };
    const handleReset = () => {
        updateQueryString({
            location: location,
            navigate: navigate,
            validArgs: [],
        });
    };
    return (
        <div>
            <div className="flex flex-col space-y-4 mb-5">
                <TextBox
                    id="filterName"
                    label={t("contest_ranklist.username")}
                    initText={username}
                    onChange={setUsername}
                />
            </div>
            <div className="flex justify-center space-x-2">
                <Button color="indigo" minWidth="8rem" onClick={handleSubmit}>
                    {t("problem_filter.search")}
                </Button>
                <Button color="gray" minWidth="8rem" onClick={handleReset}>
                    {t("problem_filter.reset")}
                </Button>
            </div>
        </div>
    );
}

function ContestRanklist({ data }) {
    const { t } = useTranslation();
    const ranklistContent = data.ranklist.map((item, index) => (
        <RanklistRow result={item} index={index} key={index} />
    ));
    const problemsContent = data.problems.map((item, index) => (
        <th className="px-4" key={index}>
            <Link className="link" to={item.href}>
                {item.text}
            </Link>
        </th>
    ));
    return (
        <div className="space-y-2">
            <DropdownFrame title="Szűrés">
                <div className="px-8 py-6">
                    <RanklistFilter />
                </div>
            </DropdownFrame>
            <RoundedTable>
                <thead className="bg-grey-800">
                    <tr className="divide-x divide-dividecol">
                        <th className="w-0">#</th>
                        <th>{t("contest_ranklist.username")}</th>
                        <th className="w-0">=</th>
                        {problemsContent}
                    </tr>
                </thead>
                <tbody className="divide-y divide-dividecol">
                    {ranklistContent}
                </tbody>
            </RoundedTable>
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ContestRanklist;
