import DropdownList from "../util/DropdownList";
import RoundedFrame from "./RoundedFrame";

function DropdownListFrame({ title, tree, leaf: Leaf }) {
    return (
        <RoundedFrame title={title}>
            <div className="px-8 pt-4 pb-2 bg-grey-850 rounded-b-container">
                <DropdownList tree={tree} leaf={Leaf} />
            </div>
        </RoundedFrame>
    );
}

export default DropdownListFrame;
