import MapDataFrame from '../../components/container/MapDataFrame';
import DropdownMenu from '../../components/input/DropdownMenu';
import RoundedFrame from '../../components/container/RoundedFrame';
import SVGTitleComponent from '../../svg/SVGTitleComponent';
import {SVGAttachment, SVGAttachmentDescription, SVGAttachmentFile, SVGInformation, SVGSubmit} from '../../svg/SVGs';
import {Link, useOutletContext} from "react-router-dom";
import checkData from "../../util/CheckData";

function ProblemInfo({info}) {
    const tagsContent =
        <div className="flex flex-wrap">
            {info.tags.map((tagName, index) =>
                <span className="tag" key={index}>{tagName}</span>)}
        </div>

    const titleComponent = <SVGTitleComponent svg={<SVGInformation cls="w-6 h-6 mr-2" />} title="Információk"/>
    return (
        <MapDataFrame titleComponent={titleComponent} data={[
            ["Azonosító",       info.id],
            ["Cím",             info.title],
            ["Időlimit",        `${info.timeLimit} ms`],
            ["Memórialimit",    `${info.memoryLimit} MiB`],
            ["Címkék",          tagsContent],
            ["Típus",           info.type]
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

function ProblemAttachments({attachments}) {
    const attachmentsContent = attachments.map((item, index) => {
        const typeLabel = item.type === "file" ? "Fájl" : (item.type === "statement" ? "Leírás" : "Csatolmány");
        return (
            <li key={index}>
                <Link className="link no-underline flex items-start" to={item.href}>
                    {item.type === "file" && <SVGAttachmentFile/>}
                    {item.type === "statement" && <SVGAttachmentDescription/>}
                    <span className="underline">{typeLabel} ({item.name})</span>
                </Link>
            </li>
        )
    });
    const titleComponent = <SVGTitleComponent svg={<SVGAttachment/>} title="Mellékletek"/>
    return (
        <RoundedFrame titleComponent={titleComponent}>
            <div className="px-6 py-5">
                <ul>
                    {attachmentsContent}
                </ul>
            </div>
        </RoundedFrame>
    )
}

function ProblemStatement() {
    const data = useOutletContext()
    if (!checkData(data)) {
        return
    }
    return (
        <div className="flex flex-col lg:flex-row">
            <div className="w-full mb-3">
                <object data="/assets/statement.pdf" type="application/pdf" width="100%"
                        className="h-[36rem] lg:h-[52rem]">
                </object>
            </div>
            <div className="w-full lg:w-96 mb-3 lg:ml-3 shrink-0">
                <div className="mb-3">
                    <ProblemInfo info={data.info} />
                </div>
                <div className="mb-3">
                    <ProblemSubmit />
                </div>
                <div className="mb-3">
                    <ProblemAttachments attachments={data.attachments} />
                </div>
            </div>
        </div>
    )
}

export default ProblemStatement;
