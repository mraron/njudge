import RoundedFrame from './RoundedFrame';
import React, { useState } from 'react';
import { SVGDropdownFilterArrow } from '../svg/SVGs';

function DropdownFrame({ children }) {
    const [isOpen, setOpen] = useState(false);
    return (
        <RoundedFrame>
            <button onClick={() => setOpen(!isOpen)} className={`${isOpen? "bg-grey-750 hover:bg-grey-725 rounded-tl-md rounded-tr-md": "bg-grey-800 hover:bg-grey-775 rounded-md"} transiton-all duration-200 border-default flex items-center justify-center`}>
                <SVGDropdownFilterArrow isOpen={isOpen} />
            </button>
            <div className={`${isOpen? "": "h-0 overflow-hidden"}`}>
                {children}
            </div>
        </RoundedFrame>
    );
}

export default DropdownFrame;