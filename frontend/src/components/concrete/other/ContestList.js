import { useTranslation } from "react-i18next";
import RoundedFrame from "../../container/RoundedFrame";
import Button from "../../util/Button";

function ContestFrame({ name, date, active }) {
    const { t } = useTranslation();
    const buttons = [
        <Button key={0} color="gray">
            {t("contests.view")}
        </Button>,
    ];
    if (active) {
        buttons.push(
            <Button color="indigo" key={buttons.length}>
                {t("contests.register")}
            </Button>,
        );
    }
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start space-x-4">
                    <span className="text-base emph-strong break-words min-w-0">
                        {name}
                    </span>
                    <span className="date-label">{date}</span>
                </div>
                <div className="mt-2 flex space-x-2">{buttons}</div>
            </div>
        </RoundedFrame>
    );
}

function ContestList({ contestData }) {
    const contestItems = contestData.map((data, index) => (
        <ContestFrame
            key={index}
            name={data[0]}
            date={data[1]}
            active={data[2]}
        />
    ));
    return <div className="space-y-3">{contestItems}</div>;
}

export default ContestList;
