import DropdownList from "../util/DropdownList";
import RoundedFrame from "./RoundedFrame";

function DropdownListFrame({title, tree}) {
    return (
        <RoundedFrame title={title}>
            <div className="rounded-md overflow-hidden">
                <div className="px-8 pt-4 pb-2 bg-grey-850">
                    <DropdownList tree={tree}/>
                </div>
            </div>
        </RoundedFrame>
    );

}

export default DropdownListFrame;