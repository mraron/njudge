import { useContext, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { useTranslation } from "react-i18next";
import Checkbox from "../../input/Checkbox";
import Button from "../../util/Button";
import DropdownFrame from "../../container/DropdownFrame";
import updateQueryString from "../../../util/updateQueryString";
import UserContext from "../../../contexts/user/UserContext";
import queryString from "query-string";

function SubmissionFilter({ optionOwn }) {
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
        updateQueryString({
            location: location,
            navigate: navigate,
            args: ["ac", "own"],
            values: [onlyAC ? 1 : 0, onlyOwn ? 1 : 0],
            validArgs: ["ac", "own"],
        });
    };
    return (
        <div className="flex flex-col">
            {isLoggedIn && optionOwn && (
                <div className="mb-3">
                    <Checkbox
                        label={t("submission_filter.own_solutions")}
                        onChange={setOnlyOwn}
                        initChecked={onlyOwn}
                    />
                </div>
            )}
            <div className="mb-5">
                <Checkbox
                    label={t("submission_filter.full_solutions")}
                    onChange={setOnlyAC}
                    initChecked={onlyAC}
                />
            </div>
            <div className="flex justify-center space-x-2">
                <Button color="indigo" onClick={handleSubmit} minWidth="8rem">
                    {t("submission_filter.search")}
                </Button>
                <Button color="gray" onClick={handleReset} minWidth="8rem">
                    {t("submission_filter.reset")}{" "}
                </Button>
            </div>
        </div>
    );
}

function SubmissionFilterFrame({ optionOwn = true }) {
    const { t } = useTranslation();
    return (
        <DropdownFrame title={t("submission_filter.filter")}>
            <div className="px-8 py-6">
                <SubmissionFilter optionOwn={optionOwn} />
            </div>
        </DropdownFrame>
    );
}

export default SubmissionFilterFrame;
