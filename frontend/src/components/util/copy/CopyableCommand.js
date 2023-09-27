import CopyButton from "./CopyButton";
import RoundedFrame from "../../container/RoundedFrame";

function CopyableCommand({ text, cls }) {
    return (
        <RoundedFrame cls={`${cls} overflow-auto`}>
            <div className="relative bg-codebgcol flex">
                <div className="flex items-center w-full px-4 py-3 overflow-x-auto whitespace-nowrap font-code">
                    {text}
                </div>
                <div className="absolute right-1 top-1/2 transform -translate-y-1/2 opacity-75 hover:opacity-100 transition duration-200">
                    <CopyButton text={text} />
                </div>
            </div>
        </RoundedFrame>
    );
}

export default CopyableCommand;
