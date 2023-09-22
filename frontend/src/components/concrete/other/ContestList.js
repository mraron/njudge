import RoundedFrame from '../../container/RoundedFrame'
import {useTranslation} from "react-i18next";

function ContestFrame({name, date, active}) {
    const {t} = useTranslation()
    const buttons = [
        <button className="btn-gray padding-btn-default mr-1" key={0}>{t("contests.view")}</button>
    ]
    if (active) {
        buttons.push(
            <button className="btn-indigo padding-btn-default ml-1" key={buttons.length}>{t("contests.register")}</button>
        )
    }
    return (
        <RoundedFrame>
            <div className="px-6 py-5 sm:px-10 sm:py-8">
                <div className="flex justify-between items-start">
                    <span className="text-lg font-semibold">{name}</span>
                    <span className="ml-4 date-label">{date}</span>
                </div>
                <div className="mt-2 flex">
                    {buttons}
                </div>
            </div>
        </RoundedFrame>
    )
}

function ContestList({contestData}) {
    const contestItems = contestData.map((data, index) => {
        return (
            <div className="mb-3" key={index}>
                <ContestFrame name={data[0]} date={data[1]} active={data[2]}/>
            </div>
        );
    });
    return (
        <div>
            {contestItems}
        </div>
    )
}

export default ContestList;