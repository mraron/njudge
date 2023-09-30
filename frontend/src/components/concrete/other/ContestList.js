import { useTranslation } from "react-i18next";
import RoundedFrame from "../../container/RoundedFrame";
import Button from "../../util/Button";

function ContestFrame({ name, date, active }) {
    const { t } = useTranslation();
    const buttons = [
        <div className="mr-2" key={0}>
            <Button theme="gray">{t("contests.view")}</Button>
        </div>,
    ];
    if (active) {
        buttons.push(
            <Button theme="indigo" key={buttons.length}>
                {t("contests.register")}
            </Button>,
        );
    }
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start">
                    <span className="text-base font-semibold break-words min-w-0">
                        {name}
                    </span>
                    <span className="ml-4 date-label">{date}</span>
                </div>
                <div className="mt-2 flex">{buttons}</div>
            </div>
        </RoundedFrame>
    );
}

function ContestList({ contestData }) {
    const contestItems = contestData.map((data, index) => {
        return (
            <div className="mb-3" key={index}>
                <ContestFrame name={data[0]} date={data[1]} active={data[2]} />
            </div>
        );
    });
    return <>{contestItems}</>;
}

export default ContestList;
