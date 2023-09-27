import CopyButton from "./CopyButton";
import RoundedFrame from "../../container/RoundedFrame";

function CopyableCode({ text, maxHeight = "auto", titleComponent, cls }) {
    return (
        <RoundedFrame
            cls={`${cls} overflow-auto`}
            titleComponent={titleComponent}>
            <div
                className="relative bg-codebgcol flex"
                style={{ minHeight: "3.75rem", maxHeight: maxHeight }}>
                <pre className="w-full px-4 py-3 overflow-x-auto">{text}</pre>
                <div className="absolute top-2.5 right-2.5 opacity-75 hover:opacity-100 transition duration-200">
                    <CopyButton text={text} />
                </div>
            </div>
        </RoundedFrame>
    );
}

export default CopyableCode;
