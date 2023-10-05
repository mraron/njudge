import { useTranslation } from "react-i18next";
import RoundedFrame from "../../container/RoundedFrame";
import Button from "../../util/Button";
import { Link } from "react-router-dom";

function ContestFrame({ contest }) {
    const { t } = useTranslation();
    const { name, href, date, active } = contest;
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start space-x-4 mb-2">
                    <span className="text-base emph-strong break-words min-w-0">
                        {name}
                    </span>
                    <span className="date-label">{date}</span>
                </div>
                <div className="flex space-x-2">
                    <Link to={href}>
                        <Button color="gray">{t("contests.view")}</Button>
                    </Link>
                    {active && (
                        <Button color="indigo">{t("contests.register")}</Button>
                    )}
                </div>
            </div>
        </RoundedFrame>
    );
}

function ContestList({ contests }) {
    const contestsContent = contests.map((item, index) => (
        <ContestFrame key={index} contest={item} />
    ));
    return <div className="space-y-3">{contestsContent}</div>;
}

export default ContestList;
