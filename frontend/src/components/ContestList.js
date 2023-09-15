import RoundedFrame from './RoundedFrame'

function ContestFrame({name, date, active}) {
    const buttons = [
        <button className="btn-gray mr-1" key={0}>Megtekintés</button>
    ]
    if (active) {
        buttons.push(
            <button className="btn-indigo ml-1" key={buttons.length}>Regisztráció</button>
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