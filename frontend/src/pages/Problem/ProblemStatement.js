import MapDataFrame from '../../components/MapDataFrame';
import DropdownMenu from '../../components/DropdownMenu';
import RoundedFrame from '../../components/RoundedFrame';
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGAttachment, SVGAttachmentDescription, SVGAttachmentFile, SVGInformation, SVGSubmit} from '../../svg/SVGs';

const attachmentSVGs = {
    "file": <SVGAttachmentFile/>,
    "description": <SVGAttachmentDescription/>
}

function ProblemInfo() {
    const tagNames = ["oszd meg és uralkodj", "dp", "adatszerkezetek", "bináris keresés"]
    const tags =
        <div className="flex flex-wrap">
            {tagNames.map((tagName, index) =>
                <span className="tag" key={index}>{tagName}</span>)}
        </div>

    const titleComponent = <SVGTitleComponent svg={<SVGInformation/>} title="Információk"/>

    return (
        <MapDataFrame titleComponent={titleComponent} data={[
            ["Azonosító", "OKTV23_Szivarvanyszamok"],
            ["Cím", "Az óvodai lét elviselhetetlen könnyűsége"],
            ["Időlimit", "300 ms"],
            ["Memórialimit", "31 MiB"],
            ["Címkék", tags],
            ["Típus", "batch"]
        ]}/>
    )
}

function ProblemSubmit() {
    const titleComponent = <SVGTitleComponent svg={<SVGSubmit/>} title="Megoldás beküldése"/>
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-6 py-5">
                <div className="flex flex-col">
                    <div className="mb-4">
                        <DropdownMenu itemNames={["C++ 11", "C++ 14", "C++ 17", "Go", "Java", "Python 3"]}/>
                    </div>
                    <div className="mb-2 mx-1 text-label">
                        Nincs kiválasztva fájl.
                    </div>
                    <div className="flex justify-center">
                        <button className="btn-gray w-1/2">Tallózás</button>
                        <button className="ml-2 btn-indigo w-1/2">Beküldés</button>
                    </div>
                </div>
            </div>
        </RoundedFrame>
    )
}

function ProblemAttachments() {
    const attachments = [
        ["minta.zip", "file"],
        ["english", "description"],
        ["hungarian", "description"]
    ];
    const attachmentElems = attachments.map((pair, index) => {
        const type = pair[1] === "file" ? "Fájl" : (pair[1] === "description" ? "Leírás" : "Csatolmány");
        return (
            <li key={index}
                className="flex items-center cursor-pointer text-indigo-200 hover:text-indigo-100 transition duration-200">
                {attachmentSVGs[pair[1]]}
                <span className="underline">{type} ({pair[0]})</span>
            </li>
        )
    });
    const titleComponent = <SVGTitleComponent svg={<SVGAttachment/>} title="Mellékletek"/>
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-6 py-5">
                <ul>
                    {attachmentElems}
                </ul>
            </div>
        </RoundedFrame>
    )
}

function ProblemStatement() {
    return (
        <div className="flex flex-col lg:flex-row">
            <div className="w-full mb-3">
                <object data="/assets/statement.pdf" type="application/pdf" width="100%"
                        className="h-[36rem] lg:h-[52rem]">

                </object>
            </div>
            <div className="w-full lg:w-96 mb-3 lg:ml-3 shrink-0">
                <div className="mb-3">
                    <ProblemInfo/>
                </div>
                <div className="mb-3">
                    <ProblemSubmit/>
                </div>
                <div className="mb-3">
                    <ProblemAttachments/>
                </div>
            </div>
        </div>
    )
}

export default ProblemStatement;
