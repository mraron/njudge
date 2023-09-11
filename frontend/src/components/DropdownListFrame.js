import DropdownList from "./DropdownList";
import RoundedFrame from "./RoundedFrame";

function DropdownListFrame({ title, tree }) {
    return (
        <RoundedFrame>
            <div className="rounded-md overflow-hidden">
                <h1 className="text-lg font-medium px-6 py-4 text-center border-b-1 border-default">{title}</h1>
                <div className="px-8 pt-4 pb-2 bg-grey-850">
                    <DropdownList tree={tree}/>
                </div>
            </div>
        </RoundedFrame>
    );
    
}

export default DropdownListFrame;