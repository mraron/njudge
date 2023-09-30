import { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import queryString from "query-string";
import Checkbox from "../../components/input/Checkbox";
import Pagination from "../../components/util/Pagination";
import DropdownFrame from "../../components/container/DropdownFrame";
import SubmissionsTable from "../../components/concrete/table/SubmissionsTable";
import UpdateQueryString from "../../util/updateQueryString";
import updateQueryString from "../../util/updateQueryString";
import UserContext from "../../contexts/user/UserContext";
import Button from "../../components/util/Button";

function SubmissionFilterFrame() {
    const { t } = useTranslation();
    const { isLoggedIn } = useContext(UserContext);
    const location = useLocation();
    const navigate = useNavigate();

    const qData = queryString.parse(location.search);
    const [onlyAC, setOnlyAC] = useState(qData["ac"] === "1");
    const [onlyOwn, setOnlyOwn] = useState(qData["own"] === "1");

    const handleReset = () => {
        updateQueryString({
            location: location,
            navigate: navigate,
            validArgs: [],
        });
    };
    const handleSubmit = () => {
        UpdateQueryString({
            location: location,
            navigate: navigate,
            args: ["ac", "own"],
            values: [onlyAC ? 1 : 0, onlyOwn ? 1 : 0],
            validArgs: ["ac", "own"],
        });
    };
    return (
        <DropdownFrame title={t("problem_submissions.filter")}>
            <div className="px-8 py-6 flex flex-col">
                {isLoggedIn && (
                    <div className="mb-3">
                        <Checkbox
                            label={t("problem_submissions.own_solutions")}
                            onChange={setOnlyOwn}
                            initChecked={onlyOwn}
                        />
                    </div>
                )}
                <div className="mb-5">
                    <Checkbox
                        label={t("problem_submissions.full_solutions")}
                        onChange={setOnlyAC}
                        initChecked={onlyAC}
                    />
                </div>
                <div className="flex justify-center">
                    <div className="mr-2">
                        <Button
                            color="indigo"
                            onClick={handleSubmit}
                            minWidth="8rem">
                            {t("problem_filter.search")}
                        </Button>
                    </div>
                    <Button color="gray" onClick={handleReset} minWidth="8rem">
                        {t("problem_filter.reset")}{" "}
                    </Button>
                </div>
            </div>
        </DropdownFrame>
    );
}

function ProblemSubmissions({ data }) {
    return (
        <div className="relative">
            <div className="mb-2">
                <SubmissionFilterFrame />
            </div>
            <div className="mb-2">
                <SubmissionsTable submissions={data.submissions} />
            </div>
            <Pagination paginationData={data.paginationData} />
        </div>
    );
}

export default ProblemSubmissions;
